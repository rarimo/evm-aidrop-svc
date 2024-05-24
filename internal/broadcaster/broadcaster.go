package broadcaster

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/evm-airdrop-svc/internal/config"
	"github.com/rarimo/evm-airdrop-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

const (
	byteSize            = 32
	transferFnSignature = "transfer(address,uint256)"
)

type Runner struct {
	log *logan.Entry
	q   *data.AirdropsQ
	config.Broadcaster
	config.AirdropConfig

	errChan chan *channelData
}

type channelData struct {
	err     error
	tx      *types.Transaction
	airdrop data.Airdrop
}

func Run(ctx context.Context, cfg *config.Config) {
	log := cfg.Log().WithField("service", "builtin-broadcaster")
	log.Info("Starting service")

	r := &Runner{
		log:           log,
		q:             data.NewAirdropsQ(cfg.DB().Clone()),
		Broadcaster:   cfg.Broadcaster(),
		AirdropConfig: cfg.AridropConfig(),

		errChan: make(chan *channelData),
	}

	running.WithBackOff(ctx, r.log, "builtin-broadcaster", r.run, 5*time.Second, 5*time.Second, 5*time.Second)
}

func (r *Runner) run(ctx context.Context) error {
	go r.waitForTxErrors(ctx)

	airdrops, err := r.q.New().FilterByStatuses(data.TxStatusPending).Limit(r.QueryLimit).Select()
	if err != nil {
		return fmt.Errorf("select airdrops: %w", err)
	}
	if len(airdrops) == 0 {
		return nil
	}

	r.log.Debugf("Got %d pending airdrops, broadcasting now", len(airdrops))

	for _, drop := range airdrops {
		if err = r.handlePending(ctx, drop); err != nil {
			r.log.WithField("airdrop", drop).
				WithError(err).Error("Failed to handle pending airdrop")
			continue
		}
	}

	return nil
}

func (r *Runner) waitForTxErrors(ctx context.Context) {
	for {
		select {
		case errData := <-r.errChan:
			r.updateAirdropStatus(ctx, errData.airdrop.ID, errData.tx.Hash().String(), data.TxStatusFailed)
			return
		}
	}
}

func (r *Runner) handlePending(ctx context.Context, airdrop data.Airdrop) (err error) {
	var txHash string

	defer func() {
		r.UnlockNonce()
		if err != nil {
			r.updateAirdropStatus(ctx, airdrop.ID, txHash, data.TxStatusFailed)
		}
	}()

	r.LockNonce()
	r.updateAirdropStatus(ctx, airdrop.ID, txHash, data.TxStatusInProgress)

	tx, err := r.genTx(ctx, airdrop)
	if err != nil {
		return fmt.Errorf("failed to generate tx: %w", err)
	}

	if err = r.broadcastTx(ctx, tx, airdrop); err != nil {
		return fmt.Errorf("failed to broadcast tx: %w", err)
	}

	return nil
}

func (r *Runner) genTx(ctx context.Context, airdrop data.Airdrop) (*types.Transaction, error) {
	receiver := common.HexToAddress(airdrop.Address)
	txData := r.buildTransferTx(airdrop)

	gasPrice, gasLimit, err := r.getGasCosts(ctx, receiver, txData)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas costs: %w", err)
	}

	tx, err := types.SignNewTx(
		r.PrivateKey,
		types.NewCancunSigner(r.ChainID),
		&types.LegacyTx{
			Nonce:    r.Nonce(),
			Gas:      gasLimit,
			GasPrice: gasPrice,
			To:       &receiver,
			Data:     txData,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to sign new tx: %w", err)
	}

	return tx, nil
}

func (r *Runner) broadcastTx(ctx context.Context, tx *types.Transaction, airdrop data.Airdrop) error {
	if err := r.RPC.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("failed to send tx: %w", err)
	}

	r.IncrementNonce()
	r.waitForTransactionMined(ctx, tx, airdrop)

	return nil
}

func (r *Runner) waitForTransactionMined(ctx context.Context, transaction *types.Transaction, airdrop data.Airdrop) {
	log := r.log.WithField("tx", transaction.Hash().Hex())

	go func() {
		log.Debugf("waiting to mine")

		if _, err := bind.WaitMined(ctx, r.RPC, transaction); err != nil {
			log.WithError(err).Error("Failed to wait for mined tx")

			r.errChan <- &channelData{
				err:     err,
				tx:      transaction,
				airdrop: airdrop,
			}
		}

		r.updateAirdropStatus(ctx, airdrop.ID, transaction.Hash().String(), data.TxStatusCompleted)

		log.Debugf("was mined")
	}()
}

// If we don't update tx status from pending, having the successful funds
// transfer, it will be possible to double-spend. With this solution the
// double-spend may still occur, if the service is restarted before the
// successful update. There is a better solution with file creation on context
// cancellation and parsing it on start.
func (r *Runner) updateAirdropStatus(ctx context.Context, id, txHash, status string) {
	running.UntilSuccess(ctx, r.log, "tx-status-updater", func(_ context.Context) (bool, error) {
		var ptr *string
		if txHash != "" {
			ptr = &txHash
		}

		err := r.q.New().Update(id, map[string]any{
			"status":  status,
			"tx_hash": ptr,
		})

		return err == nil, err
	}, 2*time.Second, 10*time.Second)
}

func (r *Runner) buildTransferTx(airdrop data.Airdrop) []byte {
	methodID := hexutil.Encode(crypto.Keccak256([]byte(transferFnSignature))[:4])
	paddedAddress := common.LeftPadBytes(common.HexToAddress(airdrop.Address).Bytes(), byteSize)
	paddedAmount := common.LeftPadBytes(r.Amount.Bytes(), byteSize)

	var txData []byte
	txData = append(txData, methodID...)
	txData = append(txData, paddedAddress...)
	txData = append(txData, paddedAmount...)

	return txData
}

func (r *Runner) getGasCosts(
	ctx context.Context,
	receiver common.Address,
	txData []byte,
) (gasPrice *big.Int, gasLimit uint64, err error) {
	gasPrice, err = r.RPC.SuggestGasPrice(ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to suggest gas price")
	}

	gasLimit, err = r.RPC.EstimateGas(ctx, ethereum.CallMsg{
		To:       &receiver,
		GasPrice: gasPrice,
		Data:     txData,
	})
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to estimate gas limit")
	}

	return
}

package broadcaster

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/evm-airdrop-svc/contracts"
	"github.com/rarimo/evm-airdrop-svc/internal/config"
	"github.com/rarimo/evm-airdrop-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

type Runner struct {
	log *logan.Entry
	q   *data.AirdropsQ
	config.Broadcaster
	config.AirdropConfig
	erc20 *contracts.ERC20Permit
}

func Run(ctx context.Context, cfg *config.Config) {
	log := cfg.Log().WithField("service", "builtin-broadcaster")
	log.Info("Starting service")

	erc20Permit, err := contracts.NewERC20Permit(cfg.AirdropConfig().TokenAddress, cfg.Broadcaster().RPC)
	if err != nil {
		panic(errors.Wrap(err, "failed to init erc20 permit transfer contract"))
	}

	r := &Runner{
		log:           log,
		q:             data.NewAirdropsQ(cfg.DB().Clone()),
		Broadcaster:   cfg.Broadcaster(),
		AirdropConfig: cfg.AirdropConfig(),
		erc20:         erc20Permit,
	}

	running.WithBackOff(ctx, r.log, "builtin-broadcaster", r.run, 5*time.Second, 5*time.Second, 5*time.Second)
}

func (r *Runner) run(ctx context.Context) error {
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

func (r *Runner) handlePending(ctx context.Context, airdrop data.Airdrop) (err error) {
	var txHash string

	defer func() {
		r.UnlockNonce()
		if err != nil {
			r.updateAirdropStatus(ctx, airdrop.ID, txHash, data.TxStatusFailed, err)
		}
	}()

	r.LockNonce()
	r.updateAirdropStatus(ctx, airdrop.ID, txHash, data.TxStatusInProgress, nil)

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
	bigAmount, ok := new(big.Int).SetString(airdrop.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse amount: %s", airdrop.Amount)
	}

	tx, err := r.getTransferTx(ctx, common.HexToAddress(airdrop.Address), bigAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer params: %w", err)
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

		receipt, err := bind.WaitMined(ctx, r.RPC, transaction)
		if err != nil {
			log.WithError(err).WithField("transaction", transaction).Error("failed to wait for mined tx")
			r.updateAirdropStatus(ctx, airdrop.ID, transaction.Hash().String(), data.TxStatusFailed, err)
		}

		if receipt.Status != types.ReceiptStatusSuccessful {
			txErr, err := r.getTxError(ctx, transaction, r.Address)
			if err != nil {
				log.WithError(err).WithField("transaction", transaction).Error("failed to get tx error")
				r.updateAirdropStatus(ctx, airdrop.ID, transaction.Hash().String(), data.TxStatusFailed, err)
				return
			}

			log.WithError(err).WithField("transaction", transaction).Error("transaction was mined with failed status")
			r.updateAirdropStatus(ctx, airdrop.ID, transaction.Hash().String(), data.TxStatusFailed, txErr)
			return
		}

		r.updateAirdropStatus(ctx, airdrop.ID, transaction.Hash().String(), data.TxStatusCompleted, nil)

		log.Debugf("was mined successfully")
	}()
}

func (r *Runner) getTxError(ctx context.Context, tx *types.Transaction, txSender common.Address) (error, error) {
	msg := ethereum.CallMsg{
		From:     txSender,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	res, err := r.RPC.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make call")
	}

	return errors.New(string(res)), nil
}

// If we don't update tx status from pending, having the successful funds
// transfer, it will be possible to double-spend. With this solution the
// double-spend may still occur, if the service is restarted before the
// successful update. There is a better solution with file creation on context
// cancellation and parsing it on start.
func (r *Runner) updateAirdropStatus(ctx context.Context, id, txHash, status string, txErr error) {
	running.UntilSuccess(ctx, r.log, "tx-status-updater", func(_ context.Context) (bool, error) {
		var ptr *string
		if txHash != "" {
			ptr = &txHash
		}

		var errMsg *string
		if txErr != nil {
			msg := txErr.Error()
			errMsg = &msg
		}
		err := r.q.New().Update(id, map[string]any{
			"status":  status,
			"tx_hash": ptr,
			"error":   errMsg,
		})

		return err == nil, err
	}, 2*time.Second, 10*time.Second)
}

func (r *Runner) getTransferTx(
	ctx context.Context,
	receiver common.Address,
	amount *big.Int,
) (tx *types.Transaction, err error) {
	gasPrice, err := r.RPC.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}

	txOptions, err := bind.NewKeyedTransactorWithChainID(r.PrivateKey, r.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tx options")
	}
	txOptions.NoSend = true
	txOptions.Nonce = new(big.Int).SetUint64(r.Nonce())
	txOptions.GasPrice = gasPrice

	tx, err = r.erc20.Transfer(txOptions, receiver, amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to simulate transfer tx")
	}

	return
}

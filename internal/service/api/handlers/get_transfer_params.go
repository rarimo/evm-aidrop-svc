package handlers

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	pkgErrors "github.com/pkg/errors"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/models"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/requests"
	"github.com/rarimo/evm-airdrop-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var ErrInsufficienciesAmount = errors.New("amount is insufficient to pay tx fee")

type TransferTxParams struct {
	amount   *big.Int
	fee      *big.Int
	gasPrice *big.Int
	gasLimit uint64
	noSend   bool
}

func GetTransferParams(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewTransferERC20Token(r)
	if err != nil {
		api.Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	txParams, err := EstimateTransfer(r, req.Data.Attributes)
	if err != nil {
		api.Log(r).WithError(err).Error("failed to estimate transfer transaction")
		if pkgErrors.Is(err, ErrInsufficienciesAmount) {
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"data/attributes/amount": errors.From(err, logan.F{
					"amount": txParams.amount,
					"fee":    txParams.fee,
				}),
			})...)
			return
		}
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewEstimateResponse(txParams.amount, txParams.fee))
}

func EstimateTransfer(r *http.Request, attr resources.TransferErc20TokenAttributes) (*TransferTxParams, error) {
	halfAmount := new(big.Int).Div(attr.Amount, big.NewInt(2))

	tx, err := MakeTransferWithPermitTx(r, attr, TransferTxParams{
		noSend: true,
		fee:    halfAmount,
		amount: halfAmount,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to build transfer tx")
	}

	broadcaster := api.Broadcaster(r)
	gasPrice := broadcaster.MultiplyGasPrice(tx.GasPrice())
	feeAmount, err := buildFeeTransferAmount(r, gasPrice, tx.Gas())
	if err != nil {
		return nil, errors.Wrap(err, "failed to build fee transfer amount")
	}

	amount := new(big.Int).Sub(attr.Amount, feeAmount)

	params := TransferTxParams{
		amount:   amount,
		fee:      feeAmount,
		gasPrice: gasPrice,
		gasLimit: tx.Gas(),
	}

	if amount.Cmp(new(big.Int)) != 1 {
		return &params, ErrInsufficienciesAmount
	}

	return &params, nil
}

func MakeTransferWithPermitTx(
	r *http.Request,
	attr resources.TransferErc20TokenAttributes,
	params TransferTxParams,
) (*types.Transaction, error) {
	var (
		R [32]byte
		S [32]byte
	)

	txOptions, err := bind.NewKeyedTransactorWithChainID(api.Broadcaster(r).PrivateKey, api.Broadcaster(r).ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tx options")
	}
	txOptions.NoSend = params.noSend
	txOptions.GasPrice = params.gasPrice
	txOptions.GasLimit = params.gasLimit

	copy(R[:], hexutil.MustDecode(attr.R))
	copy(S[:], hexutil.MustDecode(attr.S))

	tx, err := api.ERC20PermitTransfer(r).TransferWithPermit(
		txOptions,
		api.AirdropConfig(r).TokenAddress,
		attr.Sender,
		attr.Receiver,
		params.amount,
		params.fee,
		attr.Deadline,
		attr.V,
		R,
		S,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build transfer with permit transaction")
	}

	return tx, nil
}

func buildFeeTransferAmount(r *http.Request, gweiGasPrice *big.Int, gasLimit uint64) (*big.Int, error) {
	dollarInEth, err := api.PriceAPIConfig(r).ConvertPrice()
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert dollar price in eth")
	}

	// Convert GWEI gas price to ETH gas price
	ethGasPrice := new(big.Float).Quo(new(big.Float).SetInt(gweiGasPrice), big.NewFloat(1e18))
	// Convert ETH gas price to dollar correspondence
	gasPriceInGlo := new(big.Float).Quo(ethGasPrice, dollarInEth)

	feeAmount := new(big.Float).Mul(new(big.Float).SetUint64(gasLimit), gasPriceInGlo)

	amount, _ := feeAmount.Int(nil)

	return amount, nil
}

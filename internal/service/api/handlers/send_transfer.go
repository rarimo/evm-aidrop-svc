package handlers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	pkgErrors "github.com/pkg/errors"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/models"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func SendTransfer(w http.ResponseWriter, r *http.Request) {
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

	txParams.noSend = false
	tx, err := MakeTransferWithPermitTx(r, req.Data.Attributes, *txParams)
	if err != nil {
		api.Log(r).WithError(err).Error("failed to build transfer transaction", logan.F{
			"params": txParams,
		})
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewTxResponse(txParams.amount, txParams.fee, tx.Hash().Hex()))
}

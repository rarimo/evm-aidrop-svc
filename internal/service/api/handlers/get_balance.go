package handlers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/models"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetBalanceRequest(r)
	if err != nil {
		api.Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	balance, err := api.ERC20Permit(r).BalanceOf(&bind.CallOpts{}, req.Address)
	if err != nil {
		api.Log(r).WithError(err).Error("failed to get user balance")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewBalanceResponse(req.Address.String(), balance))
}

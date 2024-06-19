package handlers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/models"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetTokenDetails(w http.ResponseWriter, r *http.Request) {
	name, err := api.ERC20Permit(r).Name(&bind.CallOpts{})
	if err != nil {
		api.Log(r).WithError(err).Error("failed to get token name")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	symbol, err := api.ERC20Permit(r).Symbol(&bind.CallOpts{})
	if err != nil {
		api.Log(r).WithError(err).Error("failed to get token symbol")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	decimals, err := api.ERC20Permit(r).Decimals(&bind.CallOpts{})
	if err != nil {
		api.Log(r).WithError(err).Error("failed to get token decimals")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewTokenDetailsResponse(
		api.AirdropConfig(r).TokenAddress.String(),
		name,
		symbol,
		api.AirdropConfig(r).TokenImage.String(),
		decimals,
	))
}

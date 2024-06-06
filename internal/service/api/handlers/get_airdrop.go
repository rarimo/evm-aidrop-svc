package handlers

import (
	"net/http"

	"github.com/rarimo/evm-airdrop-svc/internal/data"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/models"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetAirdrop(w http.ResponseWriter, r *http.Request) {
	nullifier, err := requests.NewGetAirdrop(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	airdrops, err := api.AirdropsQ(r).
		FilterByNullifier(nullifier).
		Select()
	if err != nil {
		api.Log(r).WithError(err).Error("Failed to select airdrops by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(airdrops) == 0 {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	airdrop := airdrops[0]
	for _, a := range airdrops[1:] {
		if a.Status == data.TxStatusCompleted {
			airdrop = a
			break
		}
	}

	ape.Render(w, models.NewAirdropResponse(airdrop))
}

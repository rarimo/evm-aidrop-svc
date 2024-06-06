package handlers

import (
	"net/http"

	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/models"
	"gitlab.com/distributed_lab/ape"
)

func GetAirdropParams(w http.ResponseWriter, r *http.Request) {
	ape.Render(w, models.NewAirdropParams(api.AirdropParams(r)))
}

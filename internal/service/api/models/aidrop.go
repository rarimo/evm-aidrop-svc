package models

import (
	"github.com/rarimo/evm-airdrop-svc/internal/config"
	"github.com/rarimo/evm-airdrop-svc/internal/data"
	"github.com/rarimo/evm-airdrop-svc/resources"
)

func NewAirdropResponse(tx data.Airdrop) resources.AirdropResponse {
	return resources.AirdropResponse{
		Data: resources.Airdrop{
			Key: resources.Key{
				ID:   tx.ID,
				Type: resources.AIRDROP,
			},
			Attributes: resources.AirdropAttributes{
				Nullifier: tx.Nullifier,
				Address:   tx.Address,
				TxHash:    tx.TxHash,
				Amount:    tx.Amount,
				Status:    tx.Status,
				CreatedAt: tx.CreatedAt,
				UpdatedAt: tx.UpdatedAt,
			},
		},
	}
}

func NewAirdropParams(params config.GlobalParams) resources.AirdropParamsResponse {
	return resources.AirdropParamsResponse{
		Data: resources.AirdropParams{
			Key: resources.Key{
				Type: resources.AIRDROP,
			},
			Attributes: resources.AirdropParamsAttributes{
				EventId:       params.EventID,
				StartedAt:     params.AirdropStart,
				QuerySelector: params.QuerySelector,
			},
		},
	}
}

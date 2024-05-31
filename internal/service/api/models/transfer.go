package models

import (
	"math/big"

	"github.com/rarimo/evm-airdrop-svc/resources"
)

func NewEstimateResponse(amount, fee *big.Int) resources.EstimateResponse {
	return resources.EstimateResponse{
		Data: resources.Estimate{
			Key: resources.Key{
				Type: resources.TRANSFER_ERC20,
			},
			Attributes: resources.EstimateAttributes{
				Amount: amount.Int64(),
				Fee:    fee.Int64(),
			},
		},
	}
}

func NewTxResponse(amount, fee *big.Int, hash string) resources.TxResponse {
	return resources.TxResponse{
		Data: resources.Tx{
			Key: resources.Key{
				ID:   hash,
				Type: resources.TRANSFER_ERC20,
			},
			Attributes: resources.TxAttributes{
				Amount: amount.Int64(),
				Fee:    fee.Int64(),
				Hash:   hash,
			},
		},
	}
}

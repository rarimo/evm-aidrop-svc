package models

import (
	"math/big"

	"github.com/rarimo/evm-airdrop-svc/resources"
)

func NewBalanceResponse(addr string, amount *big.Int) resources.BalanceResponse {
	return resources.BalanceResponse{
		Data: resources.Balance{
			Key: resources.Key{
				ID:   addr,
				Type: resources.TOKEN,
			},
			Attributes: resources.BalanceAttributes{
				Amount: amount,
			},
		},
	}
}

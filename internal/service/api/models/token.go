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

func NewTokenDetailsResponse(addr, name, symbol, image string, decimals uint8) resources.TokenDetailsResponse {
	return resources.TokenDetailsResponse{
		Data: resources.TokenDetails{
			Key: resources.Key{
				ID:   addr,
				Type: resources.TOKEN,
			},
			Attributes: resources.TokenDetailsAttributes{
				Name:     name,
				Symbol:   symbol,
				Decimals: decimals,
				Image:    image,
			},
		},
	}
}

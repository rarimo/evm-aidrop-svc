package models

import "github.com/rarimo/evm-airdrop-svc/resources"

func NewPermitHashResponse(hash string) resources.PermitHashResponse {
	return resources.PermitHashResponse{
		Data: resources.PermitHash{
			Key: resources.Key{
				Type: resources.TRANSFER_ERC20,
			},
			Attributes: resources.PermitHashAttributes{
				Hash: hash,
			},
		},
	}
}

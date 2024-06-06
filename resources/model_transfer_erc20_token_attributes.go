/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TransferErc20TokenAttributes struct {
	// Transferred amount of tokens.
	Amount *big.Int `json:"amount"`
	// UNIX UTC timestamp in the future till which permit signature may be used.
	Deadline *big.Int `json:"deadline"`
	// Hex encoded permit the x coordinate of R value of the signature.
	R string `json:"r"`
	// EVM address TO which tokens are transferred.
	Receiver common.Address `json:"receiver"`
	// Hex encoded permit the x coordinate of S value of the signature.
	S string `json:"s"`
	// EVM address FROM which tokens are transferred.
	Sender common.Address `json:"sender"`
	// The parity of the y coordinate of R.
	V uint8 `json:"v"`
}

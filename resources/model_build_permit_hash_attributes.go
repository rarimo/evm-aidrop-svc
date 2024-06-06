/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BuildPermitHashAttributes struct {
	// Transferred amount of tokens.
	Amount *big.Int `json:"amount"`
	// UNIX UTC timestamp in the future till which permit signature may be used.
	Deadline *big.Int `json:"deadline"`
	// EVM address FROM which tokens are transferred.
	Sender common.Address `json:"sender"`
}

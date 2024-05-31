/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type TransferErc20Token struct {
	Key
	Attributes TransferErc20TokenAttributes `json:"attributes"`
}
type TransferErc20TokenRequest struct {
	Data     TransferErc20Token `json:"data"`
	Included Included           `json:"included"`
}

type TransferErc20TokenListRequest struct {
	Data     []TransferErc20Token `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *TransferErc20TokenListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *TransferErc20TokenListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustTransferErc20Token - returns TransferErc20Token from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTransferErc20Token(key Key) *TransferErc20Token {
	var transferERC20Token TransferErc20Token
	if c.tryFindEntry(key, &transferERC20Token) {
		return &transferERC20Token
	}
	return nil
}

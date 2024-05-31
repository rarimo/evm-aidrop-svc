/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type Tx struct {
	Key
	Attributes TxAttributes `json:"attributes"`
}
type TxResponse struct {
	Data     Tx       `json:"data"`
	Included Included `json:"included"`
}

type TxListResponse struct {
	Data     []Tx            `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *TxListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *TxListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustTx - returns Tx from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTx(key Key) *Tx {
	var tx Tx
	if c.tryFindEntry(key, &tx) {
		return &tx
	}
	return nil
}

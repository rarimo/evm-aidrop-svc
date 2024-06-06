/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type PermitHash struct {
	Key
	Attributes PermitHashAttributes `json:"attributes"`
}
type PermitHashResponse struct {
	Data     PermitHash `json:"data"`
	Included Included   `json:"included"`
}

type PermitHashListResponse struct {
	Data     []PermitHash    `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *PermitHashListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *PermitHashListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustPermitHash - returns PermitHash from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPermitHash(key Key) *PermitHash {
	var permitHash PermitHash
	if c.tryFindEntry(key, &permitHash) {
		return &permitHash
	}
	return nil
}

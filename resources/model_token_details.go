/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type TokenDetails struct {
	Key
	Attributes TokenDetailsAttributes `json:"attributes"`
}
type TokenDetailsResponse struct {
	Data     TokenDetails `json:"data"`
	Included Included     `json:"included"`
}

type TokenDetailsListResponse struct {
	Data     []TokenDetails  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *TokenDetailsListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *TokenDetailsListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustTokenDetails - returns TokenDetails from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTokenDetails(key Key) *TokenDetails {
	var tokenDetails TokenDetails
	if c.tryFindEntry(key, &tokenDetails) {
		return &tokenDetails
	}
	return nil
}

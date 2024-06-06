/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type Estimate struct {
	Key
	Attributes EstimateAttributes `json:"attributes"`
}
type EstimateResponse struct {
	Data     Estimate `json:"data"`
	Included Included `json:"included"`
}

type EstimateListResponse struct {
	Data     []Estimate      `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *EstimateListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *EstimateListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustEstimate - returns Estimate from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEstimate(key Key) *Estimate {
	var estimate Estimate
	if c.tryFindEntry(key, &estimate) {
		return &estimate
	}
	return nil
}

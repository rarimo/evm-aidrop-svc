/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type BuildPermitHash struct {
	Key
	Attributes BuildPermitHashAttributes `json:"attributes"`
}
type BuildPermitHashRequest struct {
	Data     BuildPermitHash `json:"data"`
	Included Included        `json:"included"`
}

type BuildPermitHashListRequest struct {
	Data     []BuildPermitHash `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
	Meta     json.RawMessage   `json:"meta,omitempty"`
}

func (r *BuildPermitHashListRequest) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *BuildPermitHashListRequest) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustBuildPermitHash - returns BuildPermitHash from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBuildPermitHash(key Key) *BuildPermitHash {
	var buildPermitHash BuildPermitHash
	if c.tryFindEntry(key, &buildPermitHash) {
		return &buildPermitHash
	}
	return nil
}

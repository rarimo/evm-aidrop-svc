package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/evm-airdrop-svc/resources"
)

func NewCreateAirdrop(r *http.Request) (req resources.CreateAirdropRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, newDecodeError("body", err)
	}

	attr := req.Data.Attributes
	return req, val.Errors{
		"data/type": val.Validate(req.Data.Type, val.Required, val.In(resources.CREATE_AIRDROP)),
		"data/attributes/address": val.Validate(
			attr.Address,
			val.Required,
			val.Match(regexp.MustCompile("^0x[0-9a-fA-F]{40}$")),
		),
	}.Filter()
}

func newDecodeError(what string, err error) error {
	return val.Errors{
		what: fmt.Errorf("decode request %s: %w", what, err),
	}
}

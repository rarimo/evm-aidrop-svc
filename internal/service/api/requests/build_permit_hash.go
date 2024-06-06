package requests

import (
	"encoding/json"
	"net/http"

	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/evm-airdrop-svc/resources"
)

func NewBuildPermitHashRequest(r *http.Request) (req resources.BuildPermitHashRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, newDecodeError("body", err)
	}

	attr := req.Data.Attributes

	err = val.Errors{
		"data/type":                val.Validate(req.Data.Type, val.Required, val.In(resources.TRANSFER_ERC20)),
		"data/attributes/sender":   val.Validate(attr.Sender.String(), val.Required, val.Match(ethAddrRegExp)),
		"data/attributes/amount":   val.Validate(attr.Amount.Int64(), val.Required, val.Min(0)),
		"data/attributes/deadline": val.Validate(attr.Deadline.Int64(), val.Required, val.By(UnixTimestampRule)),
	}.Filter()
	if err != nil {
		return req, err
	}

	return req, nil
}

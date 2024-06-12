package requests

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type GetUserBalanceRequest struct {
	Address common.Address `url:"address"`
}

func NewGetBalanceRequest(r *http.Request) (GetUserBalanceRequest, error) {
	var req GetUserBalanceRequest

	if err := urlval.Decode(r.URL.Query(), &req); err != nil {
		return req, validation.Errors{
			"query": errors.Wrap(err, "failed to decode url"),
		}.Filter()
	}

	return req, req.validate()
}

func (r *GetUserBalanceRequest) validate() error {
	return validation.Errors{
		"/address": validation.Validate(r.Address, validation.Required),
	}.Filter()
}

package config

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const priceApiYamlKey = "price_api"

var (
	ErrPriceApiRequestFailed = errors.New("failed to fetch price api")
	ErrEmptyPrice            = errors.New("dollar price in ETH is empty")
)

type PriceApiConfiger interface {
	PriceApiConfig() PriceApiConfig
}

type PriceApiConfig struct {
	URL        *url.URL `fig:"url,required"`
	Key        string   `fig:"key,required"`
	CurrencyId string   `fig:"currency_id,required"`
	QuoteTag   string   `fig:"quote_tag,required"`
}

type priceApi struct {
	once   comfig.Once
	getter kv.Getter
}

func NewPriceApiConfiger(getter kv.Getter) PriceApiConfiger {
	return &priceApi{
		getter: getter,
	}
}

func (v *priceApi) PriceApiConfig() PriceApiConfig {
	return v.once.Do(func() interface{} {
		var result PriceApiConfig

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(v.getter, priceApiYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{
				"yaml_key": priceApiYamlKey,
			}))
		}

		return result
	}).(PriceApiConfig)
}

type QuoteResponse struct {
	Data map[string]Currency `json:"data"`
}

type Currency struct {
	Id     int              `json:"id"`
	Name   string           `json:"name"`
	Symbol string           `json:"symbol"`
	Quote  map[string]Quote `json:"quote"`
}

type Quote struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

// ConvertPrice converts tokens price
func (cfg PriceApiConfig) ConvertPrice() (*big.Float, error) {
	URL := cfg.URL.JoinPath("/v2/cryptocurrency/quotes/latest")

	query := URL.Query()
	query.Set("id", cfg.CurrencyId)
	query.Set("convert", cfg.QuoteTag)

	URL.RawQuery = query.Encode()

	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request", logan.F{
			"url": URL,
		})
	}

	request.Header.Set("X-CMC_PRO_API_KEY", cfg.Key)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	if response.StatusCode != http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read response body")
		}

		return nil, errors.From(ErrPriceApiRequestFailed, logan.F{
			"status": response.StatusCode,
			"body":   string(body),
		})
	}

	var body QuoteResponse
	if err = json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "failed to decode response body")
	}

	dollarInEth := big.NewFloat(body.Data[cfg.CurrencyId].Quote[cfg.QuoteTag].Price)
	if dollarInEth.Cmp(big.NewFloat(0)) == 0 {
		return nil, ErrEmptyPrice
	}

	return dollarInEth, nil
}

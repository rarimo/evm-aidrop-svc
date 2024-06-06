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

const priceAPIYamlKey = "price_api"

var (
	ErrPriceAPIRequestFailed = errors.New("failed to fetch price api")
	ErrEmptyPrice            = errors.New("dollar price in ETH is empty")
)

type PriceAPIConfiger interface {
	PriceAPIConfig() PriceAPIConfig
}

type PriceAPIConfig struct {
	URL        *url.URL `fig:"url,required"`
	Key        string   `fig:"key,required"`
	CurrencyID string   `fig:"currency_id,required"`
	QuoteTag   string   `fig:"quote_tag,required"`
}

type priceAPI struct {
	once   comfig.Once
	getter kv.Getter
}

func NewPriceAPIConfiger(getter kv.Getter) PriceAPIConfiger {
	return &priceAPI{
		getter: getter,
	}
}

func (v *priceAPI) PriceAPIConfig() PriceAPIConfig {
	return v.once.Do(func() interface{} {
		var result PriceAPIConfig

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(v.getter, priceAPIYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{
				"yaml_key": priceAPIYamlKey,
			}))
		}

		return result
	}).(PriceAPIConfig)
}

type QuoteResponse struct {
	Data map[string]Currency `json:"data"`
}

type Currency struct {
	ID     int              `json:"id"`
	Name   string           `json:"name"`
	Symbol string           `json:"symbol"`
	Quote  map[string]Quote `json:"quote"`
}

type Quote struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

// ConvertPrice converts tokens price
func (cfg PriceAPIConfig) ConvertPrice() (*big.Float, error) {
	URL := cfg.URL.JoinPath("/v2/cryptocurrency/quotes/latest")

	query := URL.Query()
	query.Set("id", cfg.CurrencyID)
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
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read response body")
		}

		return nil, errors.From(ErrPriceAPIRequestFailed, logan.F{
			"status": response.StatusCode,
			"body":   string(body),
		})
	}

	var body QuoteResponse
	if err = json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "failed to decode response body")
	}

	dollarInEth := big.NewFloat(body.Data[cfg.CurrencyID].Quote[cfg.QuoteTag].Price)
	if dollarInEth.Cmp(big.NewFloat(0)) == 0 {
		return nil, ErrEmptyPrice
	}

	return dollarInEth, nil
}

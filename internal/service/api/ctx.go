package api

import (
	"context"
	"net/http"

	"github.com/rarimo/evm-airdrop-svc/contracts"
	"github.com/rarimo/evm-airdrop-svc/internal/config"
	"github.com/rarimo/evm-airdrop-svc/internal/data"
	zk "github.com/rarimo/zkverifier-kit"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	airdropsQCtxKey
	airdropConfigCtxKey
	verifierCtxKey
	airdropParamsCtxKey
	broadcasterCtxKey
	erc20PermitCtxKey
	erc20PermitTransferCtxKey
	priceAPIConfigCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxAirdropsQ(q *data.AirdropsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, airdropsQCtxKey, q)
	}
}

func AirdropsQ(r *http.Request) *data.AirdropsQ {
	return r.Context().Value(airdropsQCtxKey).(*data.AirdropsQ).New()
}

func CtxAirdropConfig(entry config.AirdropConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, airdropConfigCtxKey, entry)
	}
}

func AirdropConfig(r *http.Request) config.AirdropConfig {
	return r.Context().Value(airdropConfigCtxKey).(config.AirdropConfig)
}

func CtxAirdropParams(params config.GlobalParams) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, airdropParamsCtxKey, params)
	}
}

func AirdropParams(r *http.Request) config.GlobalParams {
	return r.Context().Value(airdropParamsCtxKey).(config.GlobalParams)
}

func CtxVerifier(entry *zk.Verifier) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, verifierCtxKey, entry)
	}
}

func Verifier(r *http.Request) *zk.Verifier {
	return r.Context().Value(verifierCtxKey).(*zk.Verifier)
}

func Broadcaster(r *http.Request) config.Broadcaster {
	return r.Context().Value(broadcasterCtxKey).(config.Broadcaster)
}

func CtxBroadcaster(entry config.Broadcaster) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, broadcasterCtxKey, entry)
	}
}

func CtxERC20Permit(entry *contracts.ERC20Permit) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, erc20PermitCtxKey, entry)
	}
}

func ERC20Permit(r *http.Request) *contracts.ERC20Permit {
	return r.Context().Value(erc20PermitCtxKey).(*contracts.ERC20Permit)
}

func CtxERC20PermitTransfer(entry *contracts.ERC20TransferWithPermit) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, erc20PermitTransferCtxKey, entry)
	}
}

func ERC20PermitTransfer(r *http.Request) *contracts.ERC20TransferWithPermit {
	return r.Context().Value(erc20PermitTransferCtxKey).(*contracts.ERC20TransferWithPermit)
}

func CtxPriceAPIConfig(entry config.PriceAPIConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, priceAPIConfigCtxKey, entry)
	}
}

func PriceAPIConfig(r *http.Request) config.PriceAPIConfig {
	return r.Context().Value(priceAPIConfigCtxKey).(config.PriceAPIConfig)
}

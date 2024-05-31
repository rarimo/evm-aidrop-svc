package service

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/rarimo/evm-airdrop-svc/contracts"
	"github.com/rarimo/evm-airdrop-svc/internal/config"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func Run(ctx context.Context, cfg *config.Config) {
	r := chi.NewRouter()

	erc20Permit, err := contracts.NewERC20Permit(cfg.AirdropConfig().TokenAddress, cfg.Broadcaster().RPC)
	if err != nil {
		panic(errors.Wrap(err, "failed to init erc20 permit transfer contract"))
	}

	erc20PermitTransfer, err := contracts.NewERC20TransferWithPermit(
		cfg.Broadcaster().ERC20PermitTransfer, cfg.Broadcaster().RPC,
	)
	if err != nil {
		panic(errors.Wrap(err, "failed to init erc20 permit transfer contract"))
	}

	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleware(
			api.CtxLog(cfg.Log()),
			api.CtxVerifier(cfg.Verifier().ZkVerifier),
			api.CtxAirdropConfig(cfg.AirdropConfig()),
			api.CtxAirdropParams(cfg.Verifier().Params),
			api.CtxBroadcaster(cfg.Broadcaster()),
			api.CtxPriceAPIConfig(cfg.PriceAPIConfig()),
			api.CtxERC20Permit(erc20Permit),
			api.CtxERC20PermitTransfer(erc20PermitTransfer),
		),
		handlers.DBCloneMiddleware(cfg.DB()),
	)
	r.Route("/integrations/evm-airdrop-svc", func(r chi.Router) {
		r.Route("/airdrops", func(r chi.Router) {
			r.Post("/", handlers.CreateAirdrop)
			r.Get("/{nullifier}", handlers.GetAirdrop)
			r.Get("/params", handlers.GetAirdropParams)
		})

		r.Route("/transfer", func(r chi.Router) {
			r.Post("/", handlers.SendTransfer)
			r.Get("/", handlers.GetTransferParams)
		})
	})

	cfg.Log().Info("Service started")
	ape.Serve(ctx, r, cfg, ape.ServeOpts{})
}

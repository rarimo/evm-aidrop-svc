package config

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/dig"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const broadcasterYamlKey = "broadcaster"

type Broadcaster struct {
	RPC                 *ethclient.Client
	ChainID             *big.Int
	PrivateKey          *ecdsa.PrivateKey
	Address             common.Address
	QueryLimit          uint64
	ERC20PermitTransfer common.Address

	gasMultiplier float64
	nonce         uint64
	mut           *sync.Mutex
}

type Broadcasterer interface {
	Broadcaster() Broadcaster
}

type broadcasterer struct {
	getter kv.Getter
	once   comfig.Once
}

func NewBroadcaster(getter kv.Getter) Broadcasterer {
	return &broadcasterer{
		getter: getter,
	}
}

func (b *broadcasterer) Broadcaster() Broadcaster {
	return b.once.Do(func() interface{} {
		var cfg struct {
			RPC                 *ethclient.Client `fig:"rpc,required"`
			ChainID             *big.Int          `fig:"chain_id,required"`
			QueryLimit          uint64            `fig:"query_limit"`
			SenderPrivateKey    *ecdsa.PrivateKey `fig:"sender_private_key"`
			ERC20PermitTransfer common.Address    `fig:"erc20_permit_transfer,required"`
			GasMultiplier       float64           `fig:"gas_multiplier"`
		}

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks, figure.EthereumHooks).
			From(kv.MustGetStringMap(b.getter, broadcasterYamlKey)).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out broadcaster: %w", err))
		}

		if cfg.SenderPrivateKey == nil {
			cfg.SenderPrivateKey = extractPubKey()
		}

		queryLimit := uint64(100)
		if cfg.QueryLimit > 0 {
			queryLimit = cfg.QueryLimit
		}

		gasMultiplier := float64(1)
		if cfg.GasMultiplier > 0 {
			gasMultiplier = cfg.GasMultiplier
		}

		address := crypto.PubkeyToAddress(cfg.SenderPrivateKey.PublicKey)
		nonce, err := cfg.RPC.NonceAt(context.Background(), address, nil)
		if err != nil {
			panic(fmt.Errorf("failed to get nonce %w", err))
		}

		return Broadcaster{
			RPC:                 cfg.RPC,
			PrivateKey:          cfg.SenderPrivateKey,
			Address:             address,
			ChainID:             cfg.ChainID,
			QueryLimit:          queryLimit,
			ERC20PermitTransfer: cfg.ERC20PermitTransfer,

			gasMultiplier: gasMultiplier,
			nonce:         nonce,
			mut:           &sync.Mutex{},
		}
	}).(Broadcaster)
}

func extractPubKey() *ecdsa.PrivateKey {
	var envPK struct {
		PrivateKey *ecdsa.PrivateKey `dig:"PRIVATE_KEY,clear"`
	}

	if err := dig.Out(&envPK).With(figure.EthereumHooks).Now(); err != nil {
		panic(fmt.Errorf("failed to figure out private key from ENV: %w", err))
	}

	return envPK.PrivateKey
}

func (n *Broadcaster) LockNonce() {
	n.mut.Lock()
}

func (n *Broadcaster) UnlockNonce() {
	n.mut.Unlock()
}

func (n *Broadcaster) Nonce() uint64 {
	return n.nonce
}

func (n *Broadcaster) IncrementNonce() {
	n.nonce++
}

// ResetNonce sets nonce to the value received from a node
func (n *Broadcaster) ResetNonce(client *ethclient.Client) error {
	nonce, err := client.NonceAt(context.Background(), n.Address, nil)
	if err != nil {
		return fmt.Errorf("failed to get nonce, %w", err)
	}

	n.nonce = nonce
	return nil
}

func (n *Broadcaster) MultiplyGasPrice(gasPrice *big.Int) *big.Int {
	var ONE = 1000000000 // ONE - One GWEI
	mult := big.NewFloat(0).Mul(big.NewFloat(n.gasMultiplier), big.NewFloat(float64(ONE)))
	gas, _ := big.NewFloat(0).Mul(big.NewFloat(0).SetInt(gasPrice), mult).Int(nil)
	return big.NewInt(0).Div(gas, big.NewInt(int64(ONE)))
}

package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const airdropYamlKey = "airdrop"

type AirdropConfiger interface {
	AridropConfig() *AirdropConfig
}

type AirdropConfig struct {
	Amount       int64  `fig:"amount,required"`
	TokenAddress string `fig:"token_address,required"`
}

type airdrop struct {
	once   comfig.Once
	getter kv.Getter
}

func NewAirdropConfiger(getter kv.Getter) AirdropConfiger {
	return &airdrop{
		getter: getter,
	}
}

func (v *airdrop) AridropConfig() *AirdropConfig {
	return v.once.Do(func() interface{} {
		var result AirdropConfig

		err := figure.
			Out(&result).
			From(kv.MustGetStringMap(v.getter, airdropYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{
				"yaml_key": airdropYamlKey,
			}))
		}

		return &result
	}).(*AirdropConfig)
}

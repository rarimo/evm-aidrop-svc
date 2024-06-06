package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const routingYamlKey = "routing"

type Routing struct {
	Prefix string `fig:"prefix,required"`
}

func (c *Config) NewRouting() Routing {
	return c.routing.Do(func() interface{} {
		var cfg Routing

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, routingYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{"key": routingYamlKey}))
		}

		return cfg
	}).(Routing)
}

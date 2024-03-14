package server

import (
	"github.com/generativelabs/btcserver/internal/btc"
	"github.com/generativelabs/btcserver/internal/db"
)

type Config struct {
	Btc btc.Config `mapstructure:"btc"`

	Chakra struct {
		URL             string `mapstructure:"http-url"`
		ChainID         string `mapstructure:"chain-id"`
		PrivateKey      string `mapstructure:"private-key"`
		Address         string `mapstructure:"address"`
		ContractAddress string `mapstructure:"contract-address"`
	} `mapstructure:"chakra"`

	DB db.Config `mapstructure:"database"`

	ServicePort int `mapstructure:"service-port"`
}

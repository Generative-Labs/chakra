package headeroracle

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/generative-labs/relayer/internal/btc"
)

type HeaderOracle struct {
	btcRPCClient *rpcclient.Client
}

func New(cfg *Config) (*HeaderOracle, error) {
	btcRPCClient, err := btc.CrateBtcRPCClient(&cfg.BtcRPCConfig)
	if err != nil {
		return nil, err
	}

	headerOracle := &HeaderOracle{
		btcRPCClient: btcRPCClient,
	}

	return headerOracle, nil
}

func (ho HeaderOracle) Start() {
}

func (ho HeaderOracle) Stop() {
}

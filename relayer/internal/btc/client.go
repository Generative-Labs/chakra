package btc

import (
	"github.com/btcsuite/btcd/rpcclient"
)

func CrateBtcRPCClient(cfg *RPCConfig) (*rpcclient.Client, error) {
	return rpcclient.New(&rpcclient.ConnConfig{
		Host:         cfg.RPCHost,
		User:         cfg.RPCUser,
		Pass:         cfg.RPCPass,
		DisableTLS:   cfg.DisableTLS,
		HTTPPostMode: true,
	}, nil)
}

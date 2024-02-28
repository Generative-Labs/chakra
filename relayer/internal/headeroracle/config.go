package headeroracle

import "github.com/generative-labs/relayer/internal/btc"

type Config struct {
	BtcRPCConfig btc.RPCConfig `mapstructure:"btc-rpc-config"`
}

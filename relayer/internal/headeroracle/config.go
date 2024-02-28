package headeroracle

import (
	"github.com/generative-labs/relayer/internal/btc"
	"github.com/generative-labs/relayer/internal/chakra"
)

type Config struct {
	BtcRPC btc.RPCConfig `mapstructure:"btc-rpc"`

	Chakra struct {
		RPC chakra.RPCConfig `mapstructure:"rpc"`

		Wallet chakra.WalletConfig `mapstructure:"wallet"`
	} `mapstructure:"chakra"`

	Oracle OracleConfig `mapstructure:"oracle"`
}

type OracleConfig struct {
	// StartBlockHeight defines the block from which the BTC.
	StartBlockHeight int32 `mapstructure:"start-block-height"`

	// HeaderQueueContractAddress defines the address of the Btc Header queue contract on the Chakra network.
	HeaderQueueContractAddress string `mapstructure:"header-queue-contract-address"`

	// PollInterval defines the interval for querying whether BTC generates new block.
	PollInterval uint32 `mapstructure:"poll-interval"`

	// BlocksUntilFinalization defines the number of block confirmations required for a block to be considered finalized.
	BlocksUntilFinalization int `mapstructure:"blocks-until-finalization"`
}

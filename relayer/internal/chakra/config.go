package chakra

type RPCConfig struct {
	HTTPProvider string `mapstructure:"http-url"`
	ChainID      string `mapstructure:"chain-id"`
}

type WalletConfig struct {
	PrivateKey string `mapstructure:"private-key"`
}

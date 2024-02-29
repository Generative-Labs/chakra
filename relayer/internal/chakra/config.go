package chakra

type RPCConfig struct {
	URL     string `mapstructure:"http-url"`
	ChainID string `mapstructure:"chain-id"`
}

type WalletConfig struct {
	PrivateKey string `mapstructure:"private-key"`
}

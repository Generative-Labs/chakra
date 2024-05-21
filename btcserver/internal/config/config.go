package config

type Config struct {
	Btc BtcConfig `mapstructure:"btc"`

	Chakra struct {
		URL             string `mapstructure:"http-url"`
		ChainID         string `mapstructure:"chain-id"`
		PrivateKey      string `mapstructure:"private-key"`
		Address         string `mapstructure:"address"`
		ContractAddress string `mapstructure:"contract-address"`
	} `mapstructure:"chakra"`

	DB DBConfig `mapstructure:"database"`

	ServicePort int `mapstructure:"service-port"`

	Activity struct {
		Start string `mapstructure:"start"` // 2006-01-02 15:04:05
		End   string `mapstructure:"end"`   // 2006-01-02 15:04:05
	} `mapstructure:"activity"`
}

type BtcConfig struct {
	// NetworkName defines the bitcoin network name
	NetworkName string `mapstructure:"network-name"`
	// RPCHost defines the bitcoin rpc host
	RPCHost string `mapstructure:"rpc-host"`
	// RPCUser defines the bitcoin rpc user
	RPCUser string `mapstructure:"rpc-user"`
	// RPCPass defines the bitcoin rpc password
	RPCPass string `mapstructure:"rpc-pass"`
	// DisableTLS defines the bitcoin whether tls is required
	DisableTLS bool `mapstructure:"disable-tls"`
}

type DBConfig struct {
	Driver   string `mapstructure:"driver"`
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
	Password string `mapstructure:"password"`
}

package server

import (
	"context"
	"os"

	"github.com/NethermindEth/starknet.go/curve"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/generativelabs/btcserver/internal/api"
	"github.com/generativelabs/btcserver/internal/btc"
	"github.com/generativelabs/btcserver/internal/chakra"
	"github.com/generativelabs/btcserver/internal/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	// zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	viper.SetConfigType("yaml")

	viper.SetConfigName("btc-server.yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msgf("Fatal error config file: %s ", err)
	}
}

func Run() {
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error decode config into struct: %s ", err)
	}

	backend, err := db.CreateBackend(config.DB)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error create db backend: %s ", err)
	}

	btcClient, err := btc.NewClient(config.Btc)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error create btc client: %s ", err)
	}

	ctx := context.Background()
	cAccount, err := chakra.NewChakraAccount(ctx, config.Chakra.URL, config.Chakra.PrivateKey, config.Chakra.Address)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error new chakra account: %s ", err)
	}

	log.Info().Msgf("👷 Start to run btc server, conf: %+v", config)

	err = api.NewServer(ctx, backend, cAccount, config.Chakra.ContractAddress, btcClient).Run(config.ServicePort)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error in api server: %s ", err)
	}
}

func GetPublickeyFromPrivateKey(privateKey string) string {
	privInt := utils.HexToBN(privateKey)

	pubX, _, err := curve.Curve.PrivateToPoint(privInt)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error generate public key: %s ", err)
		panic(err)
	}

	pubKey := utils.BigToHex(pubX)

	return pubKey
}

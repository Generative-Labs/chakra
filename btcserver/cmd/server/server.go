package server

import (
	"context"
	"fmt"
	"github.com/NethermindEth/starknet.go/curve"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/generativelabs/btcserver/internal/api"
	"github.com/generativelabs/btcserver/internal/chakra"
	"github.com/generativelabs/btcserver/internal/db"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/rs/zerolog"
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

	backend, err := db.CreateBackend(config.Mysql)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error create db backend: %s ", err)
	}

	ctx := context.Background()
	provider, err := chakra.NewChakraProvider(ctx, config.Chakra.URL)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error new chakra provider: %s ", err)
	}

	pubkey := GetPublickeyFromPrivateKey(config.Chakra.PrivateKey)
	cAccount, err := chakra.NewChakraAccount(config.Chakra.PrivateKey, pubkey, config.Chakra.Address, provider)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error new chakra account: %s ", err)
	}

	err = api.New(ctx, backend, cAccount, config.Chakra.ContractAddress).Run(config.ServicePort)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error in api server: %s ", err)
	}
}

func GetPublickeyFromPrivateKey(privateKey string) string {
	privInt := utils.HexToBN(privateKey)

	pubX, _, err := curve.Curve.PrivateToPoint(privInt)
	if err != nil {
		fmt.Println("can't generate public key:", err)
		panic(err)
	}

	pubkey := utils.BigToHex(pubX)

	return pubkey
}

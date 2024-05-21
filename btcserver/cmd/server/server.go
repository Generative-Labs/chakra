package server

import (
	"context"
	"os"

	"github.com/generativelabs/btcserver/internal/api"
	"github.com/generativelabs/btcserver/internal/btc"
	"github.com/generativelabs/btcserver/internal/chakra"
	"github.com/generativelabs/btcserver/internal/config"
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

func Run(migrateFlag bool) {
	var conf config.Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error decode config into struct: %s ", err)
	}

	if migrateFlag {
		migrateDB(conf)
		return
	}

	backend, err := db.CreateBackend(conf.DB)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error create db backend: %s ", err)
	}

	btcClient, err := btc.NewClient(conf.Btc)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error create btc client: %s ", err)
	}

	ctx := context.Background()
	cAccount, err := chakra.NewChakraAccount(ctx, conf.Chakra.URL, conf.Chakra.PrivateKey, conf.Chakra.Address)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error new chakra account: %s ", err)
	}

	log.Info().Msgf("👷 Start to run btc server, conf: %+v", conf)

	api.InitActivityConfig(conf.Activity.Start, conf.Activity.End)

	err = api.NewServer(ctx, backend, cAccount, conf.Chakra.ContractAddress, btcClient).Run(conf.ServicePort)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error in api server: %s ", err)
	}
}

func migrateDB(conf config.Config) {
	log.Info().Msg("MigrateEntDB...")
	log.Info().Msgf("dbConf... %+v", conf)

	db.MigrateEntDB(conf)
	log.Info().Msg("Done")
}

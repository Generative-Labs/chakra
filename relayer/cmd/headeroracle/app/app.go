package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/generative-labs/relayer/internal/headeroracle"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	viper.SetConfigName("header-oracle")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msgf("Fatal error config file: %s ", err)
	}
}

func Run() {
	var config headeroracle.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Msgf("Fatal error decode config into struct: %s ", err)
	}

	headerOracle, err := headeroracle.New(&config)
	if err != nil {
		log.Fatal().Msgf("Fatal error create btc header oracle service failed: %s ", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Info().Msgf("Received signal %s, shutting down...\n", sig)

		headerOracle.Stop()
	}()

	headerOracle.Start()
}

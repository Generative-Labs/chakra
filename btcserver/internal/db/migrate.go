package db

import (
	"context"
	"github.com/generativelabs/btcserver/internal/config"
	"github.com/generativelabs/btcserver/internal/db/ent"
	"github.com/rs/zerolog/log"
)

func MigrateEntDB(conf config.Config) {
	switch conf.DB.Driver {
	case "mysql":
		MigrateEntMySQLDB(conf)
	default:
		MigrateEntSQLiteDB(conf)
	}
}

func MigrateEntSQLiteDB(conf config.Config) {
	client, err := ent.Open("sqlite3", conf.DB.Driver+".db?_fk=1")
	if err != nil {
		log.Fatal().Msgf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Warn().Msgf("failed creating schema resources: %v", err)
	}
}

func MigrateEntMySQLDB(conf config.Config) {
	client, err := CreateMysqlDB(conf.DB)
	if err != nil {
		log.Fatal().Msgf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
}

package db

import (
	"context"
	"fmt"

	"github.com/generativelabs/btcserver/internal/config"
	"github.com/generativelabs/btcserver/internal/db/ent"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/mattn/go-sqlite3"    // sqlite3 driver
	"github.com/rs/zerolog/log"
)

const (
	SqliteDriver string = "sqlite3"
	MysqlDriver  string = "mysql"
)

type Config struct {
	Driver   string `mapstructure:"driver"`
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
	Password string `mapstructure:"password"`
}

type Backend struct {
	dbClient *ent.Client
}

func CreateBackendWithDB(dbClient *ent.Client) *Backend {
	backend := Backend{
		dbClient: dbClient,
	}
	return &backend
}

func CreateBackend(config config.DBConfig) (*Backend, error) {
	var client *ent.Client
	var err error

	switch config.Driver {
	case MysqlDriver:
		client, err = CreateMysqlDB(config)
		if err != nil {
			return nil, err
		}

	case SqliteDriver:
		client, err = CreateSqliteDB(config.Database)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.Driver)
	}

	dbClient := Backend{
		dbClient: client,
	}

	return &dbClient, nil
}

func CreateMysqlDB(config config.DBConfig) (*ent.Client, error) {
	client, err := ent.Open(MysqlDriver, fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True",
		config.User, config.Password, config.Host, config.Database))
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Msgf("failed creating schema resources: %v", err)
	}

	return client, err
}

func CreateSqliteDB(dbName string) (*ent.Client, error) {
	client, err := ent.Open(SqliteDriver, dbName+".db?_fk=1")
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Msgf("failed creating schema resources: %v", err)
	}

	return client, err
}

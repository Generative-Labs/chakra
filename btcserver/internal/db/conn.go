package db

import (
	"fmt"

	"github.com/generativelabs/btcserver/internal/db/ent"
	_ "github.com/go-sql-driver/mysql" // mysql driver
)

const (
	SqliteDriver string = "sqlite3"
	MysqlDriver  string = "mysql"
)

type Config struct {
	Driver   string `mapstructure:"driver"`
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database	"`
	Password string `mapstructure:"password"`
}

type Backend struct {
	dbClient *ent.Client
}

func CreateBackend(config Config) (*Backend, error) {
	var client *ent.Client
	var err error

	switch config.Driver {
	case MysqlDriver:
		client, err = CreateMysqlDB(config)
		if err != nil {
			return nil, err
		}

	case SqliteDriver:
		client, err = CreateSqliteDB(config)
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

func CreateMysqlDB(config Config) (*ent.Client, error) {
	return ent.Open(MysqlDriver, fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True",
		config.User, config.Password, config.Host, config.Database))
}

func CreateSqliteDB(config Config) (*ent.Client, error) {
	return ent.Open(SqliteDriver, config.Database+".db?_fk=1")
}

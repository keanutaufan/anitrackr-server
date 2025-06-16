package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
)

func NewPostgresDatabase(config Config) *bun.DB {
	dsn := config.Dsn()
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	bunDb := bun.NewDB(sqlDb, pgdialect.New())
	bunDb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return bunDb
}

func LoadPostgresConfigFromEnv() Config {
	return Config{
		Protocol: "postgres",
		Host:     os.Getenv("DB_POSTGRES_HOSTNAME"),
		Port:     os.Getenv("DB_POSTGRES_PORT"),
		User:     os.Getenv("DB_POSTGRES_USER"),
		Password: os.Getenv("DB_POSTGRES_PASSWORD"),
		Database: os.Getenv("DB_POSTGRES_DATABASE"),
		SslMode:  os.Getenv("DB_POSTGRES_SSL_MODE"),
	}
}

package pgsql

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func Connect(ctx context.Context, cfg *Config) (*DB, error) {
	datasource := cfg.datasource()

	poolCfg, err := pgxpool.ParseConfig(datasource)
	if err != nil {
		return nil, errors.Wrap(err, "connect: could not parse pgxpool config")
	}

	pool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	return &DB{pool}, errors.Wrap(err, "could not connect with pgxpool")

}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c *Config) datasource() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

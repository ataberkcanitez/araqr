package pgsql

import (
	"database/sql"
	"embed"
	"github.com/cockroachdb/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"

	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrationsFs embed.FS

func MigrateUp(conf *Config) error {
	sourceDriver, err := iofs.New(migrationsFs, "migrations")
	if err != nil {
		return errors.Wrap(err, "failed to create source driver")
	}

	datasource := conf.datasource()

	instance, err := sql.Open("pgx", datasource)
	if err != nil {
		return errors.Wrap(err, "failed to open database")
	}

	dbDriver, err := pgx.WithInstance(instance, &pgx.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to create pgx driver")
	}

	migrator, err := migrate.NewWithInstance("iofs", sourceDriver, "pgx", dbDriver)
	if err != nil {
		return errors.Wrap(err, "failed to create migrator")
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "failed to migrate up")
	}
	return nil
}

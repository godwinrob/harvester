// Package migrate contains the database schema, migrations and seeding data.
package migrate

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"github.com/ardanlabs/darwin/v3"
	"log"

	"github.com/ardanlabs/darwin/v3/dialects/postgres"
	"github.com/ardanlabs/darwin/v3/drivers/generic"
	"github.com/godwinrob/harvester/business/sdk/sqldb"
	"github.com/jmoiron/sqlx"
)

var (
	//go:embed sql/migrate.sql
	migrateDoc string

	//go:embed sql/seed.sql
	seedDoc string
)

// Reset drops all tables and the darwin migrations table.
// WARNING: This will delete all data! Only use in development.
func Reset(ctx context.Context, db *sqlx.DB) error {
	if err := sqldb.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	// Drop tables in reverse dependency order
	queries := []string{
		"DROP TABLE IF EXISTS resources CASCADE",
		"DROP TABLE IF EXISTS galaxies CASCADE",
		"DROP TABLE IF EXISTS users CASCADE",
		"DROP TABLE IF EXISTS darwin_migrations CASCADE",
	}

	for _, query := range queries {
		if _, err := db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("drop table: %w", err)
		}
	}

	log.Println("Database reset complete - all tables dropped")
	return nil
}

// Migrate attempts to bring the database up to date with the migrations
// defined in this package.
func Migrate(ctx context.Context, db *sqlx.DB) error {
	if err := sqldb.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	driver, err := generic.New(db.DB, postgres.Dialect{})
	if err != nil {
		return fmt.Errorf("construct darwin driver: %w", err)
	}

	log.Printf("migrateDoc: %s", migrateDoc)
	d := darwin.New(driver, darwin.ParseMigrations(migrateDoc))
	return d.Migrate()
}

// Seed runs the seed document defined in this package against db. The queries
// are run in a transaction and rolled back if any fail.
func Seed(ctx context.Context, db *sqlx.DB) (err error) {
	if err := sqldb.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if errTx := tx.Rollback(); errTx != nil {
			if errors.Is(errTx, sql.ErrTxDone) {
				return
			}

			err = fmt.Errorf("rollback: %w", errTx)
			return
		}
	}()

	if _, err := tx.Exec(seedDoc); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}

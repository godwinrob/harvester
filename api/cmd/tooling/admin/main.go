package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/godwinrob/harvester/business/sdk/migrate"
	"github.com/godwinrob/harvester/business/sdk/sqldb"

	"github.com/jmoiron/sqlx"
)

func main() {
	slog.Info("migrate", "status", "starting migration in 5 seconds")
	time.Sleep(2 * time.Second)

	if err := Migrate(); err != nil {
		log.Fatalln(err)
	}

	slog.Info("migrate", "status", "migration completed")
	os.Exit(0)
}

func Migrate() error {
	slog.Info("migrate", "status", "beginning database migration")
	cfg := sqldb.Config{
		User:         "postgres",
		Password:     "postgres",
		Host:         "postgres",
		Name:         "postgres",
		MaxIdleConns: 0,
		MaxOpenConns: 0,
		DisableTLS:   true,
	}

	db, err := sqldb.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			slog.Error("close-db", "error", err)
		}
	}(db)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrate.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")

	if err := migrate.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}

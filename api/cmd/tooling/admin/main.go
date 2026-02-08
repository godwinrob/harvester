package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/godwinrob/harvester/business/sdk/migrate"
	"github.com/godwinrob/harvester/business/sdk/sqldb"
	"github.com/jmoiron/sqlx"
)

func main() {

	slog.Info("migrate", "status", "starting migration in 5 seconds")
	time.Sleep(8 * time.Second)

	if err := Migrate(); err != nil {
		slog.Error("error", err)
		os.Exit(1)
	}

	slog.Info("migrate", "status", "migration completed")
	os.Exit(0)
}

func Migrate() error {
	slog.Info("migrate", "status", "beginning database migration")

	cfg := sqldb.Config{
		User:         getEnv("HARVESTER_DB_USER", "postgres"),
		Password:     getEnv("HARVESTER_DB_PASSWORD", "postgres"),
		Host:         getEnv("HARVESTER_DB_HOST", "postgres"),
		Name:         getEnv("HARVESTER_DB_NAME", "postgres"),
		MaxIdleConns: 0,
		MaxOpenConns: 0,
		DisableTLS:   getEnv("HARVESTER_DB_DISABLE_TLS", "true") == "true",
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

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

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
		slog.Error("migration failed", "error", err)
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

	// Validate configuration before attempting to connect
	if err := validateConfig(cfg); err != nil {
		return err
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

	// Reset database to avoid checksum errors in development
	// WARNING: This drops all tables!
	resetDB := getEnv("HARVESTER_DB_RESET", "false") == "true"
	if resetDB {
		slog.Info("migrate", "status", "resetting database (all tables will be dropped)")
		if err := migrate.Reset(ctx, db); err != nil {
			return fmt.Errorf("reset database: %w", err)
		}
	}

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

// validateConfig checks if the database configuration contains placeholder values
// and fails fast with a clear error message instead of attempting to connect.
func validateConfig(cfg sqldb.Config) error {
	// Check for common placeholder patterns
	placeholders := map[string]string{
		"Host":     cfg.Host,
		"User":     cfg.User,
		"Password": cfg.Password,
		"Name":     cfg.Name,
	}

	for field, value := range placeholders {
		// Check for obvious placeholder patterns
		if value == "" ||
			value == "your_db_user" ||
			value == "your_db_name" ||
			value == "your_secure_password_here" ||
			value == "localhost_or_postgres_service" ||
			value == "CHANGE_ME" ||
			value == "TODO" {
			return fmt.Errorf(`
┌─────────────────────────────────────────────────────────────────┐
│ DATABASE CONFIGURATION ERROR (Migration)                        │
├─────────────────────────────────────────────────────────────────┤
│ The database %s is set to a placeholder value: "%s"
│                                                                 │
│ Please configure the database connection by setting:            │
│   HARVESTER_DB_HOST     - Database host (e.g., postgres)        │
│   HARVESTER_DB_USER     - Database username                     │
│   HARVESTER_DB_PASSWORD - Database password                     │
│   HARVESTER_DB_NAME     - Database name                         │
│                                                                 │
│ For local development with Docker Compose:                      │
│   These should be set in infrastructure/docker/.env             │
│                                                                 │
│ For Kubernetes:                                                 │
│   Update the Secret in infrastructure/k8s/base/harvester/       │
└─────────────────────────────────────────────────────────────────┘
`, field, value)
		}
	}

	return nil
}

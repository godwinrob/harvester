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
	if err := sqldb.ValidateConfig(cfg); err != nil {
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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
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

	// Seed resource type reference data first (groups, types, type-group mappings)
	// Must run before seed.sql because resources.resource_type has FK to resource_types
	slog.Info("seeding", "status", "inserting resource type reference data")
	if err := migrate.SeedAllResourceTypeData(ctx, db); err != nil {
		return fmt.Errorf("seed resource type data: %w", err)
	}
	fmt.Println("resource type seed data complete")

	if err := migrate.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")

	// Optionally seed random resources for development/testing
	seedResources := getEnv("HARVESTER_SEED_RESOURCES", "false") == "true"
	if seedResources {
		resourceCount := 1000 // Default to 1000 resources
		slog.Info("seeding", "status", "generating random resources", "count", resourceCount)

		if err := migrate.SeedRandomResources(ctx, db, resourceCount); err != nil {
			return fmt.Errorf("seed random resources: %w", err)
		}

		fmt.Printf("seeded %d random resources\n", resourceCount)
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

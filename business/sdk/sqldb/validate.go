package sqldb

import "fmt"

// ValidateConfig checks if the database configuration contains placeholder values
// and fails fast with a clear error message instead of attempting to connect.
func ValidateConfig(cfg Config) error {
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
│ DATABASE CONFIGURATION ERROR                                    │
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
│   Update the Secret in infrastructure/k8s/dev/harvester/        │
└─────────────────────────────────────────────────────────────────┘
`, field, value)
		}
	}

	return nil
}

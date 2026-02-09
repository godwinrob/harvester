# Database Migration and Seeding

This package handles database migrations and seeding for the Harvester application.

## Files

- `migrate.go` - Core migration and seeding logic
- `seed_resources.go` - Random resource generator for testing/development
- `sql/migrate.sql` - Database schema migrations
- `sql/seed.sql` - Initial seed data (users, galaxies, sample resources)

## Seeding Random Resources

The `SeedRandomResources` function generates realistic test data for development and testing environments.

### Features

- Generates unique resource names using Star Wars Galaxies-style naming
- Randomly assigns resources to existing galaxies and users from seed data
- Uses realistic resource types from Star Wars Galaxies
- Generates random stats (CR, CD, DR, FL, HR, MA, PE, OQ, SR, UT, ER) ranging from 0-1000
- Inserts data in batches of 500 for optimal performance

### Usage

Set the environment variable `HARVESTER_SEED_RESOURCES=true` to enable random resource seeding on startup.

```bash
# Development
export HARVESTER_SEED_RESOURCES=true
export HARVESTER_DB_RESET=true  # Optional: reset DB before seeding
./admin

# Docker
docker run -e HARVESTER_SEED_RESOURCES=true harvester/admin

# Kubernetes
# The dev environment already has this configured in dev-harvester-patch-configmap.yaml
kubectl apply -k infrastructure/k8s/dev/harvester/
```

### Default Configuration

- **Resource Count**: 1000 resources
- **Galaxies Used**: Finalizer, Bria (from seed.sql)
- **Users Used**: Luke Skywalker, Darth Vader (from seed.sql)
- **Resource Types**: 21 different Star Wars Galaxies resource types

### Resource Types

The seeder includes the following resource types:
- Class 1 Liquid Petro Fuel
- Kammris Iron
- Avalon Fiberplast
- Dolovite Steel
- And 17 more realistic types from Star Wars Galaxies

### Performance

- 1000 resources are inserted in approximately 2-3 seconds
- Uses batch inserts (500 records per batch) for optimal performance
- All inserts use `ON CONFLICT DO NOTHING` to prevent duplicate errors

## Development Workflow

1. Set environment variables:
   ```bash
   export HARVESTER_DB_RESET=true          # Drop all tables and recreate
   export HARVESTER_SEED_RESOURCES=true    # Generate 1000 random resources
   ```

2. Run the admin tool:
   ```bash
   cd api/cmd/tooling/admin
   go run main.go
   ```

3. Your database now has:
   - 2 users (Luke Skywalker, Darth Vader)
   - 2 galaxies (Finalizer, Bria)
   - 2 sample resources (from seed.sql)
   - 1000 randomly generated resources

## Production Usage

⚠️ **Important**: The `HARVESTER_SEED_RESOURCES` variable should NEVER be set to `true` in production environments. This feature is for development and testing only.

Production environments should only use the standard seed data from `sql/seed.sql`.

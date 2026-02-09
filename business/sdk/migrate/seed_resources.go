package migrate

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// resourceTypeKeys returns the resource type primary keys from the embedded TSV.
// These are the keys like "aluminum_agrinium" that match the resource_types table.
func resourceTypeKeys() []string {
	var keys []string
	lines := strings.Split(strings.TrimSpace(resourceTypesData), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Split(line, "\t")
		if len(fields) > 0 {
			key := strings.TrimSpace(fields[0])
			if key != "" {
				keys = append(keys, key)
			}
		}
	}
	return keys
}

// Resource name prefixes for variety
var namePrefixes = []string{
	"akim", "bavi", "celo", "dari", "efon", "feri", "galu", "havi",
	"ibro", "jami", "kelo", "lori", "manu", "nero", "onix", "pavi",
	"quem", "ruso", "soge", "tavi", "ubri", "vero", "wesi", "xeno",
	"yalu", "zemi", "aero", "byro", "cyri", "doro", "ekon", "falo",
}

var nameSuffixes = []string{
	"aic", "ium", "ian", "ite", "ese", "ari", "oni", "asi",
	"eri", "osi", "uri", "ali", "eli", "oli", "uli", "ami",
}

// SeedRandomResources generates and inserts random resources into the database.
// This is useful for testing and development with realistic data volumes.
func SeedRandomResources(ctx context.Context, db *sqlx.DB, count int) error {
	// Fixed galaxy and user IDs from seed.sql
	galaxyIDs := []string{
		"681672b7-95a8-4871-8832-e5774799c0e3", // Finalizer
		"b4629864-500c-4f06-8e4a-d31cea7bcfae", // Bria
	}

	userIDs := []string{
		"5cf37266-3473-4006-984f-9325122678b7", // Luke Skywalker
		"45b5fbd3-755f-4379-8f07-a58d4a30fa2f", // Darth Vader
	}

	// Build the list of valid resource type keys
	rtKeys := resourceTypeKeys()
	if len(rtKeys) == 0 {
		return fmt.Errorf("no resource type keys found in embedded data")
	}

	// Seed random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate resources
	type resourceData struct {
		resourceID        string
		name              string
		galaxyID          string
		userID            string
		resourceType      string
		unavailableAt     *time.Time
		unavailableUserID *string
		cr, cd, dr, fl, hr, ma, pe, oq, sr, ut, er int
	}

	var resources []resourceData
	usedNames := make(map[string]bool)

	for i := 0; i < count; i++ {
		// Generate a unique resource name
		var name string
		for {
			prefix := namePrefixes[r.Intn(len(namePrefixes))]
			suffix := nameSuffixes[r.Intn(len(nameSuffixes))]
			name = prefix + suffix

			// Add a number if name is already used
			if usedNames[name] {
				name = fmt.Sprintf("%s%d", name, r.Intn(1000))
			}

			if !usedNames[name] {
				usedNames[name] = true
				break
			}
		}

		// Randomly mark ~30% of resources as unavailable
		var unavailableAt *time.Time
		var unavailableUserID *string
		if r.Float32() < 0.3 {
			// Set unavailable date to sometime in the last 30 days
			daysAgo := r.Intn(30)
			ts := time.Now().AddDate(0, 0, -daysAgo)
			unavailableAt = &ts
			// Randomly pick a user who marked it unavailable
			uid := userIDs[r.Intn(len(userIDs))]
			unavailableUserID = &uid
		}

		resources = append(resources, resourceData{
			resourceID:        uuid.New().String(),
			name:              name,
			galaxyID:          galaxyIDs[r.Intn(len(galaxyIDs))],
			userID:            userIDs[r.Intn(len(userIDs))],
			resourceType:      rtKeys[r.Intn(len(rtKeys))],
			unavailableAt:     unavailableAt,
			unavailableUserID: unavailableUserID,
			cr:                r.Intn(1001),
			cd:                r.Intn(1001),
			dr:                r.Intn(1001),
			fl:                r.Intn(1001),
			hr:                r.Intn(1001),
			ma:                r.Intn(1001),
			pe:                r.Intn(1001),
			oq:                r.Intn(1001),
			sr:                r.Intn(1001),
			ut:                r.Intn(1001),
			er:                r.Intn(1001),
		})
	}

	// Prepare statement for batch insert
	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO resources (
			resource_id,
			resource_name,
			galaxy_id,
			added_user_id,
			resource_type,
			unavailable_at,
			unavailable_user_id,
			cr, cd, dr, fl, "hr", ma, pe, oq, sr, ut, er
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	// Insert resources one by one (prepared statements handle escaping)
	for i, res := range resources {
		_, err := stmt.ExecContext(ctx,
			res.resourceID,
			res.name,
			res.galaxyID,
			res.userID,
			res.resourceType,
			res.unavailableAt,
			res.unavailableUserID,
			res.cr, res.cd, res.dr, res.fl, res.hr,
			res.ma, res.pe, res.oq, res.sr, res.ut, res.er,
		)
		if err != nil {
			return fmt.Errorf("insert resource %d: %w", i, err)
		}
	}

	return nil
}

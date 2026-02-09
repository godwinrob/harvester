package migrate

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	//go:embed sql/seed/resource_groups.csv
	resourceGroupsData string

	//go:embed sql/seed/resource_types.tsv
	resourceTypesData string

	//go:embed sql/seed/resource_type_groups.csv
	resourceTypeGroupsData string
)

// SeedResourceGroups inserts the resource group hierarchy into the database.
func SeedResourceGroups(ctx context.Context, db *sqlx.DB) error {
	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO resource_groups (resource_group, group_name, group_level, group_order, container_type)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	lines := strings.Split(strings.TrimSpace(resourceGroupsData), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := parseCSVLine(line)
		if len(fields) < 5 {
			return fmt.Errorf("line %d: expected 5 fields, got %d", i+1, len(fields))
		}

		groupLevel, err := strconv.Atoi(fields[2])
		if err != nil {
			return fmt.Errorf("line %d: invalid group_level %q: %w", i+1, fields[2], err)
		}

		groupOrder, err := strconv.Atoi(fields[3])
		if err != nil {
			return fmt.Errorf("line %d: invalid group_order %q: %w", i+1, fields[3], err)
		}

		_, err = stmt.ExecContext(ctx, fields[0], fields[1], groupLevel, groupOrder, fields[4])
		if err != nil {
			return fmt.Errorf("insert resource_group %d (%s): %w", i+1, fields[0], err)
		}
	}

	return nil
}

// SeedResourceTypes inserts the resource type definitions into the database.
func SeedResourceTypes(ctx context.Context, db *sqlx.DB) error {
	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO resource_types (
			resource_type, resource_type_name, resource_category, resource_group,
			enterable, max_types,
			cr_min, cr_max, cd_min, cd_max, dr_min, dr_max, fl_min, fl_max,
			hr_min, hr_max, ma_min, ma_max, pe_min, pe_max, oq_min, oq_max,
			sr_min, sr_max, ut_min, ut_max, er_min, er_max,
			container_type, inventory_type, specific_planet
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11, $12, $13, $14,
			$15, $16, $17, $18, $19, $20, $21, $22,
			$23, $24, $25, $26, $27, $28,
			$29, $30, $31
		)
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	lines := strings.Split(strings.TrimSpace(resourceTypesData), "\n")
	total := len(lines)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if (i+1)%100 == 0 {
			log.Printf("  seeding resource types: %d/%d", i+1, total)
		}

		fields := strings.Split(line, "\t")
		if len(fields) < 31 {
			return fmt.Errorf("line %d: expected 31 fields, got %d", i+1, len(fields))
		}

		// Parse numeric fields
		nums := make([]int, 27)
		// fields[4]=enterable, [5]=maxTypes, [6..27]=stat min/max pairs, [30]=specificPlanet
		numFields := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 30}
		for j, fi := range numFields {
			val, err := strconv.Atoi(strings.TrimSpace(fields[fi]))
			if err != nil {
				return fmt.Errorf("line %d field %d: invalid number %q: %w", i+1, fi, fields[fi], err)
			}
			nums[j] = val
		}

		enterable := nums[0] == 1

		_, err = stmt.ExecContext(ctx,
			strings.TrimSpace(fields[0]),  // resource_type
			strings.TrimSpace(fields[1]),  // resource_type_name
			strings.TrimSpace(fields[2]),  // resource_category
			strings.TrimSpace(fields[3]),  // resource_group
			enterable,                     // enterable
			nums[1],                       // max_types
			nums[2], nums[3],              // cr_min, cr_max
			nums[4], nums[5],              // cd_min, cd_max
			nums[6], nums[7],              // dr_min, dr_max
			nums[8], nums[9],              // fl_min, fl_max
			nums[10], nums[11],            // hr_min, hr_max
			nums[12], nums[13],            // ma_min, ma_max
			nums[14], nums[15],            // pe_min, pe_max
			nums[16], nums[17],            // oq_min, oq_max
			nums[18], nums[19],            // sr_min, sr_max
			nums[20], nums[21],            // ut_min, ut_max
			nums[22], nums[23],            // er_min, er_max
			strings.TrimSpace(fields[28]), // container_type
			strings.TrimSpace(fields[29]), // inventory_type
			nums[24],                      // specific_planet
		)
		if err != nil {
			return fmt.Errorf("insert resource_type %d (%s): %w", i+1, fields[0], err)
		}
	}

	log.Printf("  seeding resource types: %d/%d done", total, total)
	return nil
}

// SeedResourceTypeGroups inserts the resource type to group mappings.
func SeedResourceTypeGroups(ctx context.Context, db *sqlx.DB) error {
	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO resource_type_groups (resource_type, resource_group)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	lines := strings.Split(strings.TrimSpace(resourceTypeGroupsData), "\n")
	total := len(lines)
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if (i+1)%500 == 0 {
			log.Printf("  seeding type-group mappings: %d/%d", i+1, total)
		}

		fields := parseCSVLine(line)
		if len(fields) < 2 {
			return fmt.Errorf("line %d: expected 2 fields, got %d", i+1, len(fields))
		}

		_, err = stmt.ExecContext(ctx, fields[0], fields[1])
		if err != nil {
			return fmt.Errorf("insert resource_type_group %d (%s, %s): %w", i+1, fields[0], fields[1], err)
		}
	}

	log.Printf("  seeding type-group mappings: %d/%d done", total, total)
	return nil
}

// SeedAllResourceTypeData seeds all resource type related tables in order.
func SeedAllResourceTypeData(ctx context.Context, db *sqlx.DB) error {
	if err := SeedResourceGroups(ctx, db); err != nil {
		return fmt.Errorf("seed resource groups: %w", err)
	}

	if err := SeedResourceTypes(ctx, db); err != nil {
		return fmt.Errorf("seed resource types: %w", err)
	}

	if err := SeedResourceTypeGroups(ctx, db); err != nil {
		return fmt.Errorf("seed resource type groups: %w", err)
	}

	return nil
}

// parseCSVLine parses a simple CSV line, handling quoted fields.
func parseCSVLine(line string) []string {
	var fields []string
	var current strings.Builder
	inQuotes := false

	for i := 0; i < len(line); i++ {
		ch := line[i]
		switch {
		case ch == '"':
			inQuotes = !inQuotes
		case ch == ',' && !inQuotes:
			fields = append(fields, current.String())
			current.Reset()
		default:
			current.WriteByte(ch)
		}
	}
	fields = append(fields, current.String())

	return fields
}

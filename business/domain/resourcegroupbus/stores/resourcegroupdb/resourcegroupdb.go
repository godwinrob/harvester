package resourcegroupdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/godwinrob/harvester/business/sdk/sqldb"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// Store manages the set of APIs for resource group database access.
type Store struct {
	log *logger.Logger
	db  sqlx.ExtContext
}

// NewStore constructs the api for data access.
func NewStore(log *logger.Logger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// Query retrieves a list of existing resource groups from the database.
func (s *Store) Query(ctx context.Context, filter resourcegroupbus.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]resourcegroupbus.ResourceGroup, error) {
	data := map[string]any{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		resource_group, group_name, group_level, group_order, container_type
	FROM
		resource_groups`

	buf := bytes.NewBufferString(q)
	applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbRes []resourceGroup
	if err := sqldb.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbRes); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	return toBusResourceGroups(dbRes), nil
}

// Count returns the total number of resource groups in the DB.
func (s *Store) Count(ctx context.Context, filter resourcegroupbus.QueryFilter) (int, error) {
	data := map[string]any{}

	const q = `
	SELECT
		count(1)
	FROM
		resource_groups`

	buf := bytes.NewBufferString(q)
	applyFilter(filter, data, buf)

	var count struct {
		Count int `db:"count"`
	}
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, buf.String(), data, &count); err != nil {
		return 0, fmt.Errorf("db: %w", err)
	}

	return count.Count, nil
}

// QueryByID gets the specified resource group from the database.
func (s *Store) QueryByID(ctx context.Context, resourceGroupKey string) (resourcegroupbus.ResourceGroup, error) {
	data := struct {
		ResourceGroup string `db:"resource_group"`
	}{
		ResourceGroup: resourceGroupKey,
	}

	const q = `
	SELECT
		resource_group, group_name, group_level, group_order, container_type
	FROM
		resource_groups
	WHERE
		resource_group = :resource_group`

	var dbRG resourceGroup
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbRG); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return resourcegroupbus.ResourceGroup{}, fmt.Errorf("db: %w", resourcegroupbus.ErrNotFound)
		}
		return resourcegroupbus.ResourceGroup{}, fmt.Errorf("db: %w", err)
	}

	return toBusResourceGroup(dbRG), nil
}

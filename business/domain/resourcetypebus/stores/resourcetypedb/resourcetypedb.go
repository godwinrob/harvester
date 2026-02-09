package resourcetypedb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/godwinrob/harvester/business/sdk/sqldb"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// Store manages the set of APIs for resource type database access.
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

// Create inserts a new resource type into the database.
func (s *Store) Create(ctx context.Context, rt resourcetypebus.ResourceType) error {
	const q = `
	INSERT INTO resource_types
		(resource_type, resource_type_name, resource_category, resource_group,
		 enterable, max_types,
		 cr_min, cr_max, cd_min, cd_max, dr_min, dr_max, fl_min, fl_max,
		 hr_min, hr_max, ma_min, ma_max, pe_min, pe_max, oq_min, oq_max,
		 sr_min, sr_max, ut_min, ut_max, er_min, er_max,
		 container_type, inventory_type, specific_planet)
	VALUES
		(:resource_type, :resource_type_name, :resource_category, :resource_group,
		 :enterable, :max_types,
		 :cr_min, :cr_max, :cd_min, :cd_max, :dr_min, :dr_max, :fl_min, :fl_max,
		 :hr_min, :hr_max, :ma_min, :ma_max, :pe_min, :pe_max, :oq_min, :oq_max,
		 :sr_min, :sr_max, :ut_min, :ut_max, :er_min, :er_max,
		 :container_type, :inventory_type, :specific_planet)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBResourceType(rt)); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", resourcetypebus.ErrUniqueType)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Update replaces a resource type document in the database.
func (s *Store) Update(ctx context.Context, rt resourcetypebus.ResourceType) error {
	const q = `
	UPDATE
		resource_types
	SET
		"resource_type_name" = :resource_type_name,
		"resource_category" = :resource_category,
		"resource_group" = :resource_group,
		"enterable" = :enterable,
		"max_types" = :max_types,
		"cr_min" = :cr_min, "cr_max" = :cr_max,
		"cd_min" = :cd_min, "cd_max" = :cd_max,
		"dr_min" = :dr_min, "dr_max" = :dr_max,
		"fl_min" = :fl_min, "fl_max" = :fl_max,
		"hr_min" = :hr_min, "hr_max" = :hr_max,
		"ma_min" = :ma_min, "ma_max" = :ma_max,
		"pe_min" = :pe_min, "pe_max" = :pe_max,
		"oq_min" = :oq_min, "oq_max" = :oq_max,
		"sr_min" = :sr_min, "sr_max" = :sr_max,
		"ut_min" = :ut_min, "ut_max" = :ut_max,
		"er_min" = :er_min, "er_max" = :er_max,
		"container_type" = :container_type,
		"inventory_type" = :inventory_type,
		"specific_planet" = :specific_planet
	WHERE
		resource_type = :resource_type`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBResourceType(rt)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a resource type from the database.
func (s *Store) Delete(ctx context.Context, rt resourcetypebus.ResourceType) error {
	const q = `
	DELETE FROM
		resource_types
	WHERE
		resource_type = :resource_type`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBResourceType(rt)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing resource types from the database.
func (s *Store) Query(ctx context.Context, filter resourcetypebus.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]resourcetypebus.ResourceType, error) {
	data := map[string]any{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		resource_type, resource_type_name, resource_category, resource_group,
		enterable, max_types,
		cr_min, cr_max, cd_min, cd_max, dr_min, dr_max, fl_min, fl_max,
		hr_min, hr_max, ma_min, ma_max, pe_min, pe_max, oq_min, oq_max,
		sr_min, sr_max, ut_min, ut_max, er_min, er_max,
		container_type, inventory_type, specific_planet
	FROM
		resource_types`

	buf := bytes.NewBufferString(q)
	applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbRes []resourceType
	if err := sqldb.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbRes); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	return toBusResourceTypes(dbRes), nil
}

// Count returns the total number of resource types in the DB.
func (s *Store) Count(ctx context.Context, filter resourcetypebus.QueryFilter) (int, error) {
	data := map[string]any{}

	const q = `
	SELECT
		count(1)
	FROM
		resource_types`

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

// QueryByID gets the specified resource type from the database.
func (s *Store) QueryByID(ctx context.Context, resourceTypeKey string) (resourcetypebus.ResourceType, error) {
	data := struct {
		ResourceType string `db:"resource_type"`
	}{
		ResourceType: resourceTypeKey,
	}

	const q = `
	SELECT
		resource_type, resource_type_name, resource_category, resource_group,
		enterable, max_types,
		cr_min, cr_max, cd_min, cd_max, dr_min, dr_max, fl_min, fl_max,
		hr_min, hr_max, ma_min, ma_max, pe_min, pe_max, oq_min, oq_max,
		sr_min, sr_max, ut_min, ut_max, er_min, er_max,
		container_type, inventory_type, specific_planet
	FROM
		resource_types
	WHERE
		resource_type = :resource_type`

	var dbRT resourceType
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbRT); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return resourcetypebus.ResourceType{}, fmt.Errorf("db: %w", resourcetypebus.ErrNotFound)
		}
		return resourcetypebus.ResourceType{}, fmt.Errorf("db: %w", err)
	}

	return toBusResourceType(dbRT), nil
}

// BulkCreate inserts multiple resource types into the database in a single transaction.
func (s *Store) BulkCreate(ctx context.Context, resourceTypes []resourcetypebus.ResourceType) error {
	db, ok := s.db.(*sqlx.DB)
	if !ok {
		return errors.New("bulk operations require *sqlx.DB")
	}

	return sqldb.WithTransaction(db, func(tx *sqlx.Tx) error {
		const q = `
		INSERT INTO resource_types
			(resource_type, resource_type_name, resource_category, resource_group,
			 enterable, max_types,
			 cr_min, cr_max, cd_min, cd_max, dr_min, dr_max, fl_min, fl_max,
			 hr_min, hr_max, ma_min, ma_max, pe_min, pe_max, oq_min, oq_max,
			 sr_min, sr_max, ut_min, ut_max, er_min, er_max,
			 container_type, inventory_type, specific_planet)
		VALUES
			(:resource_type, :resource_type_name, :resource_category, :resource_group,
			 :enterable, :max_types,
			 :cr_min, :cr_max, :cd_min, :cd_max, :dr_min, :dr_max, :fl_min, :fl_max,
			 :hr_min, :hr_max, :ma_min, :ma_max, :pe_min, :pe_max, :oq_min, :oq_max,
			 :sr_min, :sr_max, :ut_min, :ut_max, :er_min, :er_max,
			 :container_type, :inventory_type, :specific_planet)`

		for i, rt := range resourceTypes {
			if err := sqldb.NamedExecContextWithTx(ctx, s.log, tx, q, toDBResourceType(rt)); err != nil {
				if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
					return fmt.Errorf("item[%d]: %w", i, resourcetypebus.ErrUniqueType)
				}
				return fmt.Errorf("item[%d]: %w", i, err)
			}
		}
		return nil
	})
}

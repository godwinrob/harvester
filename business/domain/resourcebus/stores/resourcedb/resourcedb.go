package resourcedb

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/godwinrob/harvester/business/sdk/sqldb"
	"github.com/godwinrob/harvester/foundation/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Store manages the set of APIs for resource database access.
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

// Create inserts a new resource into the database.
func (s *Store) Create(ctx context.Context, res resourcebus.Resource) error {
	const q = `
	INSERT INTO resources
		(resource_id, resource_name, galaxy_id, added_at, updated_at, added_user_id, resource_type, cr, cd, dr, fl, "hr", ma, pe, oq, sr, ut, er)
	VALUES
		(:resource_id, :resource_name, :galaxy_id, :added_at, :updated_at, :added_user_id, :resource_type, :cr, :cd, :dr, :fl, :hr, :ma, :pe, :oq, :sr, :ut, :er)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBResource(res)); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", resourcebus.ErrUniqueName)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Update replaces a resource document in the database.
func (s *Store) Update(ctx context.Context, res resourcebus.Resource) error {
	const q = `
	UPDATE
		resources
	SET 
		"resource_name" = :resource_name,
		"unavailable_at" = :unavailable_at,
		"unavailable_user_id" = :unavailable_user_id,
		"verified" = :verified,
		"verified_user_id" = :verified_user_id,
		"cr" = :cr,
		"cd" = :cd,
		"dr" = :dr,
		"fl" = :fl,
		"hr" = :hr,
		"ma" = :ma,
		"pe" = :pe,
		"oq" = :oq,
		"sr" = :sr,
		"ut" = :ut,
		"er" = :er
	WHERE
		resource_id = :resource_id`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBResource(res)); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return resourcebus.ErrUniqueName
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a resource from the database.
func (s *Store) Delete(ctx context.Context, res resourcebus.Resource) error {
	const q = `
	DELETE FROM
		resources
	WHERE
		resource_id = :resource_id`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBResource(res)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing resources from the database.
func (s *Store) Query(ctx context.Context, filter resourcebus.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]resourcebus.Resource, error) {
	data := map[string]any{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		resource_id, resource_name, galaxy_id, added_at, added_user_id, resource_type, unavailable_at, unavailable_user_id, verified,verified_user_id, cr, cd, dr, fl, "hr", ma, pe, oq, sr, ut, er
	FROM
		resources`

	buf := bytes.NewBufferString(q)
	applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbRes []resource
	if err := sqldb.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbRes); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	return toBusResources(dbRes)
}

// Count returns the total number of resources in the DB.
func (s *Store) Count(ctx context.Context, filter resourcebus.QueryFilter) (int, error) {
	data := map[string]any{}

	const q = `
	SELECT
		count(1)
	FROM
		resources`

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

// QueryByID gets the specified resource from the database.
func (s *Store) QueryByID(ctx context.Context, resourceID uuid.UUID) (resourcebus.Resource, error) {
	data := struct {
		ID string `db:"resource_id"`
	}{
		ID: resourceID.String(),
	}

	const q = `
	SELECT
        resource_id, resource_name, galaxy_id, added_at, added_user_id, resource_type, unavailable_at, unavailable_user_id, verified,verified_user_id, cr, cd, dr, fl, "hr", ma, pe, oq, sr, ut, er
	FROM
		resources
	WHERE 
		resource_id = :resource_id`

	var dbRe resource
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbRe); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return resourcebus.Resource{}, fmt.Errorf("db: %w", resourcebus.ErrNotFound)
		}
		return resourcebus.Resource{}, fmt.Errorf("db: %w", err)
	}

	return toBusResource(dbRe)
}

// QueryByName gets the specified resource from the database by email.
func (s *Store) QueryByName(ctx context.Context, name string) (resourcebus.Resource, error) {
	data := struct {
		Name string `db:"resource_name"`
	}{
		Name: name,
	}

	const q = `
	SELECT
        resource_id, resource_name, galaxy_id, added_at, added_user_id, resource_type, unavailable_at, unavailable_user_id, verified,verified_user_id, cr, cd, dr, fl, "hr", ma, pe, oq, sr, ut, er
	FROM
		resources
	WHERE
		resource_name = :resource_name`

	var dbRe resource
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbRe); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return resourcebus.Resource{}, fmt.Errorf("db: %w", resourcebus.ErrNotFound)
		}
		return resourcebus.Resource{}, fmt.Errorf("db: %w", err)
	}

	return toBusResource(dbRe)
}

// BulkCreate inserts multiple resources into the database in a single transaction.
func (s *Store) BulkCreate(ctx context.Context, resources []resourcebus.Resource) error {
	db, ok := s.db.(*sqlx.DB)
	if !ok {
		return errors.New("bulk operations require *sqlx.DB")
	}

	return sqldb.WithTransaction(db, func(tx *sqlx.Tx) error {
		const q = `
		INSERT INTO resources
			(resource_id, resource_name, galaxy_id, added_at, updated_at, added_user_id, resource_type, cr, cd, dr, fl, "hr", ma, pe, oq, sr, ut, er)
		VALUES
			(:resource_id, :resource_name, :galaxy_id, :added_at, :updated_at, :added_user_id, :resource_type, :cr, :cd, :dr, :fl, :hr, :ma, :pe, :oq, :sr, :ut, :er)`

		for i, res := range resources {
			if err := sqldb.NamedExecContextWithTx(ctx, s.log, tx, q, toDBResource(res)); err != nil {
				if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
					return fmt.Errorf("item[%d]: %w", i, resourcebus.ErrUniqueName)
				}
				return fmt.Errorf("item[%d]: %w", i, err)
			}
		}
		return nil
	})
}

// BulkUpdate updates multiple resources in the database in a single transaction.
func (s *Store) BulkUpdate(ctx context.Context, resources []resourcebus.Resource) error {
	db, ok := s.db.(*sqlx.DB)
	if !ok {
		return errors.New("bulk operations require *sqlx.DB")
	}

	return sqldb.WithTransaction(db, func(tx *sqlx.Tx) error {
		const q = `
		UPDATE
			resources
		SET
			"resource_name" = :resource_name,
			"unavailable_at" = :unavailable_at,
			"unavailable_user_id" = :unavailable_user_id,
			"verified" = :verified,
			"verified_user_id" = :verified_user_id,
			"cr" = :cr,
			"cd" = :cd,
			"dr" = :dr,
			"fl" = :fl,
			"hr" = :hr,
			"ma" = :ma,
			"pe" = :pe,
			"oq" = :oq,
			"sr" = :sr,
			"ut" = :ut,
			"er" = :er,
			"updated_at" = :updated_at
		WHERE
			resource_id = :resource_id`

		for i, res := range resources {
			if err := sqldb.NamedExecContextWithTx(ctx, s.log, tx, q, toDBResource(res)); err != nil {
				if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
					return fmt.Errorf("item[%d]: %w", i, resourcebus.ErrUniqueName)
				}
				return fmt.Errorf("item[%d]: %w", i, err)
			}
		}
		return nil
	})
}

// BulkDelete removes multiple resources from the database in a single transaction.
func (s *Store) BulkDelete(ctx context.Context, ids []uuid.UUID) error {
	db, ok := s.db.(*sqlx.DB)
	if !ok {
		return errors.New("bulk operations require *sqlx.DB")
	}

	return sqldb.WithTransaction(db, func(tx *sqlx.Tx) error {
		data := struct {
			IDs []string `db:"ids"`
		}{
			IDs: make([]string, len(ids)),
		}
		for i, id := range ids {
			data.IDs[i] = id.String()
		}

		const q = `DELETE FROM resources WHERE resource_id IN (:ids)`

		if err := sqldb.NamedExecContextUsingInWithTx(ctx, s.log, tx, q, data); err != nil {
			return fmt.Errorf("namedexeccontextusingintx: %w", err)
		}
		return nil
	})
}

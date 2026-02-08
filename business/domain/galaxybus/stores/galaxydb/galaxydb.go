package galaxydb

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/godwinrob/harvester/business/sdk/sqldb"
	"github.com/godwinrob/harvester/foundation/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Store manages the set of APIs for galaxy database access.
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

// Create inserts a new galaxy into the database.
func (s *Store) Create(ctx context.Context, gal galaxybus.Galaxy) error {
	const q = `
	INSERT INTO galaxies
		(galaxy_id, galaxy_name, owner_user_id, enabled, date_created, date_updated)
	VALUES
		(:galaxy_id, :galaxy_name, :owner_user_id, :enabled, :date_created, :date_updated)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBGalaxy(gal)); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", galaxybus.ErrUniqueName)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Update replaces a galaxy document in the database.
func (s *Store) Update(ctx context.Context, gal galaxybus.Galaxy) error {
	const q = `
	UPDATE
		galaxies
	SET 
		"galaxy_name" = :galaxy_name,
		"owner_user_id" = :owner_user_id,
		"enabled" = :enabled,
		"date_updated" = :date_updated
	WHERE
		galaxy_id = :galaxy_id`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBGalaxy(gal)); err != nil {
		if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
			return galaxybus.ErrUniqueName
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a galaxy from the database.
func (s *Store) Delete(ctx context.Context, gal galaxybus.Galaxy) error {
	const q = `
	DELETE FROM
		galaxies
	WHERE
		galaxy_id = :galaxy_id`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBGalaxy(gal)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing galaxies from the database.
func (s *Store) Query(ctx context.Context, filter galaxybus.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]galaxybus.Galaxy, error) {
	data := map[string]any{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		galaxy_id, galaxy_name, owner_user_id, enabled, date_created, date_updated
	FROM
		galaxies`

	buf := bytes.NewBufferString(q)
	applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbUsrs []galaxy
	if err := sqldb.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbUsrs); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	return toBusGalaxies(dbUsrs)
}

// Count returns the total number of galaxies in the DB.
func (s *Store) Count(ctx context.Context, filter galaxybus.QueryFilter) (int, error) {
	data := map[string]any{}

	const q = `
	SELECT
		count(1)
	FROM
		galaxies`

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

// QueryByID gets the specified galaxy from the database.
func (s *Store) QueryByID(ctx context.Context, galaxyID uuid.UUID) (galaxybus.Galaxy, error) {
	data := struct {
		ID string `db:"galaxy_id"`
	}{
		ID: galaxyID.String(),
	}

	const q = `
	SELECT
        galaxy_id, galaxy_name, owner_user_id, enabled, date_created, date_updated
	FROM
		galaxies
	WHERE 
		galaxy_id = :galaxy_id`

	var dbUsr galaxy
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbUsr); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return galaxybus.Galaxy{}, fmt.Errorf("db: %w", galaxybus.ErrNotFound)
		}
		return galaxybus.Galaxy{}, fmt.Errorf("db: %w", err)
	}

	return toBusGalaxy(dbUsr)
}

// QueryByName gets the specified galaxy from the database.
func (s *Store) QueryByName(ctx context.Context, galaxyName string) (galaxybus.Galaxy, error) {
	data := struct {
		Name string `db:"galaxy_name"`
	}{
		Name: galaxyName,
	}

	const q = `
	SELECT
        galaxy_id, galaxy_name, owner_user_id, enabled, date_created, date_updated
	FROM
		galaxies
	WHERE
		galaxy_name = :galaxy_name`

	var dbGalaxy galaxy
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbGalaxy); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return galaxybus.Galaxy{}, fmt.Errorf("db: %w", galaxybus.ErrNotFound)
		}
		return galaxybus.Galaxy{}, fmt.Errorf("db: %w", err)
	}

	return toBusGalaxy(dbGalaxy)
}

// BulkCreate inserts multiple galaxies into the database in a single transaction.
func (s *Store) BulkCreate(ctx context.Context, galaxies []galaxybus.Galaxy) error {
	db, ok := s.db.(*sqlx.DB)
	if !ok {
		return errors.New("bulk operations require *sqlx.DB")
	}

	return sqldb.WithTransaction(db, func(tx *sqlx.Tx) error {
		const q = `
		INSERT INTO galaxies
			(galaxy_id, galaxy_name, owner_user_id, enabled, date_created, date_updated)
		VALUES
			(:galaxy_id, :galaxy_name, :owner_user_id, :enabled, :date_created, :date_updated)`

		for i, gal := range galaxies {
			if err := sqldb.NamedExecContextWithTx(ctx, s.log, tx, q, toDBGalaxy(gal)); err != nil {
				if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
					return fmt.Errorf("item[%d]: %w", i, galaxybus.ErrUniqueName)
				}
				return fmt.Errorf("item[%d]: %w", i, err)
			}
		}
		return nil
	})
}

// BulkUpdate updates multiple galaxies in the database in a single transaction.
func (s *Store) BulkUpdate(ctx context.Context, galaxies []galaxybus.Galaxy) error {
	db, ok := s.db.(*sqlx.DB)
	if !ok {
		return errors.New("bulk operations require *sqlx.DB")
	}

	return sqldb.WithTransaction(db, func(tx *sqlx.Tx) error {
		const q = `
		UPDATE
			galaxies
		SET
			"galaxy_name" = :galaxy_name,
			"owner_user_id" = :owner_user_id,
			"enabled" = :enabled,
			"date_updated" = :date_updated
		WHERE
			galaxy_id = :galaxy_id`

		for i, gal := range galaxies {
			if err := sqldb.NamedExecContextWithTx(ctx, s.log, tx, q, toDBGalaxy(gal)); err != nil {
				if errors.Is(err, sqldb.ErrDBDuplicatedEntry) {
					return fmt.Errorf("item[%d]: %w", i, galaxybus.ErrUniqueName)
				}
				return fmt.Errorf("item[%d]: %w", i, err)
			}
		}
		return nil
	})
}

// BulkDelete removes multiple galaxies from the database in a single transaction.
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

		const q = `DELETE FROM galaxies WHERE galaxy_id IN (:ids)`

		if err := sqldb.NamedExecContextUsingInWithTx(ctx, s.log, tx, q, data); err != nil {
			return fmt.Errorf("namedexeccontextusingintx: %w", err)
		}
		return nil
	})
}

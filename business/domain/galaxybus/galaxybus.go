package galaxybus

import (
	"context"
	"errors"
	"fmt"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/google/uuid"
	"time"

	"github.com/godwinrob/harvester/business/sdk/order"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("galaxy not found")
	ErrUniqueEmail           = errors.New("email is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, gal Galaxy) error
	Update(ctx context.Context, gal Galaxy) error
	Delete(ctx context.Context, gal Galaxy) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Galaxy, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, galaxyID uuid.UUID) (Galaxy, error)
	QueryByName(ctx context.Context, galaxyName string) (Galaxy, error)
}

// Business manages the set of APIs for galaxy access.
type Business struct {
	log    *logger.Logger
	storer Storer
}

// NewBusiness constructs a galaxy business API for use.
func NewBusiness(log *logger.Logger, storer Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// Create adds a new galaxy to the system.
func (b *Business) Create(ctx context.Context, nu NewGalaxy) (Galaxy, error) {
	now := time.Now()

	gal := Galaxy{
		ID:          uuid.New(),
		Name:        nu.Name,
		OwnerUserID: nu.OwnerUserID,
		Enabled:     true,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := b.storer.Create(ctx, gal); err != nil {
		return Galaxy{}, fmt.Errorf("create: %w", err)
	}

	return gal, nil
}

// Update modifies information about a galaxy.
func (b *Business) Update(ctx context.Context, gal Galaxy, uu UpdateGalaxy) (Galaxy, error) {
	if uu.Name != nil {
		gal.Name = *uu.Name
	}

	if uu.OwnerUserID != nil {
		gal.OwnerUserID = *uu.OwnerUserID
	}

	if uu.Enabled != nil {
		gal.Enabled = *uu.Enabled
	}
	gal.DateUpdated = time.Now()

	if err := b.storer.Update(ctx, gal); err != nil {
		return Galaxy{}, fmt.Errorf("update: %w", err)
	}

	return gal, nil
}

// Delete removes the specified galaxy.
func (b *Business) Delete(ctx context.Context, gal Galaxy) error {
	if err := b.storer.Delete(ctx, gal); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing galaxies.
func (b *Business) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Galaxy, error) {
	// TODO: DON'T ALLOW MORE THAN N RECORDS REGARDLESS

	galaxies, err := b.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return galaxies, nil
}

// Count returns the total number of galaxies.
func (b *Business) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return b.storer.Count(ctx, filter)
}

// QueryByID finds the galaxy by the specified Ib.
func (b *Business) QueryByID(ctx context.Context, galaxyID uuid.UUID) (Galaxy, error) {
	galaxy, err := b.storer.QueryByID(ctx, galaxyID)
	if err != nil {
		return Galaxy{}, fmt.Errorf("query: galaxyID[%s]: %w", galaxyID, err)
	}

	return galaxy, nil
}

// QueryByName finds the galaxy by the specified Ib.
func (b *Business) QueryByName(ctx context.Context, galaxyName string) (Galaxy, error) {
	galaxy, err := b.storer.QueryByName(ctx, galaxyName)
	if err != nil {
		return Galaxy{}, fmt.Errorf("query: galaxyName[%s]: %w", galaxyName, err)
	}

	return galaxy, nil
}

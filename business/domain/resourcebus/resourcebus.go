package resourcebus

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/godwinrob/harvester/foundation/logger"

	"github.com/google/uuid"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("resource not found")
	ErrUniqueEmail           = errors.New("email is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, res Resource) error
	Update(ctx context.Context, res Resource) error
	Delete(ctx context.Context, res Resource) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Resource, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, resourceID uuid.UUID) (Resource, error)
	QueryByName(ctx context.Context, resourceName string) (Resource, error)
}

// Business manages the set of APIs for resource access.
type Business struct {
	log    *logger.Logger
	storer Storer
}

// NewBusiness constructs a resource business API for use.
func NewBusiness(log *logger.Logger, storer Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// Create adds a new resource to the system.
func (b *Business) Create(ctx context.Context, nu NewResource) (Resource, error) {
	now := time.Now()

	res := Resource{
		ID:            uuid.New(),
		Name:          nu.Name,
		GalaxyID:      nu.GalaxyID,
		AddedAtDate:   now,
		UpdatedAtDate: now,
		AddedUserID:   nu.AddedUserID,
		ResourceType:  nu.ResourceType,
		CR:            nu.CR,
		CD:            nu.CD,
		DR:            nu.DR,
		FL:            nu.FL,
		HR:            nu.HR,
		MA:            nu.MA,
		PE:            nu.PE,
		OQ:            nu.OQ,
		SR:            nu.SR,
		UT:            nu.UT,
		ER:            nu.ER,
	}

	if err := b.storer.Create(ctx, res); err != nil {
		return Resource{}, fmt.Errorf("create: %w", err)
	}

	return res, nil
}

// Update modifies information about a resource.
func (b *Business) Update(ctx context.Context, res Resource, uu UpdateResource) (Resource, error) {
	if uu.Name != nil {
		res.Name = *uu.Name
	}

	if uu.UnavailableAt != nil {
		res.UnavailableAt = *uu.UnavailableAt
	}

	if uu.UnavailableUserID != nil {
		res.UnavailableUserID = *uu.UnavailableUserID
	}

	if uu.Verified != nil {
		res.Verified = *uu.Verified
	}

	if uu.VerifiedUserID != nil {
		res.VerifiedUserID = *uu.VerifiedUserID
	}

	if uu.CR != nil {
		res.CR = *uu.CR
	}

	if uu.CD != nil {
		res.CD = *uu.CD
	}

	if uu.DR != nil {
		res.DR = *uu.DR
	}

	if uu.FL != nil {
		res.FL = *uu.FL
	}

	if uu.HR != nil {
		res.HR = *uu.HR
	}

	if uu.MA != nil {
		res.MA = *uu.MA
	}

	if uu.PE != nil {
		res.PE = *uu.PE
	}

	if uu.OQ != nil {
		res.OQ = *uu.OQ
	}

	if uu.SR != nil {
		res.SR = *uu.SR
	}

	if uu.UT != nil {
		res.UT = *uu.UT
	}

	if uu.ER != nil {
		res.ER = *uu.ER
	}

	res.UnavailableAt = time.Now()

	if err := b.storer.Update(ctx, res); err != nil {
		return Resource{}, fmt.Errorf("update: %w", err)
	}

	return res, nil
}

// Delete removes the specified resource.
func (b *Business) Delete(ctx context.Context, res Resource) error {
	if err := b.storer.Delete(ctx, res); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing resources.
func (b *Business) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Resource, error) {
	// TODO: DON'T ALLOW MORE THAN N RECORDS REGARDLESS

	resources, err := b.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return resources, nil
}

// Count returns the total number of resources.
func (b *Business) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return b.storer.Count(ctx, filter)
}

// QueryByID finds the resource by the specified Ib.
func (b *Business) QueryByID(ctx context.Context, resourceID uuid.UUID) (Resource, error) {
	resource, err := b.storer.QueryByID(ctx, resourceID)
	if err != nil {
		return Resource{}, fmt.Errorf("query: resourceID[%s]: %w", resourceID, err)
	}

	return resource, nil
}

// QueryByName finds the resource by a specified resource email.
func (b *Business) QueryByName(ctx context.Context, name string) (Resource, error) {
	resource, err := b.storer.QueryByName(ctx, name)
	if err != nil {
		return Resource{}, fmt.Errorf("query: name[%s]: %w", name, err)
	}

	return resource, nil
}

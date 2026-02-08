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
	ErrUniqueName            = errors.New("resource name is not unique")
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
	BulkCreate(ctx context.Context, resources []Resource) error
	BulkUpdate(ctx context.Context, resources []Resource) error
	BulkDelete(ctx context.Context, ids []uuid.UUID) error
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

// BulkCreate adds multiple new resources to the system in a single transaction.
func (b *Business) BulkCreate(ctx context.Context, newResources []NewResource) ([]Resource, error) {
	resources := make([]Resource, len(newResources))
	now := time.Now()

	for i, nr := range newResources {
		resources[i] = Resource{
			ID:            uuid.New(),
			Name:          nr.Name,
			GalaxyID:      nr.GalaxyID,
			AddedAtDate:   now,
			UpdatedAtDate: now,
			AddedUserID:   nr.AddedUserID,
			ResourceType:  nr.ResourceType,
			CR:            nr.CR,
			CD:            nr.CD,
			DR:            nr.DR,
			FL:            nr.FL,
			HR:            nr.HR,
			MA:            nr.MA,
			PE:            nr.PE,
			OQ:            nr.OQ,
			SR:            nr.SR,
			UT:            nr.UT,
			ER:            nr.ER,
		}
	}

	if err := b.storer.BulkCreate(ctx, resources); err != nil {
		return nil, fmt.Errorf("bulkcreate: %w", err)
	}

	return resources, nil
}

// BulkUpdate modifies multiple resources in a single transaction.
func (b *Business) BulkUpdate(ctx context.Context, updates []UpdateResourceWithID) ([]Resource, error) {
	resources := make([]Resource, len(updates))

	for i, upd := range updates {
		res, err := b.storer.QueryByID(ctx, upd.ID)
		if err != nil {
			return nil, fmt.Errorf("querybyid[%d]: %w", i, err)
		}

		if upd.Data.Name != nil {
			res.Name = *upd.Data.Name
		}
		if upd.Data.UnavailableAt != nil {
			res.UnavailableAt = *upd.Data.UnavailableAt
		}
		if upd.Data.UnavailableUserID != nil {
			res.UnavailableUserID = *upd.Data.UnavailableUserID
		}
		if upd.Data.Verified != nil {
			res.Verified = *upd.Data.Verified
		}
		if upd.Data.VerifiedUserID != nil {
			res.VerifiedUserID = *upd.Data.VerifiedUserID
		}
		if upd.Data.CR != nil {
			res.CR = *upd.Data.CR
		}
		if upd.Data.CD != nil {
			res.CD = *upd.Data.CD
		}
		if upd.Data.DR != nil {
			res.DR = *upd.Data.DR
		}
		if upd.Data.FL != nil {
			res.FL = *upd.Data.FL
		}
		if upd.Data.HR != nil {
			res.HR = *upd.Data.HR
		}
		if upd.Data.MA != nil {
			res.MA = *upd.Data.MA
		}
		if upd.Data.PE != nil {
			res.PE = *upd.Data.PE
		}
		if upd.Data.OQ != nil {
			res.OQ = *upd.Data.OQ
		}
		if upd.Data.SR != nil {
			res.SR = *upd.Data.SR
		}
		if upd.Data.UT != nil {
			res.UT = *upd.Data.UT
		}
		if upd.Data.ER != nil {
			res.ER = *upd.Data.ER
		}
		res.UpdatedAtDate = time.Now()

		resources[i] = res
	}

	if err := b.storer.BulkUpdate(ctx, resources); err != nil {
		return nil, fmt.Errorf("bulkupdate: %w", err)
	}

	return resources, nil
}

// BulkDelete removes multiple resources in a single transaction.
func (b *Business) BulkDelete(ctx context.Context, ids []uuid.UUID) error {
	if err := b.storer.BulkDelete(ctx, ids); err != nil {
		return fmt.Errorf("bulkdelete: %w", err)
	}

	return nil
}

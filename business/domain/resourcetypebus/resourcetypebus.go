package resourcetypebus

import (
	"context"
	"errors"
	"fmt"

	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/godwinrob/harvester/foundation/logger"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound   = errors.New("resource type not found")
	ErrUniqueType = errors.New("resource type key is not unique")
)

// Storer interface declares the behavior this package needs to persist and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, rt ResourceType) error
	Update(ctx context.Context, rt ResourceType) error
	Delete(ctx context.Context, rt ResourceType) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]ResourceType, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, resourceType string) (ResourceType, error)
	BulkCreate(ctx context.Context, resourceTypes []ResourceType) error
}

// Business manages the set of APIs for resource type access.
type Business struct {
	log    *logger.Logger
	storer Storer
}

// NewBusiness constructs a resource type business API for use.
func NewBusiness(log *logger.Logger, storer Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// Create adds a new resource type to the system.
func (b *Business) Create(ctx context.Context, nu NewResourceType) (ResourceType, error) {
	rt := ResourceType{
		ResourceType:     nu.ResourceType,
		ResourceTypeName: nu.ResourceTypeName,
		ResourceCategory: nu.ResourceCategory,
		ResourceGroup:    nu.ResourceGroup,
		Enterable:        nu.Enterable,
		MaxTypes:         nu.MaxTypes,
		CRmin:            nu.CRmin,
		CRmax:            nu.CRmax,
		CDmin:            nu.CDmin,
		CDmax:            nu.CDmax,
		DRmin:            nu.DRmin,
		DRmax:            nu.DRmax,
		FLmin:            nu.FLmin,
		FLmax:            nu.FLmax,
		HRmin:            nu.HRmin,
		HRmax:            nu.HRmax,
		MAmin:            nu.MAmin,
		MAmax:            nu.MAmax,
		PEmin:            nu.PEmin,
		PEmax:            nu.PEmax,
		OQmin:            nu.OQmin,
		OQmax:            nu.OQmax,
		SRmin:            nu.SRmin,
		SRmax:            nu.SRmax,
		UTmin:            nu.UTmin,
		UTmax:            nu.UTmax,
		ERmin:            nu.ERmin,
		ERmax:            nu.ERmax,
		ContainerType:    nu.ContainerType,
		InventoryType:    nu.InventoryType,
		SpecificPlanet:   nu.SpecificPlanet,
	}

	if err := b.storer.Create(ctx, rt); err != nil {
		return ResourceType{}, fmt.Errorf("create: %w", err)
	}

	return rt, nil
}

// Update modifies information about a resource type.
func (b *Business) Update(ctx context.Context, rt ResourceType, uu UpdateResourceType) (ResourceType, error) {
	if uu.ResourceTypeName != nil {
		rt.ResourceTypeName = *uu.ResourceTypeName
	}
	if uu.ResourceCategory != nil {
		rt.ResourceCategory = *uu.ResourceCategory
	}
	if uu.ResourceGroup != nil {
		rt.ResourceGroup = *uu.ResourceGroup
	}
	if uu.Enterable != nil {
		rt.Enterable = *uu.Enterable
	}
	if uu.MaxTypes != nil {
		rt.MaxTypes = *uu.MaxTypes
	}
	if uu.CRmin != nil {
		rt.CRmin = *uu.CRmin
	}
	if uu.CRmax != nil {
		rt.CRmax = *uu.CRmax
	}
	if uu.CDmin != nil {
		rt.CDmin = *uu.CDmin
	}
	if uu.CDmax != nil {
		rt.CDmax = *uu.CDmax
	}
	if uu.DRmin != nil {
		rt.DRmin = *uu.DRmin
	}
	if uu.DRmax != nil {
		rt.DRmax = *uu.DRmax
	}
	if uu.FLmin != nil {
		rt.FLmin = *uu.FLmin
	}
	if uu.FLmax != nil {
		rt.FLmax = *uu.FLmax
	}
	if uu.HRmin != nil {
		rt.HRmin = *uu.HRmin
	}
	if uu.HRmax != nil {
		rt.HRmax = *uu.HRmax
	}
	if uu.MAmin != nil {
		rt.MAmin = *uu.MAmin
	}
	if uu.MAmax != nil {
		rt.MAmax = *uu.MAmax
	}
	if uu.PEmin != nil {
		rt.PEmin = *uu.PEmin
	}
	if uu.PEmax != nil {
		rt.PEmax = *uu.PEmax
	}
	if uu.OQmin != nil {
		rt.OQmin = *uu.OQmin
	}
	if uu.OQmax != nil {
		rt.OQmax = *uu.OQmax
	}
	if uu.SRmin != nil {
		rt.SRmin = *uu.SRmin
	}
	if uu.SRmax != nil {
		rt.SRmax = *uu.SRmax
	}
	if uu.UTmin != nil {
		rt.UTmin = *uu.UTmin
	}
	if uu.UTmax != nil {
		rt.UTmax = *uu.UTmax
	}
	if uu.ERmin != nil {
		rt.ERmin = *uu.ERmin
	}
	if uu.ERmax != nil {
		rt.ERmax = *uu.ERmax
	}
	if uu.ContainerType != nil {
		rt.ContainerType = *uu.ContainerType
	}
	if uu.InventoryType != nil {
		rt.InventoryType = *uu.InventoryType
	}
	if uu.SpecificPlanet != nil {
		rt.SpecificPlanet = *uu.SpecificPlanet
	}

	if err := b.storer.Update(ctx, rt); err != nil {
		return ResourceType{}, fmt.Errorf("update: %w", err)
	}

	return rt, nil
}

// Delete removes the specified resource type.
func (b *Business) Delete(ctx context.Context, rt ResourceType) error {
	if err := b.storer.Delete(ctx, rt); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing resource types.
func (b *Business) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]ResourceType, error) {
	rts, err := b.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return rts, nil
}

// Count returns the total number of resource types.
func (b *Business) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return b.storer.Count(ctx, filter)
}

// QueryByID finds the resource type by the specified key.
func (b *Business) QueryByID(ctx context.Context, resourceType string) (ResourceType, error) {
	rt, err := b.storer.QueryByID(ctx, resourceType)
	if err != nil {
		return ResourceType{}, fmt.Errorf("query: resourceType[%s]: %w", resourceType, err)
	}

	return rt, nil
}

// BulkCreate adds multiple new resource types to the system in a single transaction.
func (b *Business) BulkCreate(ctx context.Context, newTypes []NewResourceType) ([]ResourceType, error) {
	rts := make([]ResourceType, len(newTypes))

	for i, nu := range newTypes {
		rts[i] = ResourceType{
			ResourceType:     nu.ResourceType,
			ResourceTypeName: nu.ResourceTypeName,
			ResourceCategory: nu.ResourceCategory,
			ResourceGroup:    nu.ResourceGroup,
			Enterable:        nu.Enterable,
			MaxTypes:         nu.MaxTypes,
			CRmin:            nu.CRmin,
			CRmax:            nu.CRmax,
			CDmin:            nu.CDmin,
			CDmax:            nu.CDmax,
			DRmin:            nu.DRmin,
			DRmax:            nu.DRmax,
			FLmin:            nu.FLmin,
			FLmax:            nu.FLmax,
			HRmin:            nu.HRmin,
			HRmax:            nu.HRmax,
			MAmin:            nu.MAmin,
			MAmax:            nu.MAmax,
			PEmin:            nu.PEmin,
			PEmax:            nu.PEmax,
			OQmin:            nu.OQmin,
			OQmax:            nu.OQmax,
			SRmin:            nu.SRmin,
			SRmax:            nu.SRmax,
			UTmin:            nu.UTmin,
			UTmax:            nu.UTmax,
			ERmin:            nu.ERmin,
			ERmax:            nu.ERmax,
			ContainerType:    nu.ContainerType,
			InventoryType:    nu.InventoryType,
			SpecificPlanet:   nu.SpecificPlanet,
		}
	}

	if err := b.storer.BulkCreate(ctx, rts); err != nil {
		return nil, fmt.Errorf("bulkcreate: %w", err)
	}

	return rts, nil
}

package resourcegroupbus

import (
	"context"
	"errors"
	"fmt"

	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/godwinrob/harvester/foundation/logger"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound = errors.New("resource group not found")
)

// Storer interface declares the behavior this package needs to persist and
// retrieve data.
type Storer interface {
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]ResourceGroup, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, resourceGroup string) (ResourceGroup, error)
}

// Business manages the set of APIs for resource group access.
type Business struct {
	log    *logger.Logger
	storer Storer
}

// NewBusiness constructs a resource group business API for use.
func NewBusiness(log *logger.Logger, storer Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// Query retrieves a list of existing resource groups.
func (b *Business) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]ResourceGroup, error) {
	groups, err := b.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return groups, nil
}

// Count returns the total number of resource groups.
func (b *Business) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return b.storer.Count(ctx, filter)
}

// QueryByID finds the resource group by the specified key.
func (b *Business) QueryByID(ctx context.Context, resourceGroup string) (ResourceGroup, error) {
	group, err := b.storer.QueryByID(ctx, resourceGroup)
	if err != nil {
		return ResourceGroup{}, fmt.Errorf("query: resourceGroup[%s]: %w", resourceGroup, err)
	}

	return group, nil
}

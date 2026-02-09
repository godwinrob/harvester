package resourcetypeapp

import (
	"context"
	"errors"

	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/app/sdk/page"
	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

// App manages the set of app layer api functions for the resource type domain.
type App struct {
	resourceTypeBus *resourcetypebus.Business
}

// NewApp constructs a resource type app API for use.
func NewApp(resourceTypeBus *resourcetypebus.Business) *App {
	return &App{
		resourceTypeBus: resourceTypeBus,
	}
}

// Create adds a new resource type to the system.
func (a *App) Create(ctx context.Context, app NewResourceType) (ResourceType, error) {
	nr := toBusNewResourceType(app)

	rt, err := a.resourceTypeBus.Create(ctx, nr)
	if err != nil {
		if errors.Is(err, resourcetypebus.ErrUniqueType) {
			return ResourceType{}, errs.New(errs.Aborted, resourcetypebus.ErrUniqueType)
		}
		return ResourceType{}, errs.Newf(errs.Internal, "create: rt[%+v]: %s", rt, err)
	}

	return toAppResourceType(rt), nil
}

// Update updates an existing resource type.
func (a *App) Update(ctx context.Context, resourceTypeKey string, app UpdateResourceType) (ResourceType, error) {
	uu := toBusUpdateResourceType(app)

	rt, err := a.resourceTypeBus.QueryByID(ctx, resourceTypeKey)
	if err != nil {
		return ResourceType{}, errs.Newf(errs.Internal, "resource type missing: %s", err)
	}

	updRT, err := a.resourceTypeBus.Update(ctx, rt, uu)
	if err != nil {
		return ResourceType{}, errs.Newf(errs.Internal, "update: resourceType[%s]: %s", resourceTypeKey, err)
	}

	return toAppResourceType(updRT), nil
}

// Delete removes a resource type from the system.
func (a *App) Delete(ctx context.Context, resourceTypeKey string) error {
	rt, err := a.resourceTypeBus.QueryByID(ctx, resourceTypeKey)
	if err != nil {
		return errs.Newf(errs.Internal, "resource type missing: %s", err)
	}

	if err := a.resourceTypeBus.Delete(ctx, rt); err != nil {
		return errs.Newf(errs.Internal, "delete: resourceType[%s]: %s", resourceTypeKey, err)
	}

	return nil
}

// Query returns a list of resource types with paging.
func (a *App) Query(ctx context.Context, qp QueryParams) (page.Document[ResourceType], error) {
	pg, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return page.Document[ResourceType]{}, err
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return page.Document[ResourceType]{}, err
	}

	orderBy, err := order.Parse(orderByFields, qp.OrderBy, defaultOrderBy)
	if err != nil {
		return page.Document[ResourceType]{}, err
	}

	rts, err := a.resourceTypeBus.Query(ctx, filter, orderBy, pg.Number, pg.RowsPerPage)
	if err != nil {
		return page.Document[ResourceType]{}, errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.resourceTypeBus.Count(ctx, filter)
	if err != nil {
		return page.Document[ResourceType]{}, errs.Newf(errs.Internal, "count: %s", err)
	}

	return page.NewDocument(toAppResourceTypes(rts), total, pg.Number, pg.RowsPerPage), nil
}

// QueryByID returns a resource type by its key.
func (a *App) QueryByID(ctx context.Context, resourceTypeKey string) (ResourceType, error) {
	rt, err := a.resourceTypeBus.QueryByID(ctx, resourceTypeKey)
	if err != nil {
		return ResourceType{}, errs.Newf(errs.Internal, "resource type missing: %s", err)
	}

	return toAppResourceType(rt), nil
}

// BulkCreate adds multiple new resource types to the system.
func (a *App) BulkCreate(ctx context.Context, app BulkNewResourceTypes) (BulkResourceTypes, error) {
	var bulkErrors []errs.BulkItemError
	newTypes := make([]resourcetypebus.NewResourceType, 0, len(app.Items))

	for i, item := range app.Items {
		if err := item.Validate(); err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "item",
				Error: err.Error(),
			})
			continue
		}
		newTypes = append(newTypes, toBusNewResourceType(item))
	}

	if len(bulkErrors) > 0 {
		return BulkResourceTypes{}, errs.NewBulkValidationError(bulkErrors)
	}

	rts, err := a.resourceTypeBus.BulkCreate(ctx, newTypes)
	if err != nil {
		if errors.Is(err, resourcetypebus.ErrUniqueType) {
			return BulkResourceTypes{}, errs.New(errs.Aborted, resourcetypebus.ErrUniqueType)
		}
		return BulkResourceTypes{}, errs.Newf(errs.Internal, "bulkcreate: %s", err)
	}

	return BulkResourceTypes{
		Items:   toAppResourceTypes(rts),
		Created: len(rts),
	}, nil
}

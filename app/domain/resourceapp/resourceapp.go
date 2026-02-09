package resourceapp

import (
	"context"
	"errors"
	"github.com/godwinrob/harvester/app/sdk/bulk"
	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/app/sdk/page"
	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/google/uuid"
)

// App manages the set of app layer api functions for the resource domain.
type App struct {
	resourceBus *resourcebus.Business
}

// NewApp constructs a resource app API for use.
func NewApp(resourceBus *resourcebus.Business) *App {
	return &App{
		resourceBus: resourceBus,
	}
}

// NewAppWithAuth constructs a resource app API for use with auth support.
func NewAppWithAuth(resourceBus *resourcebus.Business) *App {
	return &App{
		resourceBus: resourceBus,
	}
}

// Create adds a new resource to the system.
func (a *App) Create(ctx context.Context, app NewResource) (Resource, error) {
	nc, err := toBusNewResource(app)
	if err != nil {
		return Resource{}, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.resourceBus.Create(ctx, nc)
	if err != nil {
		if errors.Is(err, resourcebus.ErrUniqueName) {
			return Resource{}, errs.New(errs.Aborted, resourcebus.ErrUniqueName)
		}
		return Resource{}, errs.Newf(errs.Internal, "create: usr[%+v]: %s", usr, err)
	}

	return toAppResource(usr), nil
}

// Update updates an existing resource.
func (a *App) Update(ctx context.Context, resourceID string, app UpdateResource) (Resource, error) {
	uu, err := toBusUpdateResource(app)
	if err != nil {
		return Resource{}, errs.New(errs.FailedPrecondition, err)
	}

	id, err := uuid.Parse(resourceID)
	if err != nil {
		return Resource{}, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.resourceBus.QueryByID(ctx, id)
	if err != nil {
		return Resource{}, errs.Newf(errs.Internal, "resource missing in context: %s", err)
	}

	updUsr, err := a.resourceBus.Update(ctx, usr, uu)
	if err != nil {
		return Resource{}, errs.Newf(errs.Internal, "update: resourceID[%s] uu[%+v]: %s", usr.ID, uu, err)
	}

	return toAppResource(updUsr), nil
}

// Delete removes a resource from the system.
func (a *App) Delete(ctx context.Context, resourceID string) error {
	id, err := uuid.Parse(resourceID)
	if err != nil {
		return errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.resourceBus.QueryByID(ctx, id)
	if err != nil {
		return errs.Newf(errs.Internal, "resource missing in context: %s", err)
	}

	if err := a.resourceBus.Delete(ctx, usr); err != nil {
		return errs.Newf(errs.Internal, "delete: resourceID[%s]: %s", usr.ID, err)
	}

	return nil
}

// Query returns a list of resources with paging.
func (a *App) Query(ctx context.Context, qp QueryParams) (page.Document[Resource], error) {
	pg, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return page.Document[Resource]{}, err
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return page.Document[Resource]{}, err
	}

	orderBy, err := order.Parse(orderByFields, qp.OrderBy, defaultOrderBy)
	if err != nil {
		return page.Document[Resource]{}, err
	}

	usrs, err := a.resourceBus.Query(ctx, filter, orderBy, pg.Number, pg.RowsPerPage)
	if err != nil {
		return page.Document[Resource]{}, errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.resourceBus.Count(ctx, filter)
	if err != nil {
		return page.Document[Resource]{}, errs.Newf(errs.Internal, "count: %s", err)
	}

	return page.NewDocument(toAppResources(usrs), total, pg.Number, pg.RowsPerPage), nil
}

// QueryByID returns a resource by its Ia.
func (a *App) QueryByID(ctx context.Context, resourceID string) (Resource, error) {
	id, err := uuid.Parse(resourceID)
	if err != nil {
		return Resource{}, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.resourceBus.QueryByID(ctx, id)
	if err != nil {
		return Resource{}, errs.Newf(errs.Internal, "resource missing in context: %s", err)
	}

	return toAppResource(usr), nil
}

func (a *App) QueryByName(ctx context.Context, resourceName string) (Resource, error) {

	usr, err := a.resourceBus.QueryByName(ctx, resourceName)
	if err != nil {
		return Resource{}, errs.Newf(errs.Internal, "resource missing in context: %s", err)
	}

	return toAppResource(usr), nil
}

// BulkCreate adds multiple new resources to the system.
func (a *App) BulkCreate(ctx context.Context, app BulkNewResources) (BulkResources, error) {
	if err := bulk.ValidateBatchSize(len(app.Items)); err != nil {
		return BulkResources{}, errs.New(errs.FailedPrecondition, err)
	}

	// Validate and convert all items first (fail-fast)
	var bulkErrors []errs.BulkItemError
	newResources := make([]resourcebus.NewResource, 0, len(app.Items))

	for i, item := range app.Items {
		if err := item.Validate(); err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "item",
				Error: err.Error(),
			})
			continue
		}

		nr, err := toBusNewResource(item)
		if err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "item",
				Error: err.Error(),
			})
			continue
		}
		newResources = append(newResources, nr)
	}

	if len(bulkErrors) > 0 {
		return BulkResources{}, errs.NewBulkValidationError(bulkErrors)
	}

	resources, err := a.resourceBus.BulkCreate(ctx, newResources)
	if err != nil {
		if errors.Is(err, resourcebus.ErrUniqueName) {
			return BulkResources{}, errs.New(errs.Aborted, resourcebus.ErrUniqueName)
		}
		return BulkResources{}, errs.Newf(errs.Internal, "bulkcreate: %s", err)
	}

	return BulkResources{
		Items:   toAppResources(resources),
		Created: len(resources),
	}, nil
}

// BulkUpdate modifies multiple existing resources.
func (a *App) BulkUpdate(ctx context.Context, app BulkUpdateResources) (BulkResources, error) {
	if err := bulk.ValidateBatchSize(len(app.Items)); err != nil {
		return BulkResources{}, errs.New(errs.FailedPrecondition, err)
	}

	// Validate and convert all items first (fail-fast)
	var bulkErrors []errs.BulkItemError
	updates := make([]resourcebus.UpdateResourceWithID, 0, len(app.Items))

	for i, item := range app.Items {
		if err := item.Data.Validate(); err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "data",
				Error: err.Error(),
			})
			continue
		}

		id, err := uuid.Parse(item.ID)
		if err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "id",
				Error: err.Error(),
			})
			continue
		}

		ur, err := toBusUpdateResource(item.Data)
		if err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "data",
				Error: err.Error(),
			})
			continue
		}

		updates = append(updates, resourcebus.UpdateResourceWithID{
			ID:   id,
			Data: ur,
		})
	}

	if len(bulkErrors) > 0 {
		return BulkResources{}, errs.NewBulkValidationError(bulkErrors)
	}

	resources, err := a.resourceBus.BulkUpdate(ctx, updates)
	if err != nil {
		if errors.Is(err, resourcebus.ErrUniqueName) {
			return BulkResources{}, errs.New(errs.Aborted, resourcebus.ErrUniqueName)
		}
		if errors.Is(err, resourcebus.ErrNotFound) {
			return BulkResources{}, errs.New(errs.NotFound, resourcebus.ErrNotFound)
		}
		return BulkResources{}, errs.Newf(errs.Internal, "bulkupdate: %s", err)
	}

	return BulkResources{
		Items:   toAppResources(resources),
		Updated: len(resources),
	}, nil
}

// BulkDelete removes multiple resources from the system.
func (a *App) BulkDelete(ctx context.Context, app BulkDeleteResources) (BulkDeleteResult, error) {
	if err := bulk.ValidateBatchSize(len(app.IDs)); err != nil {
		return BulkDeleteResult{}, errs.New(errs.FailedPrecondition, err)
	}

	// Validate and convert all IDs first (fail-fast)
	var bulkErrors []errs.BulkItemError
	ids := make([]uuid.UUID, 0, len(app.IDs))

	for i, idStr := range app.IDs {
		id, err := uuid.Parse(idStr)
		if err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "id",
				Error: err.Error(),
			})
			continue
		}
		ids = append(ids, id)
	}

	if len(bulkErrors) > 0 {
		return BulkDeleteResult{}, errs.NewBulkValidationError(bulkErrors)
	}

	if err := a.resourceBus.BulkDelete(ctx, ids); err != nil {
		return BulkDeleteResult{}, errs.Newf(errs.Internal, "bulkdelete: %s", err)
	}

	return BulkDeleteResult{
		Deleted: len(ids),
	}, nil
}

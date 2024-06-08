package resourceapp

import (
	"context"
	"errors"
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
		if errors.Is(err, resourcebus.ErrUniqueEmail) {
			return Resource{}, errs.New(errs.Aborted, resourcebus.ErrUniqueEmail)
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

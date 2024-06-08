package galaxyapp

import (
	"context"
	"errors"
	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/app/sdk/page"
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/google/uuid"
)

// App manages the set of app layer api functions for the galaxy domain.
type App struct {
	galaxyBus *galaxybus.Business
}

// NewApp constructs a galaxy app API for use.
func NewApp(galaxyBus *galaxybus.Business) *App {
	return &App{
		galaxyBus: galaxyBus,
	}
}

// NewAppWithAuth constructs a galaxy app API for use with auth support.
func NewAppWithAuth(galaxyBus *galaxybus.Business) *App {
	return &App{
		galaxyBus: galaxyBus,
	}
}

// Create adds a new galaxy to the system.
func (a *App) Create(ctx context.Context, app NewGalaxy) (Galaxy, error) {
	nc, err := toBusNewGalaxy(app)
	if err != nil {
		return Galaxy{}, errs.New(errs.FailedPrecondition, err)
	}

	gal, err := a.galaxyBus.Create(ctx, nc)
	if err != nil {
		if errors.Is(err, galaxybus.ErrUniqueEmail) {
			return Galaxy{}, errs.New(errs.Aborted, galaxybus.ErrUniqueEmail)
		}
		return Galaxy{}, errs.Newf(errs.Internal, "create: gal[%+v]: %s", gal, err)
	}

	return toAppGalaxy(gal), nil
}

// Update updates an existing galaxy.
func (a *App) Update(ctx context.Context, galaxyID string, app UpdateGalaxy) (Galaxy, error) {
	uu, err := toBusUpdateGalaxy(app)
	if err != nil {
		return Galaxy{}, errs.New(errs.FailedPrecondition, err)
	}

	id, err := uuid.Parse(galaxyID)
	if err != nil {
		return Galaxy{}, errs.New(errs.FailedPrecondition, err)
	}

	gal, err := a.galaxyBus.QueryByID(ctx, id)
	if err != nil {
		return Galaxy{}, errs.Newf(errs.Internal, "galaxy missing in context: %s", err)
	}

	updGal, err := a.galaxyBus.Update(ctx, gal, uu)
	if err != nil {
		return Galaxy{}, errs.Newf(errs.Internal, "update: galaxyID[%s] uu[%+v]: %s", gal.ID, uu, err)
	}

	return toAppGalaxy(updGal), nil
}

// Delete removes a galaxy from the system.
func (a *App) Delete(ctx context.Context, galaxyID string) error {
	id, err := uuid.Parse(galaxyID)
	if err != nil {
		return errs.New(errs.FailedPrecondition, err)
	}

	gal, err := a.galaxyBus.QueryByID(ctx, id)
	if err != nil {
		return errs.Newf(errs.Internal, "galaxy missing in context: %s", err)
	}

	if err := a.galaxyBus.Delete(ctx, gal); err != nil {
		return errs.Newf(errs.Internal, "delete: galaxyID[%s]: %s", gal.ID, err)
	}

	return nil
}

// Query returns a list of galaxys with paging.
func (a *App) Query(ctx context.Context, qp QueryParams) (page.Document[Galaxy], error) {
	pg, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return page.Document[Galaxy]{}, err
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return page.Document[Galaxy]{}, err
	}

	orderBy, err := order.Parse(orderByFields, qp.OrderBy, defaultOrderBy)
	if err != nil {
		return page.Document[Galaxy]{}, err
	}

	gals, err := a.galaxyBus.Query(ctx, filter, orderBy, pg.Number, pg.RowsPerPage)
	if err != nil {
		return page.Document[Galaxy]{}, errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.galaxyBus.Count(ctx, filter)
	if err != nil {
		return page.Document[Galaxy]{}, errs.Newf(errs.Internal, "count: %s", err)
	}

	return page.NewDocument(toAppGalaxies(gals), total, pg.Number, pg.RowsPerPage), nil
}

// QueryByID returns a galaxy by its Ia.
func (a *App) QueryByID(ctx context.Context, galaxyID string) (Galaxy, error) {
	id, err := uuid.Parse(galaxyID)
	if err != nil {
		return Galaxy{}, errs.New(errs.FailedPrecondition, err)
	}

	gal, err := a.galaxyBus.QueryByID(ctx, id)
	if err != nil {
		return Galaxy{}, errs.Newf(errs.Internal, "galaxy missing in context: %s", err)
	}

	return toAppGalaxy(gal), nil
}

// QueryByName returns a galaxy by its Ia.
func (a *App) QueryByName(ctx context.Context, galaxyName string) (Galaxy, error) {

	gal, err := a.galaxyBus.QueryByName(ctx, galaxyName)
	if err != nil {
		return Galaxy{}, errs.Newf(errs.Internal, "galaxy missing in context: %s", err)
	}

	return toAppGalaxy(gal), nil
}

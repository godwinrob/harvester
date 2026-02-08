package galaxyapp

import (
	"context"
	"errors"
	"github.com/godwinrob/harvester/app/sdk/bulk"
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
		if errors.Is(err, galaxybus.ErrUniqueName) {
			return Galaxy{}, errs.New(errs.Aborted, galaxybus.ErrUniqueName)
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

// BulkCreate adds multiple new galaxies to the system.
func (a *App) BulkCreate(ctx context.Context, app BulkNewGalaxies) (BulkGalaxies, error) {
	if err := bulk.ValidateBatchSize(len(app.Items)); err != nil {
		return BulkGalaxies{}, errs.New(errs.FailedPrecondition, err)
	}

	// Validate and convert all items first (fail-fast)
	var bulkErrors []errs.BulkItemError
	newGalaxies := make([]galaxybus.NewGalaxy, 0, len(app.Items))

	for i, item := range app.Items {
		if err := item.Validate(); err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "item",
				Error: err.Error(),
			})
			continue
		}

		ng, err := toBusNewGalaxy(item)
		if err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "item",
				Error: err.Error(),
			})
			continue
		}
		newGalaxies = append(newGalaxies, ng)
	}

	if len(bulkErrors) > 0 {
		return BulkGalaxies{}, errs.NewBulkValidationError(bulkErrors)
	}

	galaxies, err := a.galaxyBus.BulkCreate(ctx, newGalaxies)
	if err != nil {
		if errors.Is(err, galaxybus.ErrUniqueName) {
			return BulkGalaxies{}, errs.New(errs.Aborted, galaxybus.ErrUniqueName)
		}
		return BulkGalaxies{}, errs.Newf(errs.Internal, "bulkcreate: %s", err)
	}

	return BulkGalaxies{
		Items:   toAppGalaxies(galaxies),
		Created: len(galaxies),
	}, nil
}

// BulkUpdate modifies multiple existing galaxies.
func (a *App) BulkUpdate(ctx context.Context, app BulkUpdateGalaxies) (BulkGalaxies, error) {
	if err := bulk.ValidateBatchSize(len(app.Items)); err != nil {
		return BulkGalaxies{}, errs.New(errs.FailedPrecondition, err)
	}

	// Validate and convert all items first (fail-fast)
	var bulkErrors []errs.BulkItemError
	updates := make([]galaxybus.UpdateGalaxyWithID, 0, len(app.Items))

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

		ug, err := toBusUpdateGalaxy(item.Data)
		if err != nil {
			bulkErrors = append(bulkErrors, errs.BulkItemError{
				Index: i,
				Field: "data",
				Error: err.Error(),
			})
			continue
		}

		updates = append(updates, galaxybus.UpdateGalaxyWithID{
			ID:   id,
			Data: ug,
		})
	}

	if len(bulkErrors) > 0 {
		return BulkGalaxies{}, errs.NewBulkValidationError(bulkErrors)
	}

	galaxies, err := a.galaxyBus.BulkUpdate(ctx, updates)
	if err != nil {
		if errors.Is(err, galaxybus.ErrUniqueName) {
			return BulkGalaxies{}, errs.New(errs.Aborted, galaxybus.ErrUniqueName)
		}
		if errors.Is(err, galaxybus.ErrNotFound) {
			return BulkGalaxies{}, errs.New(errs.NotFound, galaxybus.ErrNotFound)
		}
		return BulkGalaxies{}, errs.Newf(errs.Internal, "bulkupdate: %s", err)
	}

	return BulkGalaxies{
		Items:   toAppGalaxies(galaxies),
		Updated: len(galaxies),
	}, nil
}

// BulkDelete removes multiple galaxies from the system.
func (a *App) BulkDelete(ctx context.Context, app BulkDeleteGalaxies) (BulkDeleteResult, error) {
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

	if err := a.galaxyBus.BulkDelete(ctx, ids); err != nil {
		return BulkDeleteResult{}, errs.Newf(errs.Internal, "bulkdelete: %s", err)
	}

	return BulkDeleteResult{
		Deleted: len(ids),
	}, nil
}

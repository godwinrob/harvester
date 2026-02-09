// Package galaxyapi maintains the web based api for resource access.
package galaxyapi

import (
	"context"
	"github.com/godwinrob/harvester/app/domain/galaxyapp"
	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/foundation/web"
	"net/http"
)

type api struct {
	galaxyApp *galaxyapp.App
}

func newAPI(galaxyApp *galaxyapp.App) *api {
	return &api{
		galaxyApp: galaxyApp,
	}
}

func (api *api) create(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app galaxyapp.NewGalaxy
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.galaxyApp.Create(ctx, app)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) update(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app galaxyapp.UpdateGalaxy
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.galaxyApp.Update(ctx, web.Param(r, "galaxy_id"), app)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) delete(ctx context.Context, r *http.Request) (web.Encoder, error) {
	if err := api.galaxyApp.Delete(ctx, web.Param(r, "galaxy_id")); err != nil {
		return nil, err
	}

	return nil, nil
}

func (api *api) query(ctx context.Context, r *http.Request) (web.Encoder, error) {
	qp, err := parseQueryParams(r)
	if err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.galaxyApp.Query(ctx, qp)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) queryByID(ctx context.Context, r *http.Request) (web.Encoder, error) {
	usr, err := api.galaxyApp.QueryByID(ctx, web.Param(r, "galaxy_id"))
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) queryByName(ctx context.Context, r *http.Request) (web.Encoder, error) {
	usr, err := api.galaxyApp.QueryByName(ctx, web.Param(r, "name"))
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) bulkCreate(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app galaxyapp.BulkNewGalaxies
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	result, err := api.galaxyApp.BulkCreate(ctx, app)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (api *api) bulkUpdate(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app galaxyapp.BulkUpdateGalaxies
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	result, err := api.galaxyApp.BulkUpdate(ctx, app)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (api *api) bulkDelete(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app galaxyapp.BulkDeleteGalaxies
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	result, err := api.galaxyApp.BulkDelete(ctx, app)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Package resourceapi maintains the web based api for resource access.
package resourceapi

import (
	"context"
	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/foundation/web"
	"net/http"
)

type api struct {
	resourceApp *resourceapp.App
}

func newAPI(resourceApp *resourceapp.App) *api {
	return &api{
		resourceApp: resouceApp,
	}
}

func (api *api) create(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app resourceApp.NewResource
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.resourceApp.Create(ctx, app)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) update(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app resourceApp.UpdateResource
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.resourceApp.Update(ctx, web.Param(r, "resource_id"), app)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) delete(ctx context.Context, r *http.Request) (web.Encoder, error) {
	if err := api.resourceApp.Delete(ctx, web.Param(r, "resource_id")); err != nil {
		return nil, err
	}

	return nil, nil
}

func (api *api) query(ctx context.Context, r *http.Request) (web.Encoder, error) {
	qp, err := parseQueryParams(r)
	if err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.resourceApp.Query(ctx, qp)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) queryByID(ctx context.Context, r *http.Request) (web.Encoder, error) {
	usr, err := api.resourceApp.QueryByID(ctx, web.Param(r, "resource_id"))
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) queryByName(ctx context.Context, r *http.Request) (web.Encoder, error) {
	usr, err := api.resourceApp.QueryByName(ctx, web.Param(r, "name"))
	if err != nil {
		return nil, err
	}

	return usr, nil
}

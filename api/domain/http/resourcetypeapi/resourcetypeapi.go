// Package resourcetypeapi maintains the web based api for resource type access.
package resourcetypeapi

import (
	"context"
	"net/http"

	"github.com/godwinrob/harvester/app/domain/resourcetypeapp"
	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/foundation/web"
)

type api struct {
	resourceTypeApp *resourcetypeapp.App
}

func newAPI(resourceTypeApp *resourcetypeapp.App) *api {
	return &api{
		resourceTypeApp: resourceTypeApp,
	}
}

func (api *api) create(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app resourcetypeapp.NewResourceType
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	rt, err := api.resourceTypeApp.Create(ctx, app)
	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (api *api) update(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app resourcetypeapp.UpdateResourceType
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	rt, err := api.resourceTypeApp.Update(ctx, web.Param(r, "resource_type"), app)
	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (api *api) delete(ctx context.Context, r *http.Request) (web.Encoder, error) {
	if err := api.resourceTypeApp.Delete(ctx, web.Param(r, "resource_type")); err != nil {
		return nil, err
	}

	return nil, nil
}

func (api *api) query(ctx context.Context, r *http.Request) (web.Encoder, error) {
	qp, err := parseQueryParams(r)
	if err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	rts, err := api.resourceTypeApp.Query(ctx, qp)
	if err != nil {
		return nil, err
	}

	return rts, nil
}

func (api *api) queryByID(ctx context.Context, r *http.Request) (web.Encoder, error) {
	rt, err := api.resourceTypeApp.QueryByID(ctx, web.Param(r, "resource_type"))
	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (api *api) bulkCreate(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app resourcetypeapp.BulkNewResourceTypes
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	result, err := api.resourceTypeApp.BulkCreate(ctx, app)
	if err != nil {
		return nil, err
	}

	return result, nil
}

package resourcegroupapi

import (
	"context"
	"net/http"

	"github.com/godwinrob/harvester/app/domain/resourcegroupapp"
	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/foundation/web"
)

type api struct {
	resourceGroupApp *resourcegroupapp.App
}

func newAPI(resourceGroupApp *resourcegroupapp.App) *api {
	return &api{
		resourceGroupApp: resourceGroupApp,
	}
}

func (api *api) query(ctx context.Context, r *http.Request) (web.Encoder, error) {
	qp, err := parseQueryParams(r)
	if err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	groups, err := api.resourceGroupApp.Query(ctx, qp)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (api *api) queryByID(ctx context.Context, r *http.Request) (web.Encoder, error) {
	group, err := api.resourceGroupApp.QueryByID(ctx, web.Param(r, "resource_group"))
	if err != nil {
		return nil, err
	}

	return group, nil
}

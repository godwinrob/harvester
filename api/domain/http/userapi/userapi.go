// Package userapi maintains the web based api for user access.
package userapi

import (
	"context"
	"fmt"
	"github.com/godwinrob/harvester/app/domain/userapp"
	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/foundation/web"
	"net/http"
)

type api struct {
	userApp *userapp.App
}

func newAPI(userApp *userapp.App) *api {
	return &api{
		userApp: userApp,
	}
}

func create(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(fmt.Sprintf("title:%s", "nothin")))
}

func (api *api) create(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app userapp.NewUser
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.userApp.Create(ctx, app)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) update(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app userapp.UpdateUser
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.userApp.Update(ctx, web.Param(r, "user_id"), app)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) updateRole(ctx context.Context, r *http.Request) (web.Encoder, error) {
	var app userapp.UpdateUserRole
	if err := web.Decode(r, &app); err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.userApp.UpdateRole(ctx, web.Param(r, "user_id"), app)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) delete(ctx context.Context, r *http.Request) (web.Encoder, error) {
	if err := api.userApp.Delete(ctx, web.Param(r, "user_id")); err != nil {
		return nil, err
	}

	return nil, nil
}

func (api *api) query(ctx context.Context, r *http.Request) (web.Encoder, error) {
	qp, err := parseQueryParams(r)
	if err != nil {
		return nil, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := api.userApp.Query(ctx, qp)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (api *api) queryByID(ctx context.Context, r *http.Request) (web.Encoder, error) {
	usr, err := api.userApp.QueryByID(ctx, web.Param(r, "user_id"))
	if err != nil {
		return nil, err
	}

	return usr, nil
}

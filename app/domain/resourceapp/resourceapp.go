package userapp

import (
	"context"
	"errors"
	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/app/sdk/page"
	"github.com/godwinrob/harvester/business/domain/userbus"
	"github.com/godwinrob/harvester/business/sdk/order"
	"github.com/google/uuid"
)

// App manages the set of app layer api functions for the user domain.
type App struct {
	userBus *userbus.Business
}

// NewApp constructs a user app API for use.
func NewApp(userBus *userbus.Business) *App {
	return &App{
		userBus: userBus,
	}
}

// NewAppWithAuth constructs a user app API for use with auth support.
func NewAppWithAuth(userBus *userbus.Business) *App {
	return &App{
		userBus: userBus,
	}
}

// Create adds a new user to the system.
func (a *App) Create(ctx context.Context, app NewUser) (User, error) {
	nc, err := toBusNewUser(app)
	if err != nil {
		return User{}, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.userBus.Create(ctx, nc)
	if err != nil {
		if errors.Is(err, userbus.ErrUniqueEmail) {
			return User{}, errs.New(errs.Aborted, userbus.ErrUniqueEmail)
		}
		return User{}, errs.Newf(errs.Internal, "create: usr[%+v]: %s", usr, err)
	}

	return toAppUser(usr), nil
}

// Update updates an existing user.
func (a *App) Update(ctx context.Context, userID string, app UpdateUser) (User, error) {
	uu, err := toBusUpdateUser(app)
	if err != nil {
		return User{}, errs.New(errs.FailedPrecondition, err)
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return User{}, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.userBus.QueryByID(ctx, id)
	if err != nil {
		return User{}, errs.Newf(errs.Internal, "user missing in context: %s", err)
	}

	updUsr, err := a.userBus.Update(ctx, usr, uu)
	if err != nil {
		return User{}, errs.Newf(errs.Internal, "update: userID[%s] uu[%+v]: %s", usr.ID, uu, err)
	}

	return toAppUser(updUsr), nil
}

// UpdateRole updates an existing user's role.
func (a *App) UpdateRole(ctx context.Context, userID string, app UpdateUserRole) (User, error) {
	uu, err := toBusUpdateUserRole(app)
	if err != nil {
		return User{}, errs.New(errs.FailedPrecondition, err)
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return User{}, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.userBus.QueryByID(ctx, id)
	if err != nil {
		return User{}, errs.Newf(errs.Internal, "user missing in context: %s", err)
	}

	updUsr, err := a.userBus.Update(ctx, usr, uu)
	if err != nil {
		return User{}, errs.Newf(errs.Internal, "updaterole: userID[%s] uu[%+v]: %s", usr.ID, uu, err)
	}

	return toAppUser(updUsr), nil
}

// Delete removes a user from the system.
func (a *App) Delete(ctx context.Context, userID string) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.userBus.QueryByID(ctx, id)
	if err != nil {
		return errs.Newf(errs.Internal, "user missing in context: %s", err)
	}

	if err := a.userBus.Delete(ctx, usr); err != nil {
		return errs.Newf(errs.Internal, "delete: userID[%s]: %s", usr.ID, err)
	}

	return nil
}

// Query returns a list of users with paging.
func (a *App) Query(ctx context.Context, qp QueryParams) (page.Document[User], error) {
	pg, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return page.Document[User]{}, err
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return page.Document[User]{}, err
	}

	orderBy, err := order.Parse(orderByFields, qp.OrderBy, defaultOrderBy)
	if err != nil {
		return page.Document[User]{}, err
	}

	usrs, err := a.userBus.Query(ctx, filter, orderBy, pg.Number, pg.RowsPerPage)
	if err != nil {
		return page.Document[User]{}, errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.userBus.Count(ctx, filter)
	if err != nil {
		return page.Document[User]{}, errs.Newf(errs.Internal, "count: %s", err)
	}

	return page.NewDocument(toAppUsers(usrs), total, pg.Number, pg.RowsPerPage), nil
}

// QueryByID returns a user by its Ia.
func (a *App) QueryByID(ctx context.Context, userID string) (User, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return User{}, errs.New(errs.FailedPrecondition, err)
	}

	usr, err := a.userBus.QueryByID(ctx, id)
	if err != nil {
		return User{}, errs.Newf(errs.Internal, "user missing in context: %s", err)
	}

	return toAppUser(usr), nil
}

// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"context"

	"github.com/godwinrob/harvester/api/sdk/http/mid"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
	"github.com/jmoiron/sqlx"
)

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(*logger.Logger, *sqlx.DB, *web.App)
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, db *sqlx.DB, routeAdder RouteAdder) *web.App {
	l := func(ctx context.Context, msg string, args ...any) {
		log.Info(ctx, msg, args...)
	}

	app := web.NewApp(l, mid.Logger(log), mid.Error(log), mid.Panics())

	routeAdder.Add(log, db, app)

	return app
}

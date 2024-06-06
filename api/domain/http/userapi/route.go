package userapi

import (
	"github.com/godwinrob/harvester/app/domain/userapp"
	"github.com/godwinrob/harvester/business/domain/userbus"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log     *logger.Logger
	UserBus *userbus.Business
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	api := newAPI(userapp.NewApp(cfg.UserBus))
	app.HandleFunc("GET /users", api.query)
	app.HandleFunc("GET /users/{user_id}", api.queryByID)
	app.HandleFunc("POST /users", api.create)
	app.HandleFunc("PUT /users/role/{user_id}", api.updateRole)
	app.HandleFunc("PUT /users/{user_id}", api.update)
	app.HandleFunc("DELETE /users/{user_id}", api.delete)
}

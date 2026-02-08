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

	api := newAPI(userapp.NewApp(cfg.UserBus))
	app.HandleFunc("GET /v1/users", api.query)
	app.HandleFunc("GET /v1/users/{user_id}", api.queryByID)
	app.HandleFunc("POST /v1/users", api.create)
	app.HandleFunc("POST /v1/users/bulk", api.bulkCreate)
	app.HandleFunc("PUT /v1/users/role/{user_id}", api.updateRole)
	app.HandleFunc("PUT /v1/users/bulk", api.bulkUpdate)
	app.HandleFunc("PUT /v1/users/{user_id}", api.update)
	app.HandleFunc("DELETE /v1/users/bulk", api.bulkDelete)
	app.HandleFunc("DELETE /v1/users/{user_id}", api.delete)
}

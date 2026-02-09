package resourceapi

import (
	"github.com/godwinrob/harvester/app/domain/resourceapp"
	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log         *logger.Logger
	ResourceBus *resourcebus.Business
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {

	api := newAPI(resourceapp.NewApp(cfg.ResourceBus))
	app.HandleFunc("POST /v1/resources", api.create)
	app.HandleFunc("POST /v1/resources/bulk", api.bulkCreate)
	app.HandleFunc("GET /v1/resources", api.query)
	app.HandleFunc("GET /v1/resources/{resource_id}", api.queryByID)
	app.HandleFunc("GET /v1/resources/name/{name}", api.queryByName)
	app.HandleFunc("PUT /v1/resources/bulk", api.bulkUpdate)
	app.HandleFunc("PUT /v1/resources/{resource_id}", api.update)
	app.HandleFunc("DELETE /v1/resources/bulk", api.bulkDelete)
	app.HandleFunc("DELETE /v1/resources/{resource_id}", api.delete)
}

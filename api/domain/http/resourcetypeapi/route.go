package resourcetypeapi

import (
	"github.com/godwinrob/harvester/app/domain/resourcetypeapp"
	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log             *logger.Logger
	ResourceTypeBus *resourcetypebus.Business
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	api := newAPI(resourcetypeapp.NewApp(cfg.ResourceTypeBus))
	app.HandleFunc("GET /v1/resource-types", api.query)
	app.HandleFunc("GET /v1/resource-types/{resource_type}", api.queryByID)
	app.HandleFunc("POST /v1/resource-types", api.create)
	app.HandleFunc("POST /v1/resource-types/bulk", api.bulkCreate)
	app.HandleFunc("PUT /v1/resource-types/{resource_type}", api.update)
	app.HandleFunc("DELETE /v1/resource-types/{resource_type}", api.delete)
}

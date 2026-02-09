package resourcegroupapi

import (
	"github.com/godwinrob/harvester/app/domain/resourcegroupapp"
	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log              *logger.Logger
	ResourceGroupBus *resourcegroupbus.Business
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	api := newAPI(resourcegroupapp.NewApp(cfg.ResourceGroupBus))
	app.HandleFunc("GET /v1/resource-groups", api.query)
	app.HandleFunc("GET /v1/resource-groups/{resource_group}", api.queryByID)
}

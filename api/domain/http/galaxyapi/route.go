package galaxyapi

import (
	"github.com/godwinrob/harvester/app/domain/galaxyapp"
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log       *logger.Logger
	GalaxyBus *galaxybus.Business
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {

	api := newAPI(galaxyapp.NewApp(cfg.GalaxyBus))
	app.HandleFunc("POST /v1/galaxies", api.create)
	app.HandleFunc("GET /v1/galaxies", api.query)
	app.HandleFunc("GET /v1/galaxies/{galaxy_id}", api.queryByID)
	app.HandleFunc("GET /v1/galaxies/name/{name}", api.queryByName)
	app.HandleFunc("PUT /v1/galaxies/{galaxy_id}", api.update)
	app.HandleFunc("DELETE /v1/galaxies/{galaxy_id}", api.delete)
}

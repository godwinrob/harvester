package all

import (
	"github.com/godwinrob/harvester/api/domain/http/galaxyapi"
	"github.com/godwinrob/harvester/api/domain/http/resourceapi"
	"github.com/godwinrob/harvester/api/domain/http/resourcegroupapi"
	"github.com/godwinrob/harvester/api/domain/http/resourcetypeapi"
	"github.com/godwinrob/harvester/api/domain/http/userapi"
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/business/domain/galaxybus/stores/galaxydb"
	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/business/domain/resourcebus/stores/resourcedb"
	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
	"github.com/godwinrob/harvester/business/domain/resourcegroupbus/stores/resourcegroupdb"
	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
	"github.com/godwinrob/harvester/business/domain/resourcetypebus/stores/resourcetypedb"
	"github.com/godwinrob/harvester/business/domain/userbus"
	"github.com/godwinrob/harvester/business/domain/userbus/stores/userdb"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (a add) Add(log *logger.Logger, db *sqlx.DB, app *web.App) {

	userapi.Routes(app, userapi.Config{
		Log:     log,
		UserBus: userbus.NewBusiness(log, userdb.NewStore(log, db)),
	})

	galaxyapi.Routes(app, galaxyapi.Config{
		Log:       log,
		GalaxyBus: galaxybus.NewBusiness(log, galaxydb.NewStore(log, db)),
	})

	resourceapi.Routes(app, resourceapi.Config{
		Log:         log,
		ResourceBus: resourcebus.NewBusiness(log, resourcedb.NewStore(log, db)),
	})

	resourcetypeapi.Routes(app, resourcetypeapi.Config{
		Log:             log,
		ResourceTypeBus: resourcetypebus.NewBusiness(log, resourcetypedb.NewStore(log, db)),
	})

	resourcegroupapi.Routes(app, resourcegroupapi.Config{
		Log:              log,
		ResourceGroupBus: resourcegroupbus.NewBusiness(log, resourcegroupdb.NewStore(log, db)),
	})
}

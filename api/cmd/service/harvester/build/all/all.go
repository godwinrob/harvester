package all

import (
	"github.com/godwinrob/harvester/api/domain/http/userapi"
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
}

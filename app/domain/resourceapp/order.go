package resourceapp

import (
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var defaultOrderBy = order.NewBy("date_created", order.ASC)

var orderByFields = map[string]string{
	"galaxy_id":   galaxybus.OrderByID,
	"galaxy_name": galaxybus.OrderByName,
	"enabled":     galaxybus.OrderByEnabled,
}

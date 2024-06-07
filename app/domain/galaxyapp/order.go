package galaxyapp

import (
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var defaultOrderBy = order.NewBy("galaxy_id", order.ASC)

var orderByFields = map[string]string{
	"galaxy_id":    galaxybus.OrderByID,
	"galaxy_name":  galaxybus.OrderByName,
	"date_created": galaxybus.OrderByCreatedDate,
	"enabled":      galaxybus.OrderByEnabled,
}

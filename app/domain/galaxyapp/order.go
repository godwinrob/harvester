package galaxyapp

import (
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var defaultOrderBy = order.NewBy("galaxy_id", order.ASC)

var orderByFields = map[string]string{
	"galaxy_id":     galaxybus.OrderByID,
	"name":          galaxybus.OrderByName,
	"galaxy_name":   galaxybus.OrderByName,
	"ownerUserID":   galaxybus.OrderByOwnerUserID,
	"owner_user_id": galaxybus.OrderByOwnerUserID,
	"dateCreated":   galaxybus.OrderByDateCreated,
	"date_created":  galaxybus.OrderByDateCreated,
	"enabled":       galaxybus.OrderByEnabled,
}

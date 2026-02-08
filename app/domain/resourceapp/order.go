package resourceapp

import (
	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var defaultOrderBy = order.NewBy(resourcebus.OrderByID, order.ASC)

var orderByFields = map[string]string{
	"resource_id":   resourcebus.OrderByID,
	"resource_name": resourcebus.OrderByName,
	"enabled":       resourcebus.OrderByEnabled,
}

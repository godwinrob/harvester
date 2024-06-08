package resourcedb

import (
	"fmt"

	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var orderByFields = map[string]string{
	resourcebus.OrderByID:      "resource_id",
	resourcebus.OrderByName:    "resource_name",
	resourcebus.OrderByEnabled: "enabled",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}

package resourcegroupdb

import (
	"fmt"

	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var orderByFields = map[string]string{
	resourcegroupbus.OrderByResourceGroup: "resource_group",
	resourcegroupbus.OrderByGroupName:     "group_name",
	resourcegroupbus.OrderByGroupLevel:    "group_level",
	resourcegroupbus.OrderByGroupOrder:    "group_order",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}

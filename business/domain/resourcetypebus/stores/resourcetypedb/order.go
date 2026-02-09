package resourcetypedb

import (
	"fmt"

	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var orderByFields = map[string]string{
	resourcetypebus.OrderByResourceType:     "resource_type",
	resourcetypebus.OrderByResourceTypeName: "resource_type_name",
	resourcetypebus.OrderByResourceCategory: "resource_category",
	resourcetypebus.OrderByResourceGroup:    "resource_group",
	resourcetypebus.OrderByEnterable:        "enterable",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}

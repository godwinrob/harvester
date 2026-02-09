package galaxydb

import (
	"fmt"

	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var orderByFields = map[string]string{
	galaxybus.OrderByID:          "galaxy_id",
	galaxybus.OrderByName:        "galaxy_name",
	galaxybus.OrderByOwnerUserID: "owner_user_id",
	galaxybus.OrderByDateCreated: "date_created",
	galaxybus.OrderByEnabled:     "enabled",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}

package resourcedb

import (
	"fmt"

	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var orderByFields = map[string]string{
	resourcebus.OrderByID:            "resource_id",
	resourcebus.OrderByName:          "resource_name",
	resourcebus.OrderByResourceType:  "resource_type",
	resourcebus.OrderByVerified:      "verified",
	resourcebus.OrderByUnavailableAt: "unavailable_at",
	resourcebus.OrderByAddedAt:       "added_at",
	resourcebus.OrderByEnabled:       "enabled",
	resourcebus.OrderByCR:            "cr",
	resourcebus.OrderByCD:            "cd",
	resourcebus.OrderByDR:            "dr",
	resourcebus.OrderByFL:            "fl",
	resourcebus.OrderByHR:            "hr",
	resourcebus.OrderByMA:            "ma",
	resourcebus.OrderByPE:            "pe",
	resourcebus.OrderByOQ:            "oq",
	resourcebus.OrderBySR:            "sr",
	resourcebus.OrderByUT:            "ut",
	resourcebus.OrderByER:            "er",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}

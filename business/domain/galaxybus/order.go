package galaxybus

import "github.com/godwinrob/harvester/business/sdk/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByID, order.ASC)

// Set of fields that the results can be ordered by.
const (
	OrderByID      = "galaxy_id"
	OrderByName    = "name"
	OrderByEnabled = "enabled"
)

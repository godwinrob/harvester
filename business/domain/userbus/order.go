package userbus

import "github.com/godwinrob/harvester/business/sdk/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByID, order.ASC)

// Set of fields that the results can be ordered by.
const (
	OrderByID          = "user_id"
	OrderByName        = "name"
	OrderByEmail       = "email"
	OrderByRoles       = "roles"
	OrderByGuild       = "guild"
	OrderByDateCreated = "date_created"
	OrderByEnabled     = "enabled"
)

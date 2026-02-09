package resourcegroupbus

import "github.com/godwinrob/harvester/business/sdk/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByGroupOrder, order.ASC)

// Set of fields that the results can be ordered by.
const (
	OrderByResourceGroup = "resource_group"
	OrderByGroupName     = "group_name"
	OrderByGroupLevel    = "group_level"
	OrderByGroupOrder    = "group_order"
)

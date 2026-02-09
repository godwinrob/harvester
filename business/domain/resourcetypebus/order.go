package resourcetypebus

import "github.com/godwinrob/harvester/business/sdk/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByResourceType, order.ASC)

// Set of fields that the results can be ordered by.
const (
	OrderByResourceType     = "resource_type"
	OrderByResourceTypeName = "resource_type_name"
	OrderByResourceCategory = "resource_category"
	OrderByResourceGroup    = "resource_group"
	OrderByEnterable        = "enterable"
)

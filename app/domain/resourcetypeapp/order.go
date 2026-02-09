package resourcetypeapp

import (
	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var defaultOrderBy = order.NewBy(resourcetypebus.OrderByResourceType, order.ASC)

var orderByFields = map[string]string{
	"resource_type":      resourcetypebus.OrderByResourceType,
	"resourceType":       resourcetypebus.OrderByResourceType,
	"resource_type_name": resourcetypebus.OrderByResourceTypeName,
	"resourceTypeName":   resourcetypebus.OrderByResourceTypeName,
	"resource_category":  resourcetypebus.OrderByResourceCategory,
	"resourceCategory":   resourcetypebus.OrderByResourceCategory,
	"resource_group":     resourcetypebus.OrderByResourceGroup,
	"resourceGroup":      resourcetypebus.OrderByResourceGroup,
	"enterable":          resourcetypebus.OrderByEnterable,
}

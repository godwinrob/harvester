package resourcegroupapp

import (
	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var defaultOrderBy = order.NewBy(resourcegroupbus.OrderByGroupOrder, order.ASC)

var orderByFields = map[string]string{
	"resource_group": resourcegroupbus.OrderByResourceGroup,
	"resourceGroup":  resourcegroupbus.OrderByResourceGroup,
	"group_name":     resourcegroupbus.OrderByGroupName,
	"groupName":      resourcegroupbus.OrderByGroupName,
	"group_level":    resourcegroupbus.OrderByGroupLevel,
	"groupLevel":     resourcegroupbus.OrderByGroupLevel,
	"group_order":    resourcegroupbus.OrderByGroupOrder,
	"groupOrder":     resourcegroupbus.OrderByGroupOrder,
}

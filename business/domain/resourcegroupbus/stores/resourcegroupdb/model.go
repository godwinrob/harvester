package resourcegroupdb

import (
	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
)

type resourceGroup struct {
	ResourceGroup string `db:"resource_group"`
	GroupName     string `db:"group_name"`
	GroupLevel    int16  `db:"group_level"`
	GroupOrder    int16  `db:"group_order"`
	ContainerType string `db:"container_type"`
}

func toBusResourceGroup(db resourceGroup) resourcegroupbus.ResourceGroup {
	return resourcegroupbus.ResourceGroup{
		ResourceGroup: db.ResourceGroup,
		GroupName:     db.GroupName,
		GroupLevel:    db.GroupLevel,
		GroupOrder:    db.GroupOrder,
		ContainerType: db.ContainerType,
	}
}

func toBusResourceGroups(dbs []resourceGroup) []resourcegroupbus.ResourceGroup {
	bus := make([]resourcegroupbus.ResourceGroup, len(dbs))
	for i, db := range dbs {
		bus[i] = toBusResourceGroup(db)
	}
	return bus
}

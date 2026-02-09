package resourcegroupapp

import (
	"encoding/json"

	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
)

// QueryParams represents the set of possible query strings.
type QueryParams struct {
	Page          string
	Rows          string
	OrderBy       string
	ResourceGroup string
	GroupName     string
	GroupLevel    string
	ContainerType string
}

// ResourceGroup represents information about a resource group.
type ResourceGroup struct {
	ResourceGroup string `json:"resourceGroup"`
	GroupName     string `json:"groupName"`
	GroupLevel    int16  `json:"groupLevel"`
	GroupOrder    int16  `json:"groupOrder"`
	ContainerType string `json:"containerType"`
}

// Encode implements the encoder interface.
func (app ResourceGroup) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

func toAppResourceGroup(bus resourcegroupbus.ResourceGroup) ResourceGroup {
	return ResourceGroup{
		ResourceGroup: bus.ResourceGroup,
		GroupName:     bus.GroupName,
		GroupLevel:    bus.GroupLevel,
		GroupOrder:    bus.GroupOrder,
		ContainerType: bus.ContainerType,
	}
}

func toAppResourceGroups(groups []resourcegroupbus.ResourceGroup) []ResourceGroup {
	app := make([]ResourceGroup, len(groups))
	for i, g := range groups {
		app[i] = toAppResourceGroup(g)
	}
	return app
}

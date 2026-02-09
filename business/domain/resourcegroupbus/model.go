package resourcegroupbus

// ResourceGroup represents a node in the resource type hierarchy.
type ResourceGroup struct {
	ResourceGroup string
	GroupName     string
	GroupLevel    int16
	GroupOrder    int16
	ContainerType string
}

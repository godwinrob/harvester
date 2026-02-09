package resourcegroupbus

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ResourceGroup *string
	GroupName     *string
	GroupLevel    *int16
	ContainerType *string
}

package resourcetypebus

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ResourceType     *string
	ResourceTypeName *string
	ResourceCategory *string
	ResourceGroup    *string
	Enterable        *bool
	ContainerType    *string
}

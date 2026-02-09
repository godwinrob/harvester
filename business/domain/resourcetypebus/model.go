package resourcetypebus

// ResourceType represents a resource type definition with stat ranges.
type ResourceType struct {
	ResourceType     string
	ResourceTypeName string
	ResourceCategory string
	ResourceGroup    string
	Enterable        bool
	MaxTypes         int16
	CRmin            int16
	CRmax            int16
	CDmin            int16
	CDmax            int16
	DRmin            int16
	DRmax            int16
	FLmin            int16
	FLmax            int16
	HRmin            int16
	HRmax            int16
	MAmin            int16
	MAmax            int16
	PEmin            int16
	PEmax            int16
	OQmin            int16
	OQmax            int16
	SRmin            int16
	SRmax            int16
	UTmin            int16
	UTmax            int16
	ERmin            int16
	ERmax            int16
	ContainerType    string
	InventoryType    string
	SpecificPlanet   int16
}

// NewResourceType contains information needed to create a new resource type.
type NewResourceType struct {
	ResourceType     string
	ResourceTypeName string
	ResourceCategory string
	ResourceGroup    string
	Enterable        bool
	MaxTypes         int16
	CRmin            int16
	CRmax            int16
	CDmin            int16
	CDmax            int16
	DRmin            int16
	DRmax            int16
	FLmin            int16
	FLmax            int16
	HRmin            int16
	HRmax            int16
	MAmin            int16
	MAmax            int16
	PEmin            int16
	PEmax            int16
	OQmin            int16
	OQmax            int16
	SRmin            int16
	SRmax            int16
	UTmin            int16
	UTmax            int16
	ERmin            int16
	ERmax            int16
	ContainerType    string
	InventoryType    string
	SpecificPlanet   int16
}

// UpdateResourceType contains information needed to update a resource type.
type UpdateResourceType struct {
	ResourceTypeName *string
	ResourceCategory *string
	ResourceGroup    *string
	Enterable        *bool
	MaxTypes         *int16
	CRmin            *int16
	CRmax            *int16
	CDmin            *int16
	CDmax            *int16
	DRmin            *int16
	DRmax            *int16
	FLmin            *int16
	FLmax            *int16
	HRmin            *int16
	HRmax            *int16
	MAmin            *int16
	MAmax            *int16
	PEmin            *int16
	PEmax            *int16
	OQmin            *int16
	OQmax            *int16
	SRmin            *int16
	SRmax            *int16
	UTmin            *int16
	UTmax            *int16
	ERmin            *int16
	ERmax            *int16
	ContainerType    *string
	InventoryType    *string
	SpecificPlanet   *int16
}

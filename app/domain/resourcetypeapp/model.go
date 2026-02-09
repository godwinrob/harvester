package resourcetypeapp

import (
	"encoding/json"

	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
	"github.com/godwinrob/harvester/foundation/validate"
)

// QueryParams represents the set of possible query strings.
type QueryParams struct {
	Page             string
	Rows             string
	OrderBy          string
	ResourceType     string
	ResourceTypeName string
	ResourceCategory string
	ResourceGroup    string
	Enterable        string
	ContainerType    string
}

// ResourceType represents information about a resource type.
type ResourceType struct {
	ResourceType     string `json:"resourceType"`
	ResourceTypeName string `json:"resourceTypeName"`
	ResourceCategory string `json:"resourceCategory"`
	ResourceGroup    string `json:"resourceGroup"`
	Enterable        bool   `json:"enterable"`
	MaxTypes         int16  `json:"maxTypes"`
	CRmin            int16  `json:"crMin"`
	CRmax            int16  `json:"crMax"`
	CDmin            int16  `json:"cdMin"`
	CDmax            int16  `json:"cdMax"`
	DRmin            int16  `json:"drMin"`
	DRmax            int16  `json:"drMax"`
	FLmin            int16  `json:"flMin"`
	FLmax            int16  `json:"flMax"`
	HRmin            int16  `json:"hrMin"`
	HRmax            int16  `json:"hrMax"`
	MAmin            int16  `json:"maMin"`
	MAmax            int16  `json:"maMax"`
	PEmin            int16  `json:"peMin"`
	PEmax            int16  `json:"peMax"`
	OQmin            int16  `json:"oqMin"`
	OQmax            int16  `json:"oqMax"`
	SRmin            int16  `json:"srMin"`
	SRmax            int16  `json:"srMax"`
	UTmin            int16  `json:"utMin"`
	UTmax            int16  `json:"utMax"`
	ERmin            int16  `json:"erMin"`
	ERmax            int16  `json:"erMax"`
	ContainerType    string `json:"containerType"`
	InventoryType    string `json:"inventoryType"`
	SpecificPlanet   int16  `json:"specificPlanet"`
}

// Encode implements the encoder interface.
func (app ResourceType) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

func toAppResourceType(bus resourcetypebus.ResourceType) ResourceType {
	return ResourceType{
		ResourceType:     bus.ResourceType,
		ResourceTypeName: bus.ResourceTypeName,
		ResourceCategory: bus.ResourceCategory,
		ResourceGroup:    bus.ResourceGroup,
		Enterable:        bus.Enterable,
		MaxTypes:         bus.MaxTypes,
		CRmin:            bus.CRmin,
		CRmax:            bus.CRmax,
		CDmin:            bus.CDmin,
		CDmax:            bus.CDmax,
		DRmin:            bus.DRmin,
		DRmax:            bus.DRmax,
		FLmin:            bus.FLmin,
		FLmax:            bus.FLmax,
		HRmin:            bus.HRmin,
		HRmax:            bus.HRmax,
		MAmin:            bus.MAmin,
		MAmax:            bus.MAmax,
		PEmin:            bus.PEmin,
		PEmax:            bus.PEmax,
		OQmin:            bus.OQmin,
		OQmax:            bus.OQmax,
		SRmin:            bus.SRmin,
		SRmax:            bus.SRmax,
		UTmin:            bus.UTmin,
		UTmax:            bus.UTmax,
		ERmin:            bus.ERmin,
		ERmax:            bus.ERmax,
		ContainerType:    bus.ContainerType,
		InventoryType:    bus.InventoryType,
		SpecificPlanet:   bus.SpecificPlanet,
	}
}

func toAppResourceTypes(rts []resourcetypebus.ResourceType) []ResourceType {
	app := make([]ResourceType, len(rts))
	for i, rt := range rts {
		app[i] = toAppResourceType(rt)
	}
	return app
}

// =============================================================================

// NewResourceType defines the data needed to add a new resource type.
type NewResourceType struct {
	ResourceType     string `json:"resourceType" validate:"required"`
	ResourceTypeName string `json:"resourceTypeName" validate:"required"`
	ResourceCategory string `json:"resourceCategory"`
	ResourceGroup    string `json:"resourceGroup"`
	Enterable        bool   `json:"enterable"`
	MaxTypes         int16  `json:"maxTypes"`
	CRmin            int16  `json:"crMin"`
	CRmax            int16  `json:"crMax"`
	CDmin            int16  `json:"cdMin"`
	CDmax            int16  `json:"cdMax"`
	DRmin            int16  `json:"drMin"`
	DRmax            int16  `json:"drMax"`
	FLmin            int16  `json:"flMin"`
	FLmax            int16  `json:"flMax"`
	HRmin            int16  `json:"hrMin"`
	HRmax            int16  `json:"hrMax"`
	MAmin            int16  `json:"maMin"`
	MAmax            int16  `json:"maMax"`
	PEmin            int16  `json:"peMin"`
	PEmax            int16  `json:"peMax"`
	OQmin            int16  `json:"oqMin"`
	OQmax            int16  `json:"oqMax"`
	SRmin            int16  `json:"srMin"`
	SRmax            int16  `json:"srMax"`
	UTmin            int16  `json:"utMin"`
	UTmax            int16  `json:"utMax"`
	ERmin            int16  `json:"erMin"`
	ERmax            int16  `json:"erMax"`
	ContainerType    string `json:"containerType"`
	InventoryType    string `json:"inventoryType"`
	SpecificPlanet   int16  `json:"specificPlanet"`
}

// Decode implements the decoder interface.
func (app *NewResourceType) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app NewResourceType) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}
	return nil
}

func toBusNewResourceType(app NewResourceType) resourcetypebus.NewResourceType {
	return resourcetypebus.NewResourceType{
		ResourceType:     app.ResourceType,
		ResourceTypeName: app.ResourceTypeName,
		ResourceCategory: app.ResourceCategory,
		ResourceGroup:    app.ResourceGroup,
		Enterable:        app.Enterable,
		MaxTypes:         app.MaxTypes,
		CRmin:            app.CRmin,
		CRmax:            app.CRmax,
		CDmin:            app.CDmin,
		CDmax:            app.CDmax,
		DRmin:            app.DRmin,
		DRmax:            app.DRmax,
		FLmin:            app.FLmin,
		FLmax:            app.FLmax,
		HRmin:            app.HRmin,
		HRmax:            app.HRmax,
		MAmin:            app.MAmin,
		MAmax:            app.MAmax,
		PEmin:            app.PEmin,
		PEmax:            app.PEmax,
		OQmin:            app.OQmin,
		OQmax:            app.OQmax,
		SRmin:            app.SRmin,
		SRmax:            app.SRmax,
		UTmin:            app.UTmin,
		UTmax:            app.UTmax,
		ERmin:            app.ERmin,
		ERmax:            app.ERmax,
		ContainerType:    app.ContainerType,
		InventoryType:    app.InventoryType,
		SpecificPlanet:   app.SpecificPlanet,
	}
}

// =============================================================================

// UpdateResourceType defines the data needed to update a resource type.
type UpdateResourceType struct {
	ResourceTypeName *string `json:"resourceTypeName"`
	ResourceCategory *string `json:"resourceCategory"`
	ResourceGroup    *string `json:"resourceGroup"`
	Enterable        *bool   `json:"enterable"`
	MaxTypes         *int16  `json:"maxTypes"`
	CRmin            *int16  `json:"crMin"`
	CRmax            *int16  `json:"crMax"`
	CDmin            *int16  `json:"cdMin"`
	CDmax            *int16  `json:"cdMax"`
	DRmin            *int16  `json:"drMin"`
	DRmax            *int16  `json:"drMax"`
	FLmin            *int16  `json:"flMin"`
	FLmax            *int16  `json:"flMax"`
	HRmin            *int16  `json:"hrMin"`
	HRmax            *int16  `json:"hrMax"`
	MAmin            *int16  `json:"maMin"`
	MAmax            *int16  `json:"maMax"`
	PEmin            *int16  `json:"peMin"`
	PEmax            *int16  `json:"peMax"`
	OQmin            *int16  `json:"oqMin"`
	OQmax            *int16  `json:"oqMax"`
	SRmin            *int16  `json:"srMin"`
	SRmax            *int16  `json:"srMax"`
	UTmin            *int16  `json:"utMin"`
	UTmax            *int16  `json:"utMax"`
	ERmin            *int16  `json:"erMin"`
	ERmax            *int16  `json:"erMax"`
	ContainerType    *string `json:"containerType"`
	InventoryType    *string `json:"inventoryType"`
	SpecificPlanet   *int16  `json:"specificPlanet"`
}

// Decode implements the decoder interface.
func (app *UpdateResourceType) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app UpdateResourceType) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}
	return nil
}

func toBusUpdateResourceType(app UpdateResourceType) resourcetypebus.UpdateResourceType {
	return resourcetypebus.UpdateResourceType{
		ResourceTypeName: app.ResourceTypeName,
		ResourceCategory: app.ResourceCategory,
		ResourceGroup:    app.ResourceGroup,
		Enterable:        app.Enterable,
		MaxTypes:         app.MaxTypes,
		CRmin:            app.CRmin,
		CRmax:            app.CRmax,
		CDmin:            app.CDmin,
		CDmax:            app.CDmax,
		DRmin:            app.DRmin,
		DRmax:            app.DRmax,
		FLmin:            app.FLmin,
		FLmax:            app.FLmax,
		HRmin:            app.HRmin,
		HRmax:            app.HRmax,
		MAmin:            app.MAmin,
		MAmax:            app.MAmax,
		PEmin:            app.PEmin,
		PEmax:            app.PEmax,
		OQmin:            app.OQmin,
		OQmax:            app.OQmax,
		SRmin:            app.SRmin,
		SRmax:            app.SRmax,
		UTmin:            app.UTmin,
		UTmax:            app.UTmax,
		ERmin:            app.ERmin,
		ERmax:            app.ERmax,
		ContainerType:    app.ContainerType,
		InventoryType:    app.InventoryType,
		SpecificPlanet:   app.SpecificPlanet,
	}
}

// =============================================================================

// BulkNewResourceTypes defines the data needed to bulk create resource types.
type BulkNewResourceTypes struct {
	Items []NewResourceType `json:"items" validate:"required,min=1,max=100,dive"`
}

// Decode implements the decoder interface.
func (app *BulkNewResourceTypes) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app BulkNewResourceTypes) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}
	return nil
}

// BulkResourceTypes represents the result of a bulk resource type operation.
type BulkResourceTypes struct {
	Items   []ResourceType `json:"items"`
	Created int            `json:"created,omitempty"`
}

// Encode implements the encoder interface.
func (app BulkResourceTypes) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

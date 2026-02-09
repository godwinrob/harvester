package resourceapp

import (
	"encoding/json"
	"fmt"
	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/google/uuid"
	"time"

	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/foundation/validate"
)

// QueryParams represents the set of possible query strings.
type QueryParams struct {
	Page         string
	Rows         string
	OrderBy      string
	ID           string
	Name         string
	ResourceType  string
	ResourceGroup string
	AddedAtDate   string
}

// Resource represents information about an individual resource.
type Resource struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	GalaxyID          string `json:"galaxyID"`
	AddedAtDate       string `json:"addedAtDate"`
	UpdatedAtDate     string `json:"updatedAtDate"`
	AddedUserID       string `json:"addedUserID"`
	ResourceType      string `json:"resourceType"`
	UnavailableAt     string `json:"unavailableAt"`
	UnavailableUserID string `json:"unavailableUserID"`
	Verified          bool   `json:"verified"`
	VerifiedUserID    string `json:"verifiedUserID"`
	CR                int16  `json:"cr"`
	CD                int16  `json:"cd"`
	DR                int16  `json:"dr"`
	FL                int16  `json:"fl"`
	HR                int16  `json:"hr"`
	MA                int16  `json:"ma"`
	PE                int16  `json:"pe"`
	OQ                int16  `json:"oq"`
	SR                int16  `json:"sr"`
	UT                int16  `json:"ut"`
	ER                int16  `json:"er"`
}

// Encode implments the encoder interface.
func (app Resource) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

func toAppResource(bus resourcebus.Resource) Resource {
	// Handle nullable time fields - only format if not zero
	var unavailableAt string
	if !bus.UnavailableAt.IsZero() {
		unavailableAt = bus.UnavailableAt.Format(time.RFC3339)
	}

	return Resource{
		ID:                bus.ID.String(),
		Name:              bus.Name.String(),
		GalaxyID:          bus.GalaxyID.String(),
		AddedAtDate:       bus.AddedAtDate.Format(time.RFC3339),
		UpdatedAtDate:     bus.UpdatedAtDate.Format(time.RFC3339),
		AddedUserID:       bus.AddedUserID.String(),
		ResourceType:      bus.ResourceType,
		UnavailableAt:     unavailableAt,
		UnavailableUserID: bus.UnavailableUserID.String(),
		Verified:          bus.Verified,
		VerifiedUserID:    bus.VerifiedUserID.String(),
		CR:                int16(bus.CR),
		CD:                int16(bus.CD),
		DR:                int16(bus.DR),
		FL:                int16(bus.FL),
		HR:                int16(bus.HR),
		MA:                int16(bus.MA),
		PE:                int16(bus.PE),
		OQ:                int16(bus.OQ),
		SR:                int16(bus.SR),
		UT:                int16(bus.UT),
		ER:                int16(bus.ER),
	}
}

func toAppResources(resources []resourcebus.Resource) []Resource {
	app := make([]Resource, len(resources))
	for i, res := range resources {
		app[i] = toAppResource(res)
	}

	return app
}

// =============================================================================

// NewResource defines the data needed to add a new resource.
type NewResource struct {
	Name         string `json:"name" validate:"required"`
	GalaxyID     string `json:"galaxyID" validate:"required"`
	AddedUserID  string `json:"addedUserID" validate:"required"`
	ResourceType string `json:"resourceType" validate:"required"`
	CR           int16  `json:"cr"`
	CD           int16  `json:"cd"`
	DR           int16  `json:"dr"`
	FL           int16  `json:"fl"`
	HR           int16  `json:"hr"`
	MA           int16  `json:"ma"`
	PE           int16  `json:"pe"`
	OQ           int16  `json:"oq" validate:"required"`
	SR           int16  `json:"sr"`
	UT           int16  `json:"ut"`
	ER           int16  `json:"er"`
}

// Decode implments the decoder interface.
func (app *NewResource) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app NewResource) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

func toBusNewResource(app NewResource) (resourcebus.NewResource, error) {

	name, err := resourcebus.Names.Parse(app.Name)
	if err != nil {
		return resourcebus.NewResource{}, fmt.Errorf("parse: %w", err)
	}

	galaxyID, err := uuid.Parse(app.GalaxyID)
	if err != nil {
		return resourcebus.NewResource{}, fmt.Errorf("parse: %w", err)
	}

	addedUserID, err := uuid.Parse(app.AddedUserID)
	if err != nil {
		return resourcebus.NewResource{}, fmt.Errorf("parse: %w", err)
	}

	bus := resourcebus.NewResource{
		Name:         name,
		GalaxyID:     galaxyID,
		AddedUserID:  addedUserID,
		ResourceType: app.ResourceType,
		CR:           app.CR,
		CD:           app.CD,
		DR:           app.DR,
		FL:           app.FL,
		HR:           app.HR,
		MA:           app.MA,
		PE:           app.PE,
		OQ:           app.OQ,
		SR:           app.SR,
		UT:           app.UT,
		ER:           app.ER,
	}

	return bus, nil
}

// =============================================================================

// UpdateResource defines the data needed to update a resource.
type UpdateResource struct {
	Name              *string    `json:"name"`
	GalaxyID          *string    `json:"galaxyID"`
	AddedUserID       *string    `json:"addedUserID"`
	ResourceType      *string    `json:"resourceType"`
	UnavailableAt     *time.Time `json:"unavailableAt"`
	UnavailableUserID *string    `json:"unavailableUserID"`
	Verified          *bool      `json:"verified"`
	VerifiedUserID    *string    `json:"verifiedUserID"`
	CR                *int16     `json:"cr"`
	CD                *int16     `json:"cd"`
	DR                *int16     `json:"dr"`
	FL                *int16     `json:"fl"`
	HR                *int16     `json:"hr"`
	MA                *int16     `json:"ma"`
	PE                *int16     `json:"pe"`
	OQ                *int16     `json:"oq"`
	SR                *int16     `json:"sr"`
	UT                *int16     `json:"ut"`
	ER                *int16     `json:"er"`
}

// Decode implments the decoder interface.
func (app *UpdateResource) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app UpdateResource) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

func toBusUpdateResource(app UpdateResource) (resourcebus.UpdateResource, error) {

	var name *resourcebus.Name
	if app.Name != nil {
		nm, err := resourcebus.Names.Parse(*app.Name)
		if err != nil {
			return resourcebus.UpdateResource{}, fmt.Errorf("parse: %w", err)
		}
		name = &nm
	}

	var unavailableUserID *uuid.UUID
	if app.UnavailableUserID != nil {
		id, err := uuid.Parse(*app.UnavailableUserID)
		if err != nil {
			return resourcebus.UpdateResource{}, fmt.Errorf("parse: %w", err)
		}
		unavailableUserID = &id
	}

	var verifiedUserID *uuid.UUID
	if app.VerifiedUserID != nil {
		id, err := uuid.Parse(*app.VerifiedUserID)
		if err != nil {
			return resourcebus.UpdateResource{}, fmt.Errorf("parse: %w", err)
		}
		verifiedUserID = &id
	}

	bus := resourcebus.UpdateResource{
		Name:              name,
		UnavailableAt:     app.UnavailableAt,
		UnavailableUserID: unavailableUserID,
		Verified:          app.Verified,
		VerifiedUserID:    verifiedUserID,
		CR:                app.CR,
		CD:                app.CD,
		DR:                app.DR,
		FL:                app.FL,
		HR:                app.HR,
		MA:                app.MA,
		PE:                app.PE,
		OQ:                app.OQ,
		SR:                app.SR,
		UT:                app.UT,
		ER:                app.ER,
	}

	return bus, nil
}

// =============================================================================

// BulkNewResources defines the data needed to bulk create resources.
type BulkNewResources struct {
	Items []NewResource `json:"items" validate:"required,min=1,max=100,dive"`
}

// Decode implements the decoder interface.
func (app *BulkNewResources) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app BulkNewResources) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

// BulkResources represents the result of a bulk resource operation.
type BulkResources struct {
	Items   []Resource `json:"items"`
	Created int        `json:"created,omitempty"`
	Updated int        `json:"updated,omitempty"`
}

// Encode implements the encoder interface.
func (app BulkResources) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

// BulkUpdateResourceItem represents a single resource update in a bulk operation.
type BulkUpdateResourceItem struct {
	ID   string         `json:"id" validate:"required,uuid"`
	Data UpdateResource `json:"data" validate:"required"`
}

// BulkUpdateResources defines the data needed to bulk update resources.
type BulkUpdateResources struct {
	Items []BulkUpdateResourceItem `json:"items" validate:"required,min=1,max=100,dive"`
}

// Decode implements the decoder interface.
func (app *BulkUpdateResources) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app BulkUpdateResources) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

// BulkDeleteResources defines the data needed to bulk delete resources.
type BulkDeleteResources struct {
	IDs []string `json:"ids" validate:"required,min=1,max=100,dive,uuid"`
}

// Decode implements the decoder interface.
func (app *BulkDeleteResources) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app BulkDeleteResources) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

// BulkDeleteResult represents the result of a bulk delete operation.
type BulkDeleteResult struct {
	Deleted int `json:"deleted"`
}

// Encode implements the encoder interface.
func (app BulkDeleteResult) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

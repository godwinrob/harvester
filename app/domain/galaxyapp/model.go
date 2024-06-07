package galaxyapp

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/foundation/validate"
)

// QueryParams represents the set of possible query strings.
type QueryParams struct {
	Page        string
	Rows        string
	OrderBy     string
	ID          string
	Name        string
	Email       string
	DateCreated string
}

// Galaxy represents information about an individual galaxy.
type Galaxy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	OwnerUserID string `json:"ownerUserID"`
	Enabled     bool   `json:"enabled"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

// Encode implments the encoder interface.
func (app Galaxy) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

func toAppGalaxy(bus galaxybus.Galaxy) Galaxy {

	return Galaxy{
		ID:          bus.ID.String(),
		Name:        bus.Name.String(),
		OwnerUserID: bus.OwnerUserID.String(),
		Enabled:     bus.Enabled,
		DateCreated: bus.DateCreated.Format(time.RFC3339),
		DateUpdated: bus.DateUpdated.Format(time.RFC3339),
	}
}

func toAppGalaxies(galaxies []galaxybus.Galaxy) []Galaxy {
	app := make([]Galaxy, len(galaxies))
	for i, gal := range galaxies {
		app[i] = toAppGalaxy(gal)
	}

	return app
}

// =============================================================================

// NewGalaxy defines the data needed to add a new galaxy.
type NewGalaxy struct {
	Name        string `json:"name" validate:"required"`
	OwnerUserID string `json:"ownerUserID" validate:"required,uuid"`
}

// Decode implments the decoder interface.
func (app *NewGalaxy) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app NewGalaxy) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

func toBusNewGalaxy(app NewGalaxy) (galaxybus.NewGalaxy, error) {

	ownerID, err := uuid.Parse(app.OwnerUserID)
	if err != nil {
		return galaxybus.NewGalaxy{}, fmt.Errorf("parse: %w", err)
	}

	name, err := galaxybus.Names.Parse(app.Name)
	if err != nil {
		return galaxybus.NewGalaxy{}, fmt.Errorf("parse: %w", err)
	}

	bus := galaxybus.NewGalaxy{
		Name:        name,
		OwnerUserID: ownerID,
	}

	return bus, nil
}

// =============================================================================

// UpdateGalaxy defines the data needed to update a galaxy.
type UpdateGalaxy struct {
	Name        *string `json:"name"`
	OwnerUserID *string `json:"ownerUserID" validate:"omitempty,uuid"`
	Enabled     *bool   `json:"enabled"`
}

// Decode implments the decoder interface.
func (app *UpdateGalaxy) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app UpdateGalaxy) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

func toBusUpdateGalaxy(app UpdateGalaxy) (galaxybus.UpdateGalaxy, error) {
	var name *galaxybus.Name
	if app.Name != nil {
		nm, err := galaxybus.Names.Parse(*app.Name)
		if err != nil {
			return galaxybus.UpdateGalaxy{}, fmt.Errorf("parse: %w", err)
		}
		name = &nm
	}

	bus := galaxybus.UpdateGalaxy{
		Name:        name,
		OwnerUserID: app.OwnerUserID,
		Enabled:     app.Enabled,
	}

	return bus, nil
}

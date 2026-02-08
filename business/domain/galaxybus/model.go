package galaxybus

import (
	"time"

	"github.com/google/uuid"
)

// Galaxy represents information about an individual galaxy.
type Galaxy struct {
	ID          uuid.UUID
	Name        Name
	OwnerUserID uuid.UUID
	Enabled     bool
	DateCreated time.Time
	DateUpdated time.Time
}

// NewGalaxy contains information needed to create a new galaxy.
type NewGalaxy struct {
	Name        Name
	OwnerUserID uuid.UUID
}

// UpdateGalaxy contains information needed to update a galaxy.
type UpdateGalaxy struct {
	Name        *Name
	OwnerUserID *uuid.UUID
	Enabled     *bool
}

// UpdateGalaxyWithID contains an ID and update data for bulk update operations.
type UpdateGalaxyWithID struct {
	ID   uuid.UUID
	Data UpdateGalaxy
}

package resourcebus

import (
	"time"

	"github.com/google/uuid"
)

// Resource represents information about an individual resource.
type Resource struct {
	ID                uuid.UUID
	Name              Name
	GalaxyID          uuid.UUID
	AddedAtDate       time.Time
	UpdatedAtDate     time.Time
	AddedUserID       uuid.UUID
	ResourceType      string
	UnavailableAt     time.Time
	UnavailableUserID uuid.UUID
	Verified          bool
	VerifiedUserID    uuid.UUID
	CR                int16
	CD                int16
	DR                int16
	FL                int16
	HR                int16
	MA                int16
	PE                int16
	OQ                int16
	SR                int16
	UT                int16
	ER                int16
}

// NewResource contains information needed to create a new resource.
type NewResource struct {
	Name         Name
	GalaxyID     uuid.UUID
	AddedUserID  uuid.UUID
	ResourceType string
	CR           int16
	CD           int16
	DR           int16
	FL           int16
	HR           int16
	MA           int16
	PE           int16
	OQ           int16
	SR           int16
	UT           int16
	ER           int16
}

// UpdateResource contains information needed to update a resource.
type UpdateResource struct {
	Name              *Name
	UnavailableAt     *time.Time
	UnavailableUserID *uuid.UUID
	Verified          *bool
	VerifiedUserID    *uuid.UUID
	CR                *int16
	CD                *int16
	DR                *int16
	FL                *int16
	HR                *int16
	MA                *int16
	PE                *int16
	OQ                *int16
	SR                *int16
	UT                *int16
	ER                *int16
}

// UpdateResourceWithID contains an ID and update data for bulk update operations.
type UpdateResourceWithID struct {
	ID   uuid.UUID
	Data UpdateResource
}

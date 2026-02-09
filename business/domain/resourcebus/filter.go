package resourcebus

import (
	"time"

	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
// We are using pointer semantics because the With API mutates the value.
type QueryFilter struct {
	ID               *uuid.UUID
	GalaxyID         *uuid.UUID
	ResourceName     *Name
	ResourceType     *string
	ResourceGroup    *string
	StartCreatedDate *time.Time
	EndCreatedDate   *time.Time
	Verified         *bool
	CR               *int16
	CD               *int16
	DR               *int16
	FL               *int16
	HR               *int16
	MA               *int16
	PE               *int16
	OQ               *int16
	SR               *int16
	UT               *int16
	ER               *int16
}

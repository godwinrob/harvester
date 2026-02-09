package userbus

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
)

// User represents information about an individual user.
type User struct {
	ID           uuid.UUID
	Name         Name
	Email        mail.Address
	Roles        []Role
	PasswordHash []byte
	Guild        string
	Enabled      bool
	DateCreated  time.Time
	DateUpdated  time.Time
}

// NewUser contains information needed to create a new user.
type NewUser struct {
	Name     Name
	Email    mail.Address
	Roles    []Role
	Guild    string
	Password string
}

// UpdateUser contains information needed to update a user.
type UpdateUser struct {
	Name     *Name
	Email    *mail.Address
	Roles    []Role
	Guild    *string
	Password *string
	Enabled  *bool
}

// UpdateUserWithID contains an ID and update data for bulk update operations.
type UpdateUserWithID struct {
	ID   uuid.UUID
	Data UpdateUser
}

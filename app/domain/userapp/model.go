package userapp

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"time"

	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/business/domain/userbus"
	"github.com/godwinrob/harvester/foundation/validate"
)

// QueryParams represents the set of possible query strings.
type QueryParams struct {
	Page             string
	Rows             string
	OrderBy          string
	ID               string
	Name             string
	Email            string
	StartCreatedDate string
	EndCreatedDate   string
}

// User represents information about an individual user.
type User struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Roles        []string `json:"roles"`
	PasswordHash []byte   `json:"-"`
	Guild        string   `json:"guild"`
	Enabled      bool     `json:"enabled"`
	DateCreated  string   `json:"dateCreated"`
	DateUpdated  string   `json:"dateUpdated"`
}

// Encode implments the encoder interface.
func (app User) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

func toAppUser(bus userbus.User) User {
	roles := make([]string, len(bus.Roles))
	for i, role := range bus.Roles {
		roles[i] = role.String()
	}

	return User{
		ID:           bus.ID.String(),
		Name:         bus.Name.String(),
		Email:        bus.Email.Address,
		Roles:        roles,
		PasswordHash: bus.PasswordHash,
		Guild:        bus.Guild,
		Enabled:      bus.Enabled,
		DateCreated:  bus.DateCreated.Format(time.RFC3339),
		DateUpdated:  bus.DateUpdated.Format(time.RFC3339),
	}
}

func toAppUsers(users []userbus.User) []User {
	app := make([]User, len(users))
	for i, usr := range users {
		app[i] = toAppUser(usr)
	}

	return app
}

// =============================================================================

// NewUser defines the data needed to add a new user.
type NewUser struct {
	Name            string   `json:"name" validate:"required"`
	Email           string   `json:"email" validate:"required,email"`
	Roles           []string `json:"roles" validate:"required"`
	Guild           string   `json:"guild"`
	Password        string   `json:"password" validate:"required"`
	PasswordConfirm string   `json:"passwordConfirm" validate:"eqfield=Password"`
}

// Decode implments the decoder interface.
func (app *NewUser) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app NewUser) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

func toBusNewUser(app NewUser) (userbus.NewUser, error) {
	roles := make([]userbus.Role, len(app.Roles))
	for i, roleStr := range app.Roles {
		role, err := userbus.Roles.Parse(roleStr)
		if err != nil {
			return userbus.NewUser{}, fmt.Errorf("parse: %w", err)
		}
		roles[i] = role
	}

	addr, err := mail.ParseAddress(app.Email)
	if err != nil {
		return userbus.NewUser{}, fmt.Errorf("parse: %w", err)
	}

	name, err := userbus.Names.Parse(app.Name)
	if err != nil {
		return userbus.NewUser{}, fmt.Errorf("parse: %w", err)
	}

	bus := userbus.NewUser{
		Name:     name,
		Email:    *addr,
		Roles:    roles,
		Guild:    app.Guild,
		Password: app.Password,
	}

	return bus, nil
}

// =============================================================================

// UpdateUserRole defines the data needed to update a user role.
type UpdateUserRole struct {
	Roles []string `json:"roles" validate:"required"`
}

// Decode implments the decoder interface.
func (app *UpdateUserRole) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app UpdateUserRole) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

func toBusUpdateUserRole(app UpdateUserRole) (userbus.UpdateUser, error) {
	var roles []userbus.Role
	if app.Roles != nil {
		roles = make([]userbus.Role, len(app.Roles))
		for i, roleStr := range app.Roles {
			role, err := userbus.Roles.Parse(roleStr)
			if err != nil {
				return userbus.UpdateUser{}, fmt.Errorf("parse: %w", err)
			}
			roles[i] = role
		}
	}

	bus := userbus.UpdateUser{
		Roles: roles,
	}

	return bus, nil
}

// =============================================================================

// UpdateUser defines the data needed to update a user.
type UpdateUser struct {
	Name            *string `json:"name"`
	Email           *string `json:"email" validate:"omitempty,email"`
	Guild           *string `json:"guild"`
	Password        *string `json:"password"`
	PasswordConfirm *string `json:"passwordConfirm" validate:"omitempty,eqfield=Password"`
	Enabled         *bool   `json:"enabled"`
}

// Decode implments the decoder interface.
func (app *UpdateUser) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app UpdateUser) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

func toBusUpdateUser(app UpdateUser) (userbus.UpdateUser, error) {
	var addr *mail.Address
	if app.Email != nil {
		var err error
		addr, err = mail.ParseAddress(*app.Email)
		if err != nil {
			return userbus.UpdateUser{}, fmt.Errorf("parse: %w", err)
		}
	}

	var name *userbus.Name
	if app.Name != nil {
		nm, err := userbus.Names.Parse(*app.Name)
		if err != nil {
			return userbus.UpdateUser{}, fmt.Errorf("parse: %w", err)
		}
		name = &nm
	}

	bus := userbus.UpdateUser{
		Name:     name,
		Email:    addr,
		Guild:    app.Guild,
		Password: app.Password,
		Enabled:  app.Enabled,
	}

	return bus, nil
}

// =============================================================================

// BulkNewUsers defines the data needed to bulk create users.
type BulkNewUsers struct {
	Items []NewUser `json:"items" validate:"required,min=1,max=100,dive"`
}

// Decode implements the decoder interface.
func (app *BulkNewUsers) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app BulkNewUsers) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

// BulkUsers represents the result of a bulk user operation.
type BulkUsers struct {
	Items   []User `json:"items"`
	Created int    `json:"created,omitempty"`
	Updated int    `json:"updated,omitempty"`
}

// Encode implements the encoder interface.
func (app BulkUsers) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

// BulkUpdateUserItem represents a single user update in a bulk operation.
type BulkUpdateUserItem struct {
	ID   string     `json:"id" validate:"required,uuid"`
	Data UpdateUser `json:"data" validate:"required"`
}

// BulkUpdateUsers defines the data needed to bulk update users.
type BulkUpdateUsers struct {
	Items []BulkUpdateUserItem `json:"items" validate:"required,min=1,max=100,dive"`
}

// Decode implements the decoder interface.
func (app *BulkUpdateUsers) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app BulkUpdateUsers) Validate() error {
	if err := validate.Check(app); err != nil {
		return errs.Newf(errs.FailedPrecondition, "validate: %s", err)
	}

	return nil
}

// BulkDeleteUsers defines the data needed to bulk delete users.
type BulkDeleteUsers struct {
	IDs []string `json:"ids" validate:"required,min=1,max=100,dive,uuid"`
}

// Decode implements the decoder interface.
func (app *BulkDeleteUsers) Decode(data []byte) error {
	return json.Unmarshal(data, &app)
}

// Validate checks the data in the model is considered clean.
func (app BulkDeleteUsers) Validate() error {
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

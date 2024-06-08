package resourcedb

import (
	"fmt"
	"time"

	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/google/uuid"
)

type resource struct {
	ID                uuid.UUID `db:"resource_id"`
	ResourceName      string    `db:"resource_name"`
	GalaxyID          uuid.UUID `db:"galaxy_id"`
	AddedAtDate       time.Time `db:"added_at"`
	UpdatedAtDate     time.Time `db:"updated_at"`
	AddedUserID       uuid.UUID `db:"added_user_id"`
	ResourceType      string    `db:"resource_type"`
	UnavailableAt     time.Time `db:"unavailable_at"`
	UnavailableUserID uuid.UUID `db:"unavailable_user_id"`
	Verified          bool      `db:"verified"`
	VerifiedUserID    uuid.UUID `db:"verified_user_id"`
	CR                int16     `db:"cr"`
	CD                int16     `db:"cd"`
	DR                int16     `db:"dr"`
	FL                int16     `db:"fl"`
	HR                int16     `db:"hr"`
	MA                int16     `db:"ma"`
	PE                int16     `db:"pe"`
	OQ                int16     `db:"oq"`
	SR                int16     `db:"sr"`
	UT                int16     `db:"ut"`
	ER                int16     `db:"er"`
}

func toDBResource(bus resourcebus.Resource) resource {

	return resource{
		ID:                bus.ID,
		ResourceName:      bus.Name.String(),
		GalaxyID:          bus.GalaxyID,
		AddedAtDate:       bus.AddedAtDate,
		UpdatedAtDate:     bus.UpdatedAtDate,
		AddedUserID:       bus.AddedUserID,
		ResourceType:      bus.ResourceType,
		UnavailableAt:     bus.UnavailableAt,
		UnavailableUserID: bus.UnavailableUserID,
		Verified:          bus.Verified,
		VerifiedUserID:    bus.VerifiedUserID,
		CR:                bus.CR,
		CD:                bus.CD,
		DR:                bus.DR,
		FL:                bus.FL,
		HR:                bus.HR,
		MA:                bus.MA,
		PE:                bus.PE,
		OQ:                bus.OQ,
		SR:                bus.SR,
		UT:                bus.UT,
		ER:                bus.ER,
	}
}

func toBusResource(db resource) (resourcebus.Resource, error) {

	name, err := resourcebus.Names.Parse(db.ResourceName)
	if err != nil {
		return resourcebus.Resource{}, fmt.Errorf("parse name: %w", err)
	}

	bus := resourcebus.Resource{
		ID:                db.ID,
		Name:              name,
		GalaxyID:          db.GalaxyID,
		AddedAtDate:       db.AddedAtDate,
		UpdatedAtDate:     db.UpdatedAtDate,
		AddedUserID:       db.AddedUserID,
		ResourceType:      db.ResourceType,
		UnavailableAt:     db.UnavailableAt,
		UnavailableUserID: db.UnavailableUserID,
		Verified:          db.Verified,
		VerifiedUserID:    db.VerifiedUserID,
		CR:                db.CR,
		CD:                db.CD,
		DR:                db.DR,
		FL:                db.FL,
		HR:                db.HR,
		MA:                db.MA,
		PE:                db.PE,
		OQ:                db.OQ,
		SR:                db.SR,
		UT:                db.UT,
		ER:                db.ER,
	}

	return bus, nil
}

func toBusResources(dbs []resource) ([]resourcebus.Resource, error) {
	bus := make([]resourcebus.Resource, len(dbs))

	for i, db := range dbs {
		var err error
		bus[i], err = toBusResource(db)
		if err != nil {
			return nil, err
		}
	}

	return bus, nil
}

package galaxydb

import (
	"fmt"
	"time"

	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/google/uuid"
)

type galaxy struct {
	ID          uuid.UUID `db:"galaxy_id"`
	Name        string    `db:"galaxy_name"`
	OwnerUserID uuid.UUID `db:"owner_user_id"`
	Enabled     bool      `db:"enabled"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBGalaxy(bus galaxybus.Galaxy) galaxy {

	return galaxy{
		ID:          bus.ID,
		Name:        bus.Name.String(),
		OwnerUserID: bus.OwnerUserID,
		Enabled:     bus.Enabled,
		DateCreated: bus.DateCreated.UTC(),
		DateUpdated: bus.DateUpdated.UTC(),
	}
}

func toBusGalaxy(db galaxy) (galaxybus.Galaxy, error) {

	name, err := galaxybus.Names.Parse(db.Name)
	if err != nil {
		return galaxybus.Galaxy{}, fmt.Errorf("parse name: %w", err)
	}

	bus := galaxybus.Galaxy{
		ID:          db.ID,
		Name:        name,
		OwnerUserID: db.OwnerUserID,
		Enabled:     db.Enabled,
		DateCreated: db.DateCreated.In(time.Local),
		DateUpdated: db.DateUpdated.In(time.Local),
	}

	return bus, nil
}

func toBusGalaxies(dbs []galaxy) ([]galaxybus.Galaxy, error) {
	bus := make([]galaxybus.Galaxy, len(dbs))

	for i, db := range dbs {
		var err error
		bus[i], err = toBusGalaxy(db)
		if err != nil {
			return nil, err
		}
	}

	return bus, nil
}

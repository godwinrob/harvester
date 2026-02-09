package resourcetypedb

import (
	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
)

type resourceType struct {
	ResourceType     string `db:"resource_type"`
	ResourceTypeName string `db:"resource_type_name"`
	ResourceCategory string `db:"resource_category"`
	ResourceGroup    string `db:"resource_group"`
	Enterable        bool   `db:"enterable"`
	MaxTypes         int16  `db:"max_types"`
	CRmin            int16  `db:"cr_min"`
	CRmax            int16  `db:"cr_max"`
	CDmin            int16  `db:"cd_min"`
	CDmax            int16  `db:"cd_max"`
	DRmin            int16  `db:"dr_min"`
	DRmax            int16  `db:"dr_max"`
	FLmin            int16  `db:"fl_min"`
	FLmax            int16  `db:"fl_max"`
	HRmin            int16  `db:"hr_min"`
	HRmax            int16  `db:"hr_max"`
	MAmin            int16  `db:"ma_min"`
	MAmax            int16  `db:"ma_max"`
	PEmin            int16  `db:"pe_min"`
	PEmax            int16  `db:"pe_max"`
	OQmin            int16  `db:"oq_min"`
	OQmax            int16  `db:"oq_max"`
	SRmin            int16  `db:"sr_min"`
	SRmax            int16  `db:"sr_max"`
	UTmin            int16  `db:"ut_min"`
	UTmax            int16  `db:"ut_max"`
	ERmin            int16  `db:"er_min"`
	ERmax            int16  `db:"er_max"`
	ContainerType    string `db:"container_type"`
	InventoryType    string `db:"inventory_type"`
	SpecificPlanet   int16  `db:"specific_planet"`
}

func toDBResourceType(bus resourcetypebus.ResourceType) resourceType {
	return resourceType{
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

func toBusResourceType(db resourceType) resourcetypebus.ResourceType {
	return resourcetypebus.ResourceType{
		ResourceType:     db.ResourceType,
		ResourceTypeName: db.ResourceTypeName,
		ResourceCategory: db.ResourceCategory,
		ResourceGroup:    db.ResourceGroup,
		Enterable:        db.Enterable,
		MaxTypes:         db.MaxTypes,
		CRmin:            db.CRmin,
		CRmax:            db.CRmax,
		CDmin:            db.CDmin,
		CDmax:            db.CDmax,
		DRmin:            db.DRmin,
		DRmax:            db.DRmax,
		FLmin:            db.FLmin,
		FLmax:            db.FLmax,
		HRmin:            db.HRmin,
		HRmax:            db.HRmax,
		MAmin:            db.MAmin,
		MAmax:            db.MAmax,
		PEmin:            db.PEmin,
		PEmax:            db.PEmax,
		OQmin:            db.OQmin,
		OQmax:            db.OQmax,
		SRmin:            db.SRmin,
		SRmax:            db.SRmax,
		UTmin:            db.UTmin,
		UTmax:            db.UTmax,
		ERmin:            db.ERmin,
		ERmax:            db.ERmax,
		ContainerType:    db.ContainerType,
		InventoryType:    db.InventoryType,
		SpecificPlanet:   db.SpecificPlanet,
	}
}

func toBusResourceTypes(dbs []resourceType) []resourcetypebus.ResourceType {
	bus := make([]resourcetypebus.ResourceType, len(dbs))
	for i, db := range dbs {
		bus[i] = toBusResourceType(db)
	}
	return bus
}

package resourcedb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/godwinrob/harvester/business/domain/resourcebus"
)

func applyFilter(filter resourcebus.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["resource_id"] = *filter.ID
		wc = append(wc, "resource_id = :resource_id")
	}

	if filter.GalaxyID != nil {
		data["galaxy_id"] = *filter.GalaxyID
		wc = append(wc, "galaxy_id = :galaxy_id")
	}

	if filter.ResourceName != nil {
		data["resource_name"] = fmt.Sprintf("%%%s%%", *filter.ResourceName)
		wc = append(wc, "resource_name LIKE :resource_name")
	}

	if filter.StartCreatedDate != nil {
		data["start_date_created"] = filter.StartCreatedDate.UTC()
		wc = append(wc, "updated_at >= :start_date_created")
	}

	if filter.EndCreatedDate != nil {
		data["end_date_created"] = filter.EndCreatedDate.UTC()
		wc = append(wc, "updated_at <= :end_date_created")
	}

	if filter.Verified != nil {
		data["verified"] = *filter.Verified
		wc = append(wc, "verified = :verified")
	}

	if filter.ResourceType != nil {
		data["resource_type"] = *filter.ResourceType
		wc = append(wc, "resource_type = :resource_type")
	}

	if filter.ResourceGroup != nil {
		data["resource_group"] = *filter.ResourceGroup
		wc = append(wc, "resource_type IN (SELECT resource_type FROM resource_type_groups WHERE resource_group = :resource_group)")
	}

	if filter.CR != nil {
		data["cr"] = *filter.CR
		wc = append(wc, "cr >= :cr")
	}

	if filter.CD != nil {
		data["cd"] = *filter.CD
		wc = append(wc, "cd >= :cd")
	}

	if filter.DR != nil {
		data["dr"] = *filter.DR
		wc = append(wc, "dr >= :dr")
	}

	if filter.FL != nil {
		data["fl"] = *filter.FL
		wc = append(wc, "fl >= :fl")
	}

	if filter.HR != nil {
		data["hr"] = *filter.HR
		wc = append(wc, "hr >= :hr")
	}

	if filter.MA != nil {
		data["ma"] = *filter.MA
		wc = append(wc, "ma >= :ma")
	}

	if filter.PE != nil {
		data["pe"] = *filter.PE
		wc = append(wc, "pe >= :pe")
	}

	if filter.OQ != nil {
		data["oq"] = *filter.OQ
		wc = append(wc, "oq >= :oq")
	}

	if filter.SR != nil {
		data["sr"] = *filter.SR
		wc = append(wc, "sr >= :sr")
	}

	if filter.UT != nil {
		data["ut"] = *filter.UT
		wc = append(wc, "ut >= :ut")
	}

	if filter.ER != nil {
		data["er"] = *filter.ER
		wc = append(wc, "er >= :er")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}

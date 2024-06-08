package galaxydb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/godwinrob/harvester/business/domain/galaxybus"
)

func applyFilter(filter galaxybus.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["galaxy_id"] = *filter.ID
		wc = append(wc, "galaxy_id = :galaxy_id")
	}

	if filter.Name != nil {
		data["name"] = fmt.Sprintf("%%%s%%", *filter.Name)
		wc = append(wc, "name LIKE :name")
	}

	if filter.CreatedDate != nil {
		data["date_created"] = filter.CreatedDate.UTC()
		wc = append(wc, "date_created >= :date_created")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}

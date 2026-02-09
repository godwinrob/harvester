package resourcetypedb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
)

func applyFilter(filter resourcetypebus.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ResourceType != nil {
		data["resource_type"] = fmt.Sprintf("%%%s%%", *filter.ResourceType)
		wc = append(wc, "resource_type LIKE :resource_type")
	}

	if filter.ResourceTypeName != nil {
		data["resource_type_name"] = fmt.Sprintf("%%%s%%", *filter.ResourceTypeName)
		wc = append(wc, "resource_type_name LIKE :resource_type_name")
	}

	if filter.ResourceCategory != nil {
		data["resource_category"] = *filter.ResourceCategory
		wc = append(wc, "resource_category = :resource_category")
	}

	if filter.ResourceGroup != nil {
		data["resource_group"] = *filter.ResourceGroup
		wc = append(wc, "resource_group = :resource_group")
	}

	if filter.Enterable != nil {
		data["enterable"] = *filter.Enterable
		wc = append(wc, "enterable = :enterable")
	}

	if filter.ContainerType != nil {
		data["container_type"] = *filter.ContainerType
		wc = append(wc, "container_type = :container_type")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}

package resourcegroupdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
)

func applyFilter(filter resourcegroupbus.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.ResourceGroup != nil {
		data["resource_group"] = fmt.Sprintf("%%%s%%", *filter.ResourceGroup)
		wc = append(wc, "resource_group LIKE :resource_group")
	}

	if filter.GroupName != nil {
		data["group_name"] = fmt.Sprintf("%%%s%%", *filter.GroupName)
		wc = append(wc, "group_name LIKE :group_name")
	}

	if filter.GroupLevel != nil {
		data["group_level"] = *filter.GroupLevel
		wc = append(wc, "group_level = :group_level")
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

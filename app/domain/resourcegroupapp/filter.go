package resourcegroupapp

import (
	"strconv"

	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
	"github.com/godwinrob/harvester/foundation/validate"
)

func parseFilter(qp QueryParams) (resourcegroupbus.QueryFilter, error) {
	var filter resourcegroupbus.QueryFilter

	if qp.ResourceGroup != "" {
		filter.ResourceGroup = &qp.ResourceGroup
	}

	if qp.GroupName != "" {
		filter.GroupName = &qp.GroupName
	}

	if qp.GroupLevel != "" {
		level, err := strconv.ParseInt(qp.GroupLevel, 10, 16)
		if err != nil {
			return resourcegroupbus.QueryFilter{}, validate.NewFieldsError("groupLevel", err)
		}
		l := int16(level)
		filter.GroupLevel = &l
	}

	if qp.ContainerType != "" {
		filter.ContainerType = &qp.ContainerType
	}

	return filter, nil
}

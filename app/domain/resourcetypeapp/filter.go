package resourcetypeapp

import (
	"strconv"

	"github.com/godwinrob/harvester/business/domain/resourcetypebus"
	"github.com/godwinrob/harvester/foundation/validate"
)

func parseFilter(qp QueryParams) (resourcetypebus.QueryFilter, error) {
	var filter resourcetypebus.QueryFilter

	if qp.ResourceType != "" {
		filter.ResourceType = &qp.ResourceType
	}

	if qp.ResourceTypeName != "" {
		filter.ResourceTypeName = &qp.ResourceTypeName
	}

	if qp.ResourceCategory != "" {
		filter.ResourceCategory = &qp.ResourceCategory
	}

	if qp.ResourceGroup != "" {
		filter.ResourceGroup = &qp.ResourceGroup
	}

	if qp.Enterable != "" {
		enterable, err := strconv.ParseBool(qp.Enterable)
		if err != nil {
			return resourcetypebus.QueryFilter{}, validate.NewFieldsError("enterable", err)
		}
		filter.Enterable = &enterable
	}

	if qp.ContainerType != "" {
		filter.ContainerType = &qp.ContainerType
	}

	return filter, nil
}

package resourceapp

import (
	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/foundation/validate"
	"github.com/google/uuid"
)

func parseFilter(qp QueryParams) (resourcebus.QueryFilter, error) {
	var filter resourcebus.QueryFilter

	if qp.ID != "" {
		id, err := uuid.Parse(qp.ID)
		if err != nil {
			return resourcebus.QueryFilter{}, validate.NewFieldsError("resource_id", err)
		}
		filter.ID = &id
	}

	if qp.Name != "" {
		name, err := resourcebus.Names.Parse(qp.Name)
		if err != nil {
			return resourcebus.QueryFilter{}, validate.NewFieldsError("name", err)
		}
		filter.ResourceName = &name
	}

	if qp.ResourceType != "" {
		filter.ResourceType = &qp.ResourceType
	}

	if qp.ResourceGroup != "" {
		filter.ResourceGroup = &qp.ResourceGroup
	}

	//if qp.StartCreatedDate != "" {
	//	t, err := time.Parse(time.RFC3339, qp.StartCreatedDate)
	//	if err != nil {
	//		return userbus.QueryFilter{}, validate.NewFieldsError("added_at", err)
	//	}
	//	filter.AddedAtDate = &t
	//}

	return filter, nil
}

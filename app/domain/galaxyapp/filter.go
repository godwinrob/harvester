package galaxyapp

import (
	"github.com/godwinrob/harvester/business/domain/galaxybus"
	"github.com/godwinrob/harvester/foundation/validate"
	"github.com/google/uuid"
)

func parseFilter(qp QueryParams) (galaxybus.QueryFilter, error) {
	var filter galaxybus.QueryFilter

	if qp.ID != "" {
		id, err := uuid.Parse(qp.ID)
		if err != nil {
			return galaxybus.QueryFilter{}, validate.NewFieldsError("galaxy_id", err)
		}
		filter.ID = &id
	}

	if qp.Name != "" {
		name, err := galaxybus.Names.Parse(qp.Name)
		if err != nil {
			return galaxybus.QueryFilter{}, validate.NewFieldsError("galaxy_name", err)
		}
		filter.Name = &name
	}

	//if qp.StartCreatedDate != "" {
	//	t, err := time.Parse(time.RFC3339, qp.StartCreatedDate)
	//	if err != nil {
	//		return galaxybus.QueryFilter{}, validate.NewFieldsError("start_created_date", err)
	//	}
	//	filter.StartCreatedDate = &t
	//}
	//
	//if qp.EndCreatedDate != "" {
	//	t, err := time.Parse(time.RFC3339, qp.EndCreatedDate)
	//	if err != nil {
	//		return galaxybus.QueryFilter{}, validate.NewFieldsError("end_created_date", err)
	//	}
	//	filter.EndCreatedDate = &t
	//}

	return filter, nil
}

package galaxyapi

import (
	"github.com/godwinrob/harvester/app/domain/galaxyapp"
	"net/http"
)

func parseQueryParams(r *http.Request) (galaxyapp.QueryParams, error) {
	values := r.URL.Query()

	filter := galaxyapp.QueryParams{
		Page:        values.Get("page"),
		Rows:        values.Get("row"),
		OrderBy:     values.Get("orderBy"),
		ID:          values.Get("galaxy_id"),
		Name:        values.Get("name"),
		DateCreated: values.Get("date_created"),
	}

	return filter, nil
}

package resourcetypeapi

import (
	"net/http"

	"github.com/godwinrob/harvester/app/domain/resourcetypeapp"
)

func parseQueryParams(r *http.Request) (resourcetypeapp.QueryParams, error) {
	values := r.URL.Query()

	filter := resourcetypeapp.QueryParams{
		Page:             values.Get("page"),
		Rows:             values.Get("row"),
		OrderBy:          values.Get("orderBy"),
		ResourceType:     values.Get("resourceType"),
		ResourceTypeName: values.Get("resourceTypeName"),
		ResourceCategory: values.Get("resourceCategory"),
		ResourceGroup:    values.Get("resourceGroup"),
		Enterable:        values.Get("enterable"),
		ContainerType:    values.Get("containerType"),
	}

	return filter, nil
}

package resourcegroupapi

import (
	"net/http"

	"github.com/godwinrob/harvester/app/domain/resourcegroupapp"
)

func parseQueryParams(r *http.Request) (resourcegroupapp.QueryParams, error) {
	values := r.URL.Query()

	filter := resourcegroupapp.QueryParams{
		Page:          values.Get("page"),
		Rows:          values.Get("row"),
		OrderBy:       values.Get("orderBy"),
		ResourceGroup: values.Get("resourceGroup"),
		GroupName:     values.Get("groupName"),
		GroupLevel:    values.Get("groupLevel"),
		ContainerType: values.Get("containerType"),
	}

	return filter, nil
}

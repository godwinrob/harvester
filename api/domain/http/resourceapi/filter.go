package resourceapi

import (
	"github.com/godwinrob/harvester/app/domain/resourceapp"
	"net/http"
)

func parseQueryParams(r *http.Request) (resourceapp.QueryParams, error) {
	values := r.URL.Query()

	filter := resourceapp.QueryParams{
		Page:         values.Get("page"),
		Rows:         values.Get("row"),
		OrderBy:      values.Get("orderBy"),
		ID:           values.Get("resource_id"),
		Name:         values.Get("name"),
		AddedAt:      values.Get("added_at"),
		ResourceType: values.Get("resource_type"),
		CR:           values.Get("cr"),
		CD:           values.Get("cd"),
		DR:           values.Get("dr"),
		FL:           values.Get("fl"),
		HR:           values.Get("hr"),
		MA:           values.Get("ma"),
		PE:           values.Get("pe"),
		OQ:           values.Get("oq"),
		SR:           values.Get("sr"),
		UT:           values.Get("ut"),
		ER:           values.Get("er"),
	}

	return filter, nil
}

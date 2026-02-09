package resourcebus

import "github.com/godwinrob/harvester/business/sdk/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByID, order.ASC)

// Set of fields that the results can be ordered by.
const (
	OrderByID           = "resource_id"
	OrderByName         = "resource_name"
	OrderByResourceType = "resource_type"
	OrderByVerified     = "verified"
	OrderByUnavailableAt = "unavailable_at"
	OrderByAddedAt      = "added_at"
	OrderByEnabled      = "enabled"
	OrderByCR           = "cr"
	OrderByCD           = "cd"
	OrderByDR           = "dr"
	OrderByFL           = "fl"
	OrderByHR           = "hr"
	OrderByMA           = "ma"
	OrderByPE           = "pe"
	OrderByOQ           = "oq"
	OrderBySR           = "sr"
	OrderByUT           = "ut"
	OrderByER           = "er"
)

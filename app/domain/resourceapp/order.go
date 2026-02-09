package resourceapp

import (
	"github.com/godwinrob/harvester/business/domain/resourcebus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var defaultOrderBy = order.NewBy(resourcebus.OrderByID, order.ASC)

var orderByFields = map[string]string{
	"resource_id":    resourcebus.OrderByID,
	"name":           resourcebus.OrderByName,
	"resource_name":  resourcebus.OrderByName,
	"resourceType":   resourcebus.OrderByResourceType,
	"resource_type":  resourcebus.OrderByResourceType,
	"verified":       resourcebus.OrderByVerified,
	"unavailableAt":  resourcebus.OrderByUnavailableAt,
	"unavailable_at": resourcebus.OrderByUnavailableAt,
	"addedAtDate":    resourcebus.OrderByAddedAt,
	"added_at":       resourcebus.OrderByAddedAt,
	"cr":             resourcebus.OrderByCR,
	"cd":             resourcebus.OrderByCD,
	"dr":             resourcebus.OrderByDR,
	"fl":             resourcebus.OrderByFL,
	"hr":             resourcebus.OrderByHR,
	"ma":             resourcebus.OrderByMA,
	"pe":             resourcebus.OrderByPE,
	"oq":             resourcebus.OrderByOQ,
	"sr":             resourcebus.OrderBySR,
	"ut":             resourcebus.OrderByUT,
	"er":             resourcebus.OrderByER,
}

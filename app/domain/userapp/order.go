package userapp

import (
	"github.com/godwinrob/harvester/business/domain/userbus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

var defaultOrderBy = order.NewBy("user_id", order.ASC)

var orderByFields = map[string]string{
	"user_id":      userbus.OrderByID,
	"name":         userbus.OrderByName,
	"email":        userbus.OrderByEmail,
	"roles":        userbus.OrderByRoles,
	"guild":        userbus.OrderByGuild,
	"dateCreated":  userbus.OrderByDateCreated,
	"date_created": userbus.OrderByDateCreated,
	"enabled":      userbus.OrderByEnabled,
}

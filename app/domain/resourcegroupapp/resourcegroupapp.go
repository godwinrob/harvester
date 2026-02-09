package resourcegroupapp

import (
	"context"

	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/app/sdk/page"
	"github.com/godwinrob/harvester/business/domain/resourcegroupbus"
	"github.com/godwinrob/harvester/business/sdk/order"
)

// App manages the set of app layer api functions for the resource group domain.
type App struct {
	resourceGroupBus *resourcegroupbus.Business
}

// NewApp constructs a resource group app API for use.
func NewApp(resourceGroupBus *resourcegroupbus.Business) *App {
	return &App{
		resourceGroupBus: resourceGroupBus,
	}
}

// Query returns a list of resource groups with paging.
func (a *App) Query(ctx context.Context, qp QueryParams) (page.Document[ResourceGroup], error) {
	pg, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return page.Document[ResourceGroup]{}, err
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return page.Document[ResourceGroup]{}, err
	}

	orderBy, err := order.Parse(orderByFields, qp.OrderBy, defaultOrderBy)
	if err != nil {
		return page.Document[ResourceGroup]{}, err
	}

	groups, err := a.resourceGroupBus.Query(ctx, filter, orderBy, pg.Number, pg.RowsPerPage)
	if err != nil {
		return page.Document[ResourceGroup]{}, errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.resourceGroupBus.Count(ctx, filter)
	if err != nil {
		return page.Document[ResourceGroup]{}, errs.Newf(errs.Internal, "count: %s", err)
	}

	return page.NewDocument(toAppResourceGroups(groups), total, pg.Number, pg.RowsPerPage), nil
}

// QueryByID returns a resource group by its key.
func (a *App) QueryByID(ctx context.Context, resourceGroup string) (ResourceGroup, error) {
	group, err := a.resourceGroupBus.QueryByID(ctx, resourceGroup)
	if err != nil {
		return ResourceGroup{}, errs.Newf(errs.Internal, "resource group missing: %s", err)
	}

	return toAppResourceGroup(group), nil
}

package mid

import (
	"context"
	"net/http"

	"github.com/godwinrob/harvester/app/sdk/mid"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
)

// Errors executes the errors middleware functionality.
func Error(log *logger.Logger) web.Middleware {
	midFunc := func(ctx context.Context, r *http.Request, next mid.Handler) (mid.Encoder, error) {
		return mid.Errors(ctx, log, next)
	}

	return addMiddleware(midFunc)
}

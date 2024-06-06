package mid

import (
	"context"
	"net/http"

	"github.com/godwinrob/harvester/app/sdk/mid"
	"github.com/godwinrob/harvester/foundation/logger"
	"github.com/godwinrob/harvester/foundation/web"
)

// Logger executes the logger middleware functionality.
func Logger(log *logger.Logger) web.Middleware {
	midFunc := func(ctx context.Context, r *http.Request, next mid.Handler) (mid.Encoder, error) {
		return mid.Logger(ctx, log, r.URL.Path, r.URL.RawQuery, r.Method, r.RemoteAddr, next)
	}

	return addMiddleware(midFunc)
}

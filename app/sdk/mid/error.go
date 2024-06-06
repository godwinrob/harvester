package mid

import (
	"context"
	"path"

	"github.com/godwinrob/harvester/app/sdk/errs"
	"github.com/godwinrob/harvester/foundation/logger"
)

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a uniform way.
// Unexpected errors (status >= 500) are logged.
func Errors(ctx context.Context, log *logger.Logger, next Handler) (Encoder, error) {
	resp, err := next(ctx)
	if err == nil {
		return resp, nil
	}

	v, ok := err.(*errs.Error)
	if !ok {
		v = errs.New(errs.Internal, err)
		err = v
	}

	log.Error(ctx, "message", "ERROR", err, "FileName", path.Base(v.FileName), "FuncName", path.Base(v.FuncName))

	// Send the error to the web package so the error can be
	// used as the response.

	return nil, err
}

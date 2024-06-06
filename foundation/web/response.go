package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type httpStatus interface {
	HTTPStatus() int
}

func respondError(ctx context.Context, w http.ResponseWriter, err error) error {
	data, ok := err.(Encoder)
	if !ok {
		return fmt.Errorf("error value does not implement the encoder interface: %T", err)
	}

	return respond(ctx, w, data)
}

func respond(ctx context.Context, w http.ResponseWriter, data Encoder) error {

	// If the context has been canceled, it means the client is no longer
	// waiting for a response.
	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return errors.New("client disconnected, do not send response")
		}
	}

	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	var statusCode = http.StatusOK
	switch v := data.(type) {
	case httpStatus:
		statusCode = v.HTTPStatus()
	case error:
		statusCode = http.StatusInternalServerError
	}

	b, contentType, err := data.Encode()
	if err != nil {
		return fmt.Errorf("respond: encode: %w", err)
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)

	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("respond: write: %w", err)
	}

	return nil
}

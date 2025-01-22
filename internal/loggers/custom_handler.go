package loggers

import (
	"context"
	"log/slog"
)

type ContextKey string

const RequestIDCtxKey ContextKey = "request_id"

type CustomHandler struct {
	slog.Handler
}

func NewCustomHandler(handler slog.Handler) *CustomHandler {
	return &CustomHandler{Handler: handler}
}

func (h CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestID, ok := ctx.Value(RequestIDCtxKey).(string); ok {
		r.Add(slog.String("request_id", requestID))
	}

	return h.Handler.Handle(ctx, r)
}

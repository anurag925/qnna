package middlewares

import (
	"context"

	"github.com/anurag925/qnna/internal/loggers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// RequestID generates a unique request ID for each request
func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Generate a unique request ID
		requestID := uuid.NewString()

		// Set the request ID in the request context
		ctx := context.WithValue(c.Request().Context(), loggers.RequestIDCtxKey, requestID)

		request := c.Request().Clone(ctx)
		c.SetRequest(request)
		return next(c)
	}
}

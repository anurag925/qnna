package server

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/shibaone/agreements-service/pkg/api"
	"github.com/shibaone/agreements-service/pkg/errs"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	ctx := c.Request().Context()
	slog.ErrorContext(ctx, "error occurred in api handling at custom handler", "error", err)

	var code int
	var message any
	var internal error
	httpErr, ok := err.(*echo.HTTPError)
	if ok {
		internal = httpErr.Internal
	} else {
		internal = err
	}

	if errors.Is(err, sql.ErrNoRows) {
		code = http.StatusNotFound
		message = "record not found"
	} else if errors.Is(err, errs.ErrStructValidation) || errors.Is(err, errs.ErrBadRequest) {
		code = http.StatusBadRequest
		message = "invalid data in given request"
	} else if errors.Is(err, errs.ErrNotFound) {
		code = http.StatusNotFound
		message = errs.ErrNotFound.Error()
	} else if errors.Is(err, errs.ErrInvalidInput) || errors.Is(err, errs.ErrMalformedData) {
		code = http.StatusBadRequest
		message = errs.ErrInvalidInput.Error()
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		slog.ErrorContext(ctx, "internal server error", "error", err)
		if hub := sentryecho.GetHubFromContext(c); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				scope.AddBreadcrumb(&sentry.Breadcrumb{Type: "error", Category: "internal server", Message: err.Error(), Level: sentry.LevelError}, 1)
				hub.CaptureException(err)
			})
		}
		code = http.StatusInternalServerError
		message = "something went wrong please try again later"
	}

	// Return the error response in JSON format
	errorResponse := api.Response{
		Success: false,
		Error: &echo.HTTPError{
			Code:     code,
			Message:  message,
			Internal: internal,
		},
	}

	slog.InfoContext(ctx, "error response sent via common handler", "response", errorResponse)

	// Check if the response has already been committed
	// If not, commit the response with the error details
	if !c.Response().Committed {
		c.JSON(code, errorResponse)
	}
}

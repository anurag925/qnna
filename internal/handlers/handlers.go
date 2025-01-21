package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/anurag925/qnna/pkg/api"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Health(c echo.Context) error
}

type handler struct {
}

var _ Handler = (*handler)(nil)

func NewHandler() (Handler, error) {
	return &handler{}, nil
}

func (h *handler) bindAndValidate(c echo.Context, obj any) error {
	ctx := h.ctx(c)
	slog.InfoContext(ctx, "binding request...")
	if err := c.Bind(obj); err != nil {
		slog.ErrorContext(ctx, "failed to bind request", "error", err)
		return err
	}
	slog.InfoContext(ctx, "validating request...")
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		slog.ErrorContext(ctx, "failed to validate request", "error", err)
		return err
	}
	return nil
}

func (h *handler) ctx(c echo.Context) context.Context {
	return c.Request().Context()
}

func (h *handler) success(c echo.Context, status int, data any) error {
	var resp = &api.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(status, resp)
}

func (h *handler) Health(c echo.Context) error {
	var resp = &api.Response{
		Success: true,
		Data:    "ok",
	}

	return c.JSON(http.StatusOK, resp)
}

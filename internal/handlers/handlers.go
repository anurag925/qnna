package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/anurag925/qnna/internal/repositories"
	"github.com/anurag925/qnna/pkg/api"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Health(c echo.Context) error
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type handler struct {
	userRepo *repositories.UserRepository
}

var _ Handler = (*handler)(nil)

func NewHandler(userRepo *repositories.UserRepository) (Handler, error) {
	return &handler{userRepo: userRepo}, nil
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

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func (h *handler) render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func (h *handler) isGet(c echo.Context) bool {
	return c.Request().Method == http.MethodGet
}

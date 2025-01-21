package api

import "github.com/labstack/echo/v4"

type Response struct {
	Success bool            `json:"success"`
	Error   *echo.HTTPError `json:"error,omitempty"`
	Data    any             `json:"data,omitempty"`
}

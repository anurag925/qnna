package server

import (
	"encoding/json"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/anurag925/qnna/configs"
	"github.com/labstack/echo/v4"
)

func (s *Rest) routes() {
	root := s.router.Group("")
	{
		root.GET("/health", s.handles.Health)

	}
	apiV1 := root.Group("/api/v1")
	{
		apiV1.POST("/signup", s.handles.SignUp)
	}
	s.printRoutes()
}

func (s *Rest) printRoutes() error {
	if !configs.Get().Debug {
		return nil
	}
	slog.Debug("printing routes ...")
	// print routes
	routes := make([]*echo.Route, 0)

	for _, route := range s.router.Routes() {
		if route.Method == echo.RouteNotFound {
			continue
		}
		routes = append(routes, route)
	}
	slices.SortFunc(routes, func(a, b *echo.Route) int {
		val := strings.Compare(a.Path, b.Path)
		if val == 0 {
			return strings.Compare(a.Method, b.Method)
		}
		return val
	})
	data, err := json.MarshalIndent(routes, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile("routes.json", data, 0644); err != nil {
		return err
	}
	return nil
}

// 1) integration with user service
// 2) setup of nats
// 3) ads redirection service setup d

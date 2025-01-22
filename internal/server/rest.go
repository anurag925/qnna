package server

import (
	"log/slog"

	"github.com/anurag925/qnna/internal/handlers"
	"github.com/anurag925/qnna/internal/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

type Rest struct {
	router  *echo.Echo
	handles handlers.Handler
}

func NewRest(handles handlers.Handler) *Rest {
	router := echo.New()

	// Logger Middleware logs HTTP requests
	router.Use(slogecho.New(slog.Default()))

	// Recover Middleware recovers from panics anywhere in the chain
	router.Use(middleware.Recover())

	// router.Use(sentryecho.New(sentryecho.Options{
	// 	Repanic: true,
	// 	Timeout: 10 * time.Millisecond,
	// }))

	// Secure Middleware for protection against cross-site scripting (XSS) attack,
	// content type sniffing, Clickjacking, insecure connection and other code injection attacks.
	router.Use(middleware.Secure())

	// requestId Middleware generates a unique request ID for each request
	router.Use(middlewares.RequestID)

	// CORS Middleware
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccessControlAllowMethods,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowCredentials,
		},
	}))
	router.HTTPErrorHandler = customHTTPErrorHandler

	srv := &Rest{router: router, handles: handles}
	srv.routes()

	return srv
}

func (s *Rest) Run(addr string) error {

	// Start server
	go func() {
		if err := s.router.Start(addr); err != nil {
			slog.Error("unable to start server", "error", err)
		}
	}()

	return nil
}

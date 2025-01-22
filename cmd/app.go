package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/anurag925/qnna/configs"
	"github.com/anurag925/qnna/internal/handlers"
	"github.com/anurag925/qnna/internal/loggers"
	"github.com/anurag925/qnna/internal/repositories"
	"github.com/anurag925/qnna/internal/server"
	"github.com/anurag925/qnna/internal/utils"

	"github.com/urfave/cli/v2"
)

type App struct {
	app *cli.App
}

func New() *App {
	a := &App{
		app: cli.NewApp(),
	}
	cfg := &cli.StringFlag{}
	a.app.Usage = "ads-users-service"
	a.app.Flags = []cli.Flag{cfg}

	a.commands()

	return a
}

func (a *App) Run(args []string) error {
	if args == nil {
		args = os.Args
	}
	return a.app.Run(args)
}

func (a *App) Start(c *cli.Context) error {
	cfg := configs.Get()
	ctx := context.Background()
	loggers.Init(cfg.Debug)

	// connect to database
	db := utils.ConnectSQLite(ctx, cfg.Debug)

	// initialize repositories
	usersRepo := repositories.NewUserRepository(db)

	// initialize services

	// initialize handlers
	// rest handler
	handler, err := handlers.NewHandler(usersRepo)
	if err != nil {
		return err
	}

	// run rest server
	if err = server.NewRest(handler).Run(cfg.HTTPListenHostPort); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// wait for signal
	<-quit
	slog.Info("shutting down")

	return nil
}

func (a *App) commands() {
	a.app.Commands = []*cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "start server",
			Action: func(c *cli.Context) error {
				return a.Start(c)
			},
		},
	}
}

package main

import (
	"log/slog"
	"os"
)

// @title					 ads-users-service API
// @version					1.0
// @description				This services basically tells the user about what are the dids that are linked with the users main did.
// @contact.name				API Support
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	if err := New().Run(os.Args); err != nil {
		slog.Error("error in starting the app", "error", err)
		os.Exit(1)
	}
}

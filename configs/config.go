package configs

import (
	"log/slog"
	"sync"

	v11env "github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Debug     bool   `env:"DEBUG" envDefault:"false"`
	SentryDNS string `env:"SENTRY_DNS" envDefault:"https://examplePublicKey@o0.ingest.sentry.io/0"`

	HTTPListenHostPort string `env:"HTTP_LISTEN_HOST_PORT" envDefault:"127.0.0.1:3002"`
}

var instance Config
var instanceOnce sync.Once

// Get returns the config instance
func Get() *Config {
	instanceOnce.Do(func() {
		if err := v11env.Parse(&instance); err != nil {
			slog.Error("unable to load secrets from .env", "error", err)
			panic(err)
		}
		if instance.Debug {
			slog.Debug("printing the configs", "config", instance)
		}
	})
	return &instance
}

// this will only contain keys that doesn't have a default value
// and will be used only for tests to load those values

func LoadConfigForTest() {
	instanceOnce.Do(func() {
		instance = Config{
			Debug:              false,
			SentryDNS:          "",
			HTTPListenHostPort: "127.0.0.1:2090",
		}
		slog.Info("printing the configs", "config", instance)
	})
}

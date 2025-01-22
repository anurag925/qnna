package loggers

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func Init(debug bool) {
	logLevel := slog.LevelInfo
	if debug {
		logLevel = slog.LevelDebug
	}
	slogHandler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      logLevel,
		TimeFormat: time.DateTime,
		AddSource:  true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if err, ok := a.Value.Any().(error); ok {
				aErr := tint.Err(err)
				aErr.Key = a.Key
				return aErr
			}
			return a
		},
	})
	slog.SetDefault(slog.New(NewCustomHandler(slogHandler)))
}

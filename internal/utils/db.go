package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/pandoratoolbox/bun/extra/bunslog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func ConnectSQLite(ctx context.Context, ping bool) *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(fmt.Sprintf("error opening database %s", err))
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	hook := bunslog.NewQueryHook(
		bunslog.WithQueryLogLevel(slog.LevelInfo),
		bunslog.WithSlowQueryLogLevel(slog.LevelWarn),
		bunslog.WithErrorQueryLogLevel(slog.LevelError),
		bunslog.WithSlowQueryThreshold(3*time.Second),
	)
	db.AddQueryHook(hook)
	if ping {
		if err := db.Ping(); err != nil {
			panic(fmt.Sprintf("error pinging to database %s", err))
		}
	}
	return db
}

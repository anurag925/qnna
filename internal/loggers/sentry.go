package loggers

// import (
// 	"fmt"

// 	"github.com/getsentry/sentry-go"
// )

// func InitSentry(dns string, debug bool, env string) {
// 	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
// 	if err := sentry.Init(sentry.ClientOptions{
// 		Dsn: dns,
// 		// Set TracesSampleRate to 1.0 to capture 100%
// 		// of transactions for tracing.
// 		// We recommend adjusting this value in production,
// 		TracesSampleRate: 1.0,
// 		Debug:            debug,
// 		Environment:      env,
// 		ServerName:       "ads-users-service",
// 	}); err != nil {
// 		fmt.Printf("Sentry initialization failed: %v\n", err)
// 	}

// }

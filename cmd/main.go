package main

import (
	"os"
	"os/signal"

	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/middlewares/cors"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/middlewares/metrics"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/middlewares/panics"

	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/config"
	transport "github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/router"
)

func main() {
	cfg := config.LoadConfig()

	router := transport.HandlersCreate(cfg, metrics.Middleware, cors.Middleware, panics.Middleware)
	router.WithMetrics()
	router.Run()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	// We received an interrupt signal, shut down.
	router.Shutdown()
}

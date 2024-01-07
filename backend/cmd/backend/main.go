package main

import (
	"context"
	"fmt"
	"github.com/imthaghost/musik/backend/internal/music/soundcloud"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"

	"github.com/imthaghost/musik/backend/config"
	"github.com/imthaghost/musik/backend/internal/api"
	"github.com/imthaghost/musik/backend/internal/errors/sentry"
	"github.com/imthaghost/musik/backend/internal/logger/zap"
)

var startup time.Time

func main() {
	startup = time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	configService := config.New{}
	configService.Load()
	cfg := configService.Get()

	// init logger service
	logService := zap.NewService()

	// init error monitor service
	errorMonitorService := sentry.NewService(&cfg)

	// init music service
	musicService := soundcloud.NewService()

	// init server and inject dependencies
	server := api.NewServer(
		cfg,
		logService,
		errorMonitorService,
		musicService,
	)

	// start message
	startMsg := fmt.Sprintf("starting http server took %v", time.Since(startup))
	logService.Msg(startMsg)

	// start server
	go func() {
		// listen and serve
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logService.Msg("shutting down server")
			errorMonitorService.Report(err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Close(ctx); err != nil {
		errorMonitorService.Report(err)
	}
}

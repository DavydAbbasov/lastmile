package main

import (
	"context"
	"lastmile/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	logger := initLogger(cfg.App.Env, cfg.App.Name)

	log.Info().
		Msg("starting service")

		// application, err := app.New(cfg, logger)
		// if err != nil {
		// 	logger.Fatal().Err(err).Msg("failed to build application")
		// }

	//graceful shutdown
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// go func() {
	// 	if err := application.Run(); err != nil {
	// 		logger.Error().Err(err).Msg("application stoped with error")
	// 	}
	// }()

	<-ctx.Done()
	logger.Info().Msg("shuting down...")

	// sutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// if err := application.Shutdown(shutdownCtx); err != nil {
	// 	logger.Error().Err(err).Msg("graceful shutdown failed")
	// } else {
	// 	logger.Info().Msg("service stopped gracefully")
	// }

}
func initLogger(env, serviceName string) zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	//base zerolog
	var base zerolog.Logger

	//logger for dev
	if env == "local" || env == "dev" {
		base = zerolog.New(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: "15:04:05",
			},
		)
	} else {
		//for prod
		base = zerolog.New(os.Stdout)
	}

	//build custum logger
	logger := base.
		With().
		Timestamp().
		Str("service", serviceName).
		Str("env", env).
		Logger()

	log.Logger = logger

	return logger
}

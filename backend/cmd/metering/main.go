package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/wqrxssq/project_seminar/backend/internal/httpserver"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	addr := httpserver.ListenAddr()
	logger.Info("starting",
		"service", "metering",
		"listen_addr", addr,
		"database_configured", os.Getenv("DATABASE_URL") != "",
		"nats_configured", os.Getenv("NATS_URL") != "",
		"s3_configured", os.Getenv("S3_ENDPOINT") != "",
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := httpserver.Run(ctx, logger, "metering", addr); err != nil {
		logger.Error("server exited", "err", err, "service", "metering")
		os.Exit(1)
	}
}

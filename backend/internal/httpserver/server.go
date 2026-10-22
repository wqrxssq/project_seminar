// Package httpserver is the shared draft process for every service.
// The only route is GET /ping. Domain handlers are not registered here.
package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const defaultAddr = ":8080"

// ListenAddr returns LISTEN_ADDR, or :8080 when that variable is empty.
func ListenAddr() string {
	if addr := os.Getenv("LISTEN_ADDR"); addr != "" {
		return addr
	}
	return defaultAddr
}

type pingBody struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

// NewMux returns the draft router. Only GET /ping is served.
func NewMux(service string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		body, err := json.Marshal(pingBody{Service: service, Status: "ok"})
		if err != nil {
			http.Error(w, "encode", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
		_, _ = w.Write([]byte("\n"))
	})
	return mux
}

// Run serves GET /ping until ctx is cancelled, then shuts the listener down.
func Run(ctx context.Context, logger *slog.Logger, service, addr string) error {
	if logger == nil {
		logger = slog.Default()
	}
	srv := &http.Server{
		Addr:              addr,
		Handler:           NewMux(service),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("http listening", "service", service, "addr", addr)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		logger.Info("http shutting down", "service", service)
		return srv.Shutdown(shutdownCtx)
	}
}

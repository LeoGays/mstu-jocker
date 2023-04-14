package httpx

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
)

func NewServer(cfg *Config, handler http.Handler) *http.Server {
	return &http.Server{
		Handler:           handler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		Addr:              cfg.Addr(),
	}
}

// StartServer starts a server, closes on context cancel.
func StartServer(ctx context.Context, server *http.Server) error {
	log.Printf("starting server at: %s", server.Addr)

	server.BaseContext = func(net.Listener) context.Context {
		return ctx
	}

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			log.Printf("server has closed")

			return nil
		}

		return err
	case <-ctx.Done():
		log.Printf("closing server due to context cancellation")

		if err := server.Close(); err != nil {
			return fmt.Errorf("failed to close http server: %w", err)
		}

		return nil
	}
}

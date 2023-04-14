package server

import (
	"github.com/go-chi/chi/v5"
	"jocer/config"
	"jocer/internal/server/generated"
	"jocer/internal/usecase"
)

var _ generated.ServerInterface = (*Server)(nil)

type (
	Server struct {
		cfg     *config.Config
		useCase usecase.UseCase
	}
)

func New(cfg *config.Config, useCase usecase.UseCase) *Server {
	return &Server{
		cfg:     cfg,
		useCase: useCase,
	}
}

func (s *Server) NewServerOptions() generated.ChiServerOptions {
	router := chi.NewRouter()

	return generated.ChiServerOptions{
		BaseRouter: router,
	}
}

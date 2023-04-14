package main

import (
	"context"
	"fmt"
	"jocer/config"
	"jocer/internal/server"
	"jocer/internal/server/generated"
	"jocer/internal/storage"
	"jocer/internal/usecase"
	"jocer/pkg/db/entx"
	"jocer/pkg/httpx"
	"jocer/pkg/logs"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadFromViper()

	dbStorage, storageClose, err := databaseStorage(ctx, cfg)
	defer storageClose()
	if err != nil {
		log.Fatal("unable to init db storage: %w", err)
	}

	useCase := usecase.NewUseCaseImpl(dbStorage)
	srv := server.New(cfg, useCase)
	options := srv.NewServerOptions()

	if err := httpx.StartServer(ctx, httpx.NewServer(cfg.HTTP, generated.HandlerWithOptions(srv, options))); err != nil {
		log.Fatal("unable to start server %w", err)
	}
}

func databaseStorage(ctx context.Context, cfg *config.Config) (storage.Storage, func(), error) {
	dbDriver, err := entx.Driver(cfg.DB)
	if err != nil {
		return nil, func() {}, fmt.Errorf("failed connect to DB: %w", err)
	}
	dbClose := func() {
		if err := dbDriver.Close(); err != nil {
			logs.FromContext(ctx).Err(err).Msg("failed to close ent client")
		}
	}
	dbClient, err := storage.NewDBClient(ctx, dbDriver, cfg.DB.Debug)
	if err != nil {
		return nil, dbClose, fmt.Errorf("failed to init ent client: %w", err)
	}

	return storage.New(dbClient), dbClose, nil
}

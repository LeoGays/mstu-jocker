package storage

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"jocer/internal/storage/ent"
	"jocer/pkg/logs"
)

type (
	Storage interface {
		Jock
	}

	StorageImpl struct {
		Client *ent.Client
	}
)

func New(client *ent.Client) Storage {
	return &StorageImpl{
		Client: client,
	}
}

func NewDBClient(ctx context.Context, drv *sql.Driver, debug bool) (*ent.Client, error) {
	logger := logs.FromContext(ctx)

	options := []ent.Option{ent.Driver(drv), ent.Log(logger.Print)}
	if debug {
		options = append(options, ent.Debug())
	}
	client := ent.NewClient(options...)
	if debug {
		client = client.Debug()
	}
	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return client, nil
}

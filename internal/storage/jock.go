package storage

import (
	"context"
	"jocer/internal/model"
	"jocer/internal/storage/mapper"
)

var _ Jock = (*StorageImpl)(nil)

type (
	Jock interface {
		GetJocks(ctx context.Context) ([]*model.Jock, error)
		Create(ctx context.Context, jock *model.Jock) (*model.Jock, error)
	}
)

func (s StorageImpl) GetJocks(ctx context.Context) ([]*model.Jock, error) {
	jocks, err := s.Client.Jock.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.CreateJockList(jocks), nil
}

func (s StorageImpl) Create(ctx context.Context, jock *model.Jock) (*model.Jock, error) {
	jockEnt, err := s.Client.Jock.Create().SetName(jock.Name).SetContent(jock.Content).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Jock{
		ID:      jockEnt.ID,
		Name:    jockEnt.Name,
		Content: jockEnt.Content,
	}, nil
}

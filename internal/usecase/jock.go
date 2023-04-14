package usecase

import (
	"context"
	"fmt"
	"jocer/internal/model"
)

var _ Jock = (*UseCaseImpl)(nil)

type Jock interface {
	GetJocks(ctx context.Context) ([]*model.Jock, error)
	CreateJock(ctx context.Context, jock *model.Jock) (*model.Jock, error)
}

func (u UseCaseImpl) GetJocks(ctx context.Context) ([]*model.Jock, error) {
	jocks, err := u.Storage.GetJocks(ctx)
	if err != nil {
		return nil, fmt.Errorf("Storage GetJocks is %v", err)
	}
	return jocks, nil
}

func (u UseCaseImpl) CreateJock(ctx context.Context, jock *model.Jock) (*model.Jock, error) {
	return u.Storage.Create(ctx, jock)
}

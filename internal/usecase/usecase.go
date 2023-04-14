package usecase

import "jocer/internal/storage"

type (
	UseCase interface {
		Jock
	}

	UseCaseImpl struct {
		Storage storage.Storage
	}
)

func NewUseCaseImpl(storage storage.Storage) *UseCaseImpl {
	return &UseCaseImpl{
		Storage: storage,
	}
}

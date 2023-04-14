package model

import "github.com/google/uuid"

type (
	Jock struct {
		ID      uuid.UUID
		Name    string
		Content string
	}
)

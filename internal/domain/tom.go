package domain

import (
	"context"

	"github.com/google/uuid"
)

type TomRepository interface {
	AddTom(context.Context, RegisterTomRequest) (uuid.UUID, error)
	UpdateTom(context.Context, UpdateTomRequest) (*Tom, error)
}

type Tom struct {
	ID    uuid.UUID
	Name  string
	Ready bool
}

type RegisterTomRequest struct {
	Name string
}

type UpdateTomRequest struct {
	ID    uuid.UUID
	Name  *string
	Ready *bool
}

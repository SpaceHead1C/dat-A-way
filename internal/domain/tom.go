package domain

import (
	"context"

	"github.com/google/uuid"
)

type TomRepository interface {
	AddTom(context.Context, RegisterTomRequest) (uuid.UUID, error)
}

type RegisterTomRequest struct {
	Name string
}

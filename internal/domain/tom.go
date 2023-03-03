package domain

import (
	"context"
	"github.com/google/uuid"
)

type TomRepository interface {
	AddTom(context.Context) (uuid.UUID, error)
}

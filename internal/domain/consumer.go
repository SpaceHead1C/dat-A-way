package domain

import (
	"context"

	"github.com/google/uuid"
)

type ConsumerRepository interface {
	AddConsumer(context.Context, AddConsumerRequest) (*Consumer, error)
	UpdateConsumer(context.Context, UpdConsumerRequest) (*Consumer, error)
}

type Consumer struct {
	ID          uuid.UUID
	Queue       string
	Name        string
	Description string
}

type AddConsumerRequest struct {
	Name        string
	Description string
}

type UpdConsumerRequest struct {
	ID          uuid.UUID
	Name        *string
	Description *string
}

package domain

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type ConsumerRepository interface {
	AddConsumer(context.Context, AddConsumerRequest) (*Consumer, error)
	UpdateConsumer(context.Context, UpdConsumerRequest) (*Consumer, error)
	GetConsumer(context.Context, uuid.UUID) (*Consumer, error)
}

type ConsumerBroker interface {
	DeclareConsumerQueue(context.Context, Consumer) error
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

func ConsumerQueue(id uuid.UUID) string {
	return fmt.Sprintf("o%s", strings.ReplaceAll(id.String(), "-", ""))
}

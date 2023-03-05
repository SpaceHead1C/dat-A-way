package domain

import (
	"context"
	"github.com/google/uuid"
)

type SubscriptionRepository interface {
	AddSubscription(context.Context, AddSubscriptionRequest) (*Subscription, error)
}

type Subscription struct {
	ConsumerID uuid.UUID
	TomID      uuid.UUID
	PropertyID uuid.UUID
}

type AddSubscriptionRequest struct {
	ConsumerID uuid.UUID
	TomID      uuid.UUID
	PropertyID uuid.UUID
}

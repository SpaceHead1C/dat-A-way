package domain

import (
	"context"

	"github.com/google/uuid"
)

type SubscriptionRepository interface {
	AddSubscription(context.Context, AddSubscriptionRequest) (*Subscription, error)
	DeleteSubscription(context.Context, DeleteSubscriptionRequest) (*Subscription, error)
	GetSubscriberIDs(context.Context, SubscriberIDsRequest) ([]uuid.UUID, error)
	SubscriberExists(context.Context, SubscriberExistanceRequest) (bool, error)
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

type DeleteSubscriptionRequest struct {
	ConsumerID uuid.UUID
	TomID      uuid.UUID
	PropertyID uuid.UUID
}

type SubscriberIDsRequest struct {
	TomID      uuid.UUID
	PropertyID *uuid.UUID
}

type SubscriberExistanceRequest struct {
	ConsumerID uuid.UUID
	TomID      uuid.UUID
}

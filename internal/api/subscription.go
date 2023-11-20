package api

import (
	"context"
	. "dataway/internal/domain"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const defaultSubscriptionManagerTimeout = time.Second

type SubscriptionManager struct {
	SubscriptionConfig
}

type SubscriptionConfig struct {
	Repository SubscriptionRepository
	Timeout    time.Duration
}

func NewSubscriptionManager(c SubscriptionConfig) (*SubscriptionManager, error) {
	if c.Repository == nil {
		return nil, fmt.Errorf("subscription repository can't be nil")
	}
	if c.Timeout == 0 {
		c.Timeout = defaultSubscriptionManagerTimeout
	}
	return &SubscriptionManager{c}, nil
}

func (sm *SubscriptionManager) Add(ctx context.Context, req AddSubscriptionRequest) (*Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, sm.Timeout)
	defer cancel()
	return sm.Repository.AddSubscription(ctx, req)
}

func (sm *SubscriptionManager) Delete(ctx context.Context, req DeleteSubscriptionRequest) (*Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, sm.Timeout)
	defer cancel()
	return sm.Repository.DeleteSubscription(ctx, req)
}

func (sm *SubscriptionManager) SubscriberIDs(ctx context.Context, req SubscriberIDsRequest) ([]uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, sm.Timeout)
	defer cancel()
	return sm.Repository.GetSubscriberIDs(ctx, req)
}

func (sm *SubscriptionManager) SubscriberExists(ctx context.Context, req SubscriberExistanceRequest) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, sm.Timeout)
	defer cancel()
	return sm.Repository.SubscriberExists(ctx, req)
}

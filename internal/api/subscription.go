package api

import (
	"context"
	. "dataway/internal/domain"
	"fmt"
	"time"
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

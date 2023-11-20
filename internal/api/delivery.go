package api

import (
	"context"
	"dataway/internal/domain"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const defaultDeliveryManagerTimeout = time.Second

type DeliveryManager struct {
	DeliveryConfig
}

type DeliveryConfig struct {
	Broker  domain.DeliveryBroker
	Timeout time.Duration
}

func NewDeliveryManager(c DeliveryConfig) (*DeliveryManager, error) {
	if c.Broker == nil {
		return nil, fmt.Errorf("delivery broker can not be nil")
	}
	if c.Timeout == 0 {
		c.Timeout = defaultDeliveryManagerTimeout
	}
	return &DeliveryManager{c}, nil
}

func (dm *DeliveryManager) Send(ctx context.Context, consumerID uuid.UUID, m domain.Message) error {
	ctx, cancel := context.WithTimeout(ctx, dm.Timeout)
	defer cancel()
	return dm.Broker.SendDelivery(ctx, consumerID, m)
}

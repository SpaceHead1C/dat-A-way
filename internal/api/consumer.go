package api

import (
	"context"
	. "dataway/internal/domain"
	"fmt"
	"time"
)

const defaultConsumerManagerTimeout = time.Second

type ConsumerManager struct {
	ConsumerConfig
}

type ConsumerConfig struct {
	Repository ConsumerRepository
	Timeout    time.Duration
}

func NewConsumerManager(c ConsumerConfig) (*ConsumerManager, error) {
	if c.Repository == nil {
		return nil, fmt.Errorf("consumer repository can't be nil")
	}
	if c.Timeout == 0 {
		c.Timeout = defaultConsumerManagerTimeout
	}
	return &ConsumerManager{c}, nil
}

func (cm *ConsumerManager) Add(ctx context.Context, req AddConsumerRequest) (*Consumer, error) {
	ctx, cancel := context.WithTimeout(ctx, cm.Timeout)
	defer cancel()
	return cm.Repository.AddConsumer(ctx, req)
}

func (cm *ConsumerManager) Update(ctx context.Context, req UpdConsumerRequest) (*Consumer, error) {
	ctx, cancel := context.WithTimeout(ctx, cm.Timeout)
	defer cancel()
	return cm.Repository.UpdateConsumer(ctx, req)
}

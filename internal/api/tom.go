package api

import (
	"context"
	"fmt"
	"time"

	. "dataway/internal/domain"

	"github.com/google/uuid"
)

const defaultTomManagerTimeout = time.Second

type TomManager struct {
	TomConfig
}

type TomConfig struct {
	Repository TomRepository
	Timeout    time.Duration
}

func NewTomManager(c TomConfig) (*TomManager, error) {
	if c.Repository == nil {
		return nil, fmt.Errorf("tom repository can not be nil")
	}
	if c.Timeout == 0 {
		c.Timeout = defaultTomManagerTimeout
	}
	return &TomManager{c}, nil
}

func (tm *TomManager) Add(ctx context.Context, req RegisterTomRequest) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, tm.Timeout)
	defer cancel()
	return tm.Repository.AddTom(ctx, req)
}

func (tm *TomManager) Update(ctx context.Context, req UpdateTomRequest) (*Tom, error) {
	ctx, cancel := context.WithTimeout(ctx, tm.Timeout)
	defer cancel()
	return tm.Repository.UpdateTom(ctx, req)
}

func (tm *TomManager) Get(ctx context.Context, id uuid.UUID) (*Tom, error) {
	ctx, cancel := context.WithTimeout(ctx, tm.Timeout)
	defer cancel()
	return tm.Repository.GetTom(ctx, id)
}

func (tm *TomManager) QueueDeclare(ctx context.Context, tom Tom) error {
	ctx, cancel := context.WithTimeout(ctx, tm.Timeout)
	defer cancel()
	return tm.Broker.DeclareTomQueue(ctx, tom)
}

func (tm *TomManager) SetAsReady(ctx context.Context, id uuid.UUID) error {
	ready := true
	ctx, cancel := context.WithTimeout(ctx, tm.Timeout)
	defer cancel()
	_, err := tm.Repository.UpdateTom(ctx, UpdateTomRequest{
		ID:    id,
		Ready: &ready,
	})
	return err
}

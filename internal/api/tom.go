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

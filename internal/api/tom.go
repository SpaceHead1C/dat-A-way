package api

import (
	"context"
	. "dataway/internal/domain"
	"fmt"
	"github.com/google/uuid"
	"time"
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
		return nil, fmt.Errorf("tom repository can't be nil")
	}
	if c.Timeout == 0 {
		c.Timeout = defaultTomManagerTimeout
	}
	return &TomManager{c}, nil
}

func (tm *TomManager) Add(ctx context.Context) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, tm.Timeout)
	defer cancel()
	return tm.Repository.AddTom(ctx)
}

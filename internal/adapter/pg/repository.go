package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const getUUIDAttemptsThreshold = 10

type Config struct {
	ConnectConfig *pgxpool.Config
	Logger        *zap.SugaredLogger
}

type Repository struct {
	*pgxpool.Pool
	l *zap.SugaredLogger
}

func NewRepository(ctx context.Context, c Config) (*Repository, error) {
	if c.ConnectConfig == nil {
		return nil, fmt.Errorf("connect config is nil")
	}
	if c.Logger == nil {
		c.Logger = zap.L().Sugar()
	}
	pool, err := pgxpool.NewWithConfig(ctx, c.ConnectConfig)
	if err != nil {
		return nil, err
	}
	return &Repository{pool, c.Logger}, nil
}

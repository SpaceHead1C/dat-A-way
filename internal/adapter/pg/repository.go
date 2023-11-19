package pg

import (
	"context"
	"dataway/pkg/log"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	getUUIDAttemptsThreshold = 10

	tableTomColumnName = "name"
)

type Config struct {
	ConnectConfig *pgxpool.Config
	Logger        *log.Logger
}

type Repository struct {
	*pgxpool.Pool
	l *log.Logger
}

func NewRepository(ctx context.Context, c Config) (*Repository, error) {
	if c.ConnectConfig == nil {
		return nil, fmt.Errorf("connect config is nil")
	}
	if c.Logger == nil {
		c.Logger = log.GlobalLogger()
	}
	pool, err := pgxpool.NewWithConfig(ctx, c.ConnectConfig)
	if err != nil {
		return nil, err
	}
	return &Repository{pool, c.Logger}, nil
}

package pg

import (
	"context"
	. "dataway/internal/domain"
	"dataway/pkg/db/pg"
	"fmt"
	"github.com/google/uuid"
)

func (r *Repository) AddConsumer(ctx context.Context, req AddConsumerRequest) (*Consumer, error) {
	args := []any{
		req.Name,
		req.Description,
	}
	query := `SELECT * FROM new_consumer($1, $2);`
	for attempts := 0; attempts < getUUIDAttemptsThreshold; attempts++ {
		var out Consumer
		if err := r.QueryRow(ctx, query, args...).Scan(&out.ID, &out.Queue, &out.Name, &out.Description); err != nil {
			if pg.IsNotUniqueError(err) {
				continue
			}
			return nil, fmt.Errorf("database error: %w, %s", err, query)
		}
		return &out, nil
	}
	return nil, errCanNotGetUniqueID
}

func (r *Repository) UpdateConsumer(ctx context.Context, req UpdConsumerRequest) (*Consumer, error) {
	var out Consumer
	args := make([]any, 3)
	args[0] = req.ID
	if req.Name != nil {
		args[1] = *req.Name
	}
	if req.Description != nil {
		args[2] = *req.Description
	}
	query := `SELECT * FROM update_consumer($1, $2, $3);`
	if err := r.QueryRow(ctx, query, args...).Scan(&out.ID, &out.Queue, &out.Name, &out.Description); err != nil {
		if pg.IsNoRowsError(err) {
			return nil, ErrConsumerNotFound
		}
		return nil, fmt.Errorf("database error: %w, %s", err, query)
	}
	return &out, nil
}

func (r *Repository) GetConsumer(ctx context.Context, id uuid.UUID) (*Consumer, error) {
	query := `SELECT * FROM get_consumer($1);`
	var out Consumer
	if err := r.QueryRow(ctx, query, id).Scan(&out.ID, &out.Queue, &out.Name, &out.Description); err != nil {
		if pg.IsNoRowsError(err) {
			return nil, ErrConsumerNotFound
		}
		return nil, fmt.Errorf("database error: %w, %s", err, query)
	}
	return &out, nil
}

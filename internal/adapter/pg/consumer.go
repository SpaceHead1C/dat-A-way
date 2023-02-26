package pg

import (
	"context"
	. "dataway/internal/domain"
	"dataway/pkg/db/pg"
	"fmt"
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

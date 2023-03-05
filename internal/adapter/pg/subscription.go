package pg

import (
	"context"
	. "dataway/internal/domain"
	"dataway/pkg/db/pg"
	"fmt"
)

func (r *Repository) AddSubscription(ctx context.Context, req AddSubscriptionRequest) (*Subscription, error) {
	var out Subscription
	args := []any{
		pg.NullUUID(req.ConsumerID),
		pg.NullUUID(req.TomID),
		pg.NullUUID(req.PropertyID),
	}
	query := `SELECT * FROM set_subscription($1, $2, $3);`
	if err := r.QueryRow(ctx, query, args...).Scan(&out.ConsumerID, &out.TomID, &out.PropertyID); err != nil {
		return nil, fmt.Errorf("database error: %w, %s", err, query)
	}
	return &out, nil
}

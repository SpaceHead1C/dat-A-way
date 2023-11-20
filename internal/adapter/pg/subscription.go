package pg

import (
	"context"
	"fmt"

	. "dataway/internal/domain"
	"dataway/pkg/db/pg"

	"github.com/google/uuid"
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

func (r *Repository) DeleteSubscription(ctx context.Context, req DeleteSubscriptionRequest) (*Subscription, error) {
	var out Subscription
	args := []any{
		pg.NullUUID(req.ConsumerID),
		pg.NullUUID(req.TomID),
		pg.NullUUID(req.PropertyID),
	}
	query := `SELECT * FROM delete_subscription($1, $2, $3);`
	if err := r.QueryRow(ctx, query, args...).Scan(&out.ConsumerID, &out.TomID, &out.PropertyID); err != nil {
		if pg.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("database error: %w, %s", err, query)
	}
	return &out, nil
}

func (r *Repository) GetSubscriberIDs(ctx context.Context, req SubscriberIDsRequest) ([]uuid.UUID, error) {
	var out []uuid.UUID
	args := make([]any, 2)
	args[0] = req.TomID
	if req.PropertyID != nil {
		args[1] = req.PropertyID
	}
	query := `SELECT get_subscriber_ids($1, $2);`
	rows, err := r.Query(ctx, query, args...)
	if err != nil {
		if pg.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("database error: %w, %s", err, query)
	}
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("database scan error: %w, %s", err, query)
		}
		out = append(out, id)
	}
	return out, nil
}

func (r *Repository) SubscriberExists(ctx context.Context, req SubscriberExistanceRequest) (bool, error) {
	args := []any{
		req.ConsumerID,
		req.TomID,
	}
	query := `SELECT subscriber_exists($1, $2);`
	var out bool
	if err := r.QueryRow(ctx, query, args...).Scan(&out); err != nil {
		return out, fmt.Errorf("database scan error: %w, %s", err, query)
	}
	return out, nil
}

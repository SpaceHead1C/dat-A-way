package pg

import (
	"context"
	"fmt"

	"dataway/internal/domain"
	"dataway/pkg/db/pg"

	"github.com/google/uuid"
)

func (r *Repository) AddTom(ctx context.Context, req domain.RegisterTomRequest) (uuid.UUID, error) {
	var out uuid.UUID
	args := []any{
		pg.NullString(req.Name),
	}
	query := `SELECT new_tom($1);`
	for attempts := 0; attempts < getUUIDAttemptsThreshold; attempts++ {
		if err := r.QueryRow(ctx, query, args...).Scan(&out); err != nil {
			if isTomNameDuplicateError(err, req.Name) {
				return out, domain.ErrNameDuplicate
			}
			if pg.IsNotUniqueError(err) {
				continue
			}
			return out, fmt.Errorf("database error: %w, %s", err, query)
		}
		return out, nil
	}
	return out, errCanNotGetUniqueID
}

func (r *Repository) UpdateTom(ctx context.Context, req domain.UpdateTomRequest) (*domain.Tom, error) {
	emptyReq := true
	args := make([]any, 3)
	args[0] = req.ID
	if req.Name != nil {
		args[1] = pg.NullString(*req.Name)
		emptyReq = false
	}
	if req.Ready != nil {
		args[2] = *req.Ready
		emptyReq = false
	}
	if emptyReq {
		return r.GetTom(ctx, req.ID)
	}
	var out domain.Tom
	query := `SELECT * FROM update_tom($1, $2, $3);`
	if err := r.QueryRow(ctx, query, args...).Scan(&out.ID, &out.Name, &out.Ready); err != nil {
		if pg.IsNoRowsError(err) {
			return nil, fmt.Errorf("tom %w", domain.ErrNotFound)
		}
		if req.Name != nil && isTomNameDuplicateError(err, *req.Name) {
			return nil, domain.ErrNameDuplicate
		}
		return nil, fmt.Errorf("database error: %w, %s", err, query)
	}
	return &out, nil
}

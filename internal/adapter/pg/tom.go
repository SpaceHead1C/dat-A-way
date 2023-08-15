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

package pg

import (
	"context"
	"dataway/pkg/db/pg"
	"fmt"
	"github.com/google/uuid"
)

func (r *Repository) AddTom(ctx context.Context) (uuid.UUID, error) {
	var out uuid.UUID
	query := `SELECT new_tom();`
	for attempts := 0; attempts < getUUIDAttemptsThreshold; attempts++ {
		if err := r.QueryRow(ctx, query).Scan(&out); err != nil {
			if pg.IsNotUniqueError(err) {
				continue
			}
			return out, fmt.Errorf("database error: %w, %s", err, query)
		}
		return out, nil
	}
	return out, errCanNotGetUniqueID
}

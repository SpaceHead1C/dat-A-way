package pg

import (
	"dataway/pkg/helper"
	"github.com/google/uuid"
)

func NullUUID(v uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{UUID: v, Valid: !helper.IsZeroUUID(v)}
}

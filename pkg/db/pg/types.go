package pg

import (
	"database/sql"
	"dataway/pkg/helper"

	"github.com/google/uuid"
)

func NullUUID(v uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{UUID: v, Valid: !helper.IsZeroUUID(v)}
}

func NullString(v string) sql.NullString {
	return sql.NullString{String: v, Valid: v != ""}
}

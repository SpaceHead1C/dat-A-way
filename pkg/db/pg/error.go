package pg

import (
	"github.com/jackc/pgx/v5/pgconn"
)

func IsNotUniqueError(err error) bool {
	if err == nil {
		return false
	}
	pgErr, ok := err.(*pgconn.PgError)
	return ok && pgErr.Code == "23505"
}

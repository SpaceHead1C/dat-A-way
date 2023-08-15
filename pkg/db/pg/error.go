package pg

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func IsNotUniqueError(err error) bool {
	if err == nil {
		return false
	}
	pgErr, ok := err.(*pgconn.PgError)
	return ok && pgErr.Code == "23505"
}

func IsNoRowsError(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

func IsDuplicateKeyError(err error, key, value string) bool {
	if !IsNotUniqueError(err) {
		return false
	}
	pgErr, _ := err.(*pgconn.PgError)
	return pgErr.Detail == fmt.Sprintf("Key (%s)=(%s) already exists.", key, value)
}

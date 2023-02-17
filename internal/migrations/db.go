package migrations

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func UpMigrations(cc *pgx.ConnConfig) error {
	db, err := sql.Open("pgx", stdlib.RegisterConnConfig(cc))
	if err != nil {
		return err
	}
	defer db.Close()
	return goose.Up(db, ".")
}


package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00002, down00002)
}

func up00002(tx *sql.Tx) error {
	query := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	return execQuery(query, tx)
}

func down00002(tx *sql.Tx) error {
	return nil
}

package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00006, down00006)
}

func up00006(tx *sql.Tx) error {
	query := `-- Toms
CREATE TABLE IF NOT EXISTS toms (
	id uuid PRIMARY KEY
);`
	return execQuery(query, tx)
}

func down00006(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS toms;`
	return execQuery(query, tx)
}

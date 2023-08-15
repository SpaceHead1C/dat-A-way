package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00012, down00012)
}

func up00012(tx *sql.Tx) error {
	query := `-- Update function for tom
CREATE FUNCTION update_tom(uuid, text) RETURNS SETOF toms AS $update_tom$
	BEGIN
		RETURN QUERY
			UPDATE toms SET "name" = $2 WHERE id = $1
			RETURNING *;
	END;
$update_tom$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00012(tx *sql.Tx) error {
	query := `DROP FUNCTION IF EXISTS update_tom(uuid, text);`
	return execQuery(query, tx)
}

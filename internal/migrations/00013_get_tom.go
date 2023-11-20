package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00013, down00013)
}

func up00013(tx *sql.Tx) error {
	query := `-- Get function for tom
CREATE FUNCTION get_tom(uuid) RETURNS SETOF toms AS $get_tom$
	BEGIN
		RETURN QUERY
			SELECT *
			FROM toms
			WHERE id = $1;
	END;
$get_tom$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00013(tx *sql.Tx) error {
	query := `DROP FUNCTION get_tom(uuid);`
	return execQuery(query, tx)
}

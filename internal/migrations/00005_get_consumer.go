package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00005, down00005)
}

func up00005(tx *sql.Tx) error {
	query := `-- Get function for consumer
CREATE OR REPLACE FUNCTION get_consumer(uuid) RETURNS SETOF consumers AS $get_consumer$
	BEGIN
		RETURN QUERY
			SELECT *
			FROM consumers
			WHERE id = $1;
	END;
$get_consumer$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00005(tx *sql.Tx) error {
	query := `DROP FUNCTION IF EXISTS get_consumer(uuid);`
	return execQuery(query, tx)
}

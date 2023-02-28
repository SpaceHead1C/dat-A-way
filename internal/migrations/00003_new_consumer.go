package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00003, down00003)
}

func up00003(tx *sql.Tx) error {
	query := `-- Consumer constructor
CREATE OR REPLACE FUNCTION new_consumer(text, text) RETURNS SETOF consumers AS $new_consumer$
	DECLARE
		res uuid;
	BEGIN
		res := uuid_generate_v4();
	
		RETURN QUERY
			INSERT INTO consumers (id, queue, "name", description)
			VALUES (res,  'q' || replace(res::TEXT, '-', ''), $1, $2)
			RETURNING *;
	END;
$new_consumer$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00003(tx *sql.Tx) error {
	query := `DROP FUNCTION IF EXISTS new_consumer(text, text);`
	return execQuery(query, tx)
}

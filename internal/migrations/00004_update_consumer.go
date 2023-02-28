package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00004, down00004)
}

func up00004(tx *sql.Tx) error {
	query := `-- Update function for consumer
	CREATE OR REPLACE FUNCTION update_consumer(uuid, text, text) RETURNS SETOF consumers AS $update_consumer$
		BEGIN
			RETURN QUERY
				UPDATE consumers SET
					"name" = COALESCE($2, "name", $2),
					description = COALESCE($3, description, $3)
				WHERE id = $1
				RETURNING *;
		END;
	$update_consumer$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00004(tx *sql.Tx) error {
	query := `DROP FUNCTION IF EXISTS update_consumer(uuid, text, text);`
	return execQuery(query, tx)
}

package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00016, down00016)
}

func up00016(tx *sql.Tx) error {
	query := `
DO $$ BEGIN
	ALTER TABLE toms ADD COLUMN ready bool NOT NULL DEFAULT FALSE;

	DROP FUNCTION update_tom(uuid, text);

	CREATE FUNCTION update_tom(uuid, text, bool) RETURNS SETOF toms AS $update_tom$
		BEGIN
			RETURN QUERY UPDATE toms SET
				"name" = COALESCE($2, "name", $2),
				ready = COALESCE($3, ready, $3)
			WHERE id = $1
			RETURNING *;
		END;
	$update_tom$ LANGUAGE plpgsql;
END $$;`
	return execQuery(query, tx)
}

func down00016(tx *sql.Tx) error {
	query := `
DO $$ BEGIN
	DROP FUNCTION update_tom(uuid, text, bool);

	CREATE FUNCTION update_tom(uuid, text) RETURNS SETOF toms AS $update_tom$
		BEGIN
			RETURN QUERY
				UPDATE toms SET "name" = $2 WHERE id = $1
				RETURNING *;
		END;
	$update_tom$ LANGUAGE plpgsql;

	ALTER TABLE toms DROP COLUMN ready;
END $$;`
	return execQuery(query, tx)
}

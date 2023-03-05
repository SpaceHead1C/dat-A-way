package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00007, down00007)
}

func up00007(tx *sql.Tx) error {
	query := `-- Tom constructor
CREATE OR REPLACE FUNCTION new_tom() RETURNS uuid AS $new_tom$
	DECLARE
		res uuid;
	BEGIN
		res := uuid_generate_v4();

		INSERT INTO toms (id) VALUES (res);
	
		RETURN res;
	END;
$new_tom$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00007(tx *sql.Tx) error {
	query := `DROP FUNCTION IF EXISTS new_tom();`
	return execQuery(query, tx)
}

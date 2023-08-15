package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00011, down00011)
}

func up00011(tx *sql.Tx) error {
	query := `
DO $$ BEGIN
	ALTER TABLE toms ADD COLUMN "name" varchar(128);

	UPDATE toms SET "name" = id::text;

	ALTER TABLE toms ALTER COLUMN "name" SET NOT NULL;

	CREATE UNIQUE INDEX toms_name_idx ON toms ("name");

	DROP FUNCTION new_tom();

	CREATE FUNCTION new_tom(text) RETURNS uuid AS $new_tom$
		DECLARE
			res uuid;
		BEGIN
			res := uuid_generate_v4();

			INSERT INTO toms (id, "name") VALUES (res, $1);
		
			RETURN res;
		END;
	$new_tom$ LANGUAGE plpgsql;
END $$;`
	return execQuery(query, tx)
}

func down00011(tx *sql.Tx) error {
	query := `
DO $$ BEGIN
	DROP FUNCTION new_tom(text);
	
	CREATE FUNCTION new_tom() RETURNS uuid AS $new_tom$
		DECLARE
			res uuid;
		BEGIN
			res := uuid_generate_v4();

			INSERT INTO toms (id) VALUES (res);
		
			RETURN res;
		END;
	$new_tom$ LANGUAGE plpgsql;

	ALTER TABLE toms DROP COLUMN "name";
END $$;`
	return execQuery(query, tx)
}

package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00015, down00015)
}

func up00015(tx *sql.Tx) error {
	query := `-- Get function for subscribtion exsistance
CREATE FUNCTION subscriber_exists(uuid, uuid) RETURNS bool AS $subscriber_exists$
	BEGIN
		RETURN EXISTS (
			SELECT NULL
			FROM subscriptions
			WHERE consumer_id = $1 AND tom_id = $2
			LIMIT 1);
	END;
$subscriber_exists$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00015(tx *sql.Tx) error {
	query := `DROP FUNCTION subscriber_exists(uuid, uuid);`
	return execQuery(query, tx)
}

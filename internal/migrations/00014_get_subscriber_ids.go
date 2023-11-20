package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00014, down00014)
}

func up00014(tx *sql.Tx) error {
	query := `-- Get function for subscriber IDs
CREATE FUNCTION get_subscriber_ids(uuid, uuid) RETURNS SETOF uuid AS $get_subscriber_ids$
	BEGIN
		IF $2 IS NULL THEN
			RETURN QUERY
				SELECT DISTINCT consumer_id
				FROM subscriptions
				WHERE tom_id = $1;
		ELSE
			RETURN QUERY
				SELECT consumer_id
				FROM subscriptions
				WHERE tom_id = $1 AND property_id = $2;
		END IF;
	END;
$get_subscriber_ids$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00014(tx *sql.Tx) error {
	query := `DROP FUNCTION get_subscriber_ids(uuid, uuid);`
	return execQuery(query, tx)
}

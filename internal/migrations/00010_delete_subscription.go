package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00010, down00010)
}

func up00010(tx *sql.Tx) error {
	query := `-- Delete function for subscription
CREATE FUNCTION delete_subscription(uuid, uuid, uuid) RETURNS SETOF subscriptions AS $delete_subscription$
    BEGIN
        RETURN QUERY
            DELETE FROM subscriptions
            WHERE consumer_id = $1 AND tom_id = $2 AND property_id = $3
            RETURNING *;
    END;
$delete_subscription$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00010(tx *sql.Tx) error {
	query := `DROP FUNCTION delete_subscription(uuid, uuid, uuid);`
	return execQuery(query, tx)
}

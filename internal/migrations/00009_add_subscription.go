package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00009, down00009)
}

func up00009(tx *sql.Tx) error {
	query := `-- Add function for subscription
CREATE OR REPLACE FUNCTION set_subscription(uuid, uuid, uuid) RETURNS SETOF subscriptions AS $set_subscription$
    BEGIN
        RETURN QUERY
            WITH insertion AS (
                INSERT INTO subscriptions (consumer_id, tom_id, property_id)
                VALUES ($1, $2, $3)
                ON CONFLICT(consumer_id, tom_id, property_id) DO NOTHING
                RETURNING *
            )
            SELECT * FROM insertion
            UNION
            SELECT * FROM subscriptions WHERE consumer_id = $1 AND tom_id = $2 AND property_id = $3;
    END;
$set_subscription$ LANGUAGE plpgsql;`
	return execQuery(query, tx)
}

func down00009(tx *sql.Tx) error {
	query := `DROP FUNCTION set_subscription(uuid, uuid, uuid);`
	return execQuery(query, tx)
}

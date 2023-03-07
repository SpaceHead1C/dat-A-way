package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up00008, down00008)
}

func up00008(tx *sql.Tx) error {
	query := `-- Subscriptions
CREATE TABLE subscriptions (
	consumer_id uuid NOT NULL,
	tom_id uuid NOT NULL,
	property_id uuid NOT NULL,
	UNIQUE(consumer_id, tom_id, property_id),
	CONSTRAINT fk_consumer
		FOREIGN KEY(consumer_id)
			REFERENCES consumers(id),
	CONSTRAINT fk_tom
		FOREIGN KEY(tom_id)
			REFERENCES toms(id)
);`
	return execQuery(query, tx)
}

func down00008(tx *sql.Tx) error {
	query := `DROP TABLE subscriptions;`
	return execQuery(query, tx)
}

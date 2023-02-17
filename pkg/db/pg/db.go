package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Address      string
	Port         uint
	User         string
	Password     string
	DatabaseName string
}

func NewConnConfig(c Config) (*pgx.ConnConfig, error) {
	return pgx.ParseConfig(connectionString(c))
}

func connectionString(c Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Password, c.Address, c.Port, c.DatabaseName)
}

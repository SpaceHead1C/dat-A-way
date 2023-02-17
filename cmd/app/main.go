package main

import (
	"context"
	"dataway/pkg/db/pg"
	"dataway/pkg/log"
	"os"
	"time"
)

func main() {
	c := newConfig()
	if err := parse(os.Args[1:], c); err != nil {
		panic(err.Error())
	}
	l, err := log.NewLogger()
	if err != nil {
		panic(err.Error())
	}
	dbCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	db, err := pg.NewDB(dbCtx, pg.Config{
		Address:      c.PostgresAddress,
		Port:         c.PostgresPort,
		User:         c.PostgresUser,
		Password:     c.PostgresPassword,
		DatabaseName: c.PostgresDBName,
	})
	cancel()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}
	l.Info("dat(A)way service is up")
}

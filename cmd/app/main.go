package main

import (
	"context"
	"dataway/internal/adapter/pg"
	"dataway/internal/migrations"
	pkgpg "dataway/pkg/db/pg"
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
	dbCC, err := pkgpg.NewConnConfig(pkgpg.Config{
		Address:      c.PostgresAddress,
		Port:         c.PostgresPort,
		User:         c.PostgresUser,
		Password:     c.PostgresPassword,
		DatabaseName: c.PostgresDBName,
	})
	if err != nil {
		panic(err.Error())
	}
	if err := migrations.UpMigrations(dbCC); err != nil {
		panic(err.Error())
	}
	dbCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	repo, err := pg.NewRepository(dbCtx, pg.Config{
		ConnectConfig: dbCC,
		Logger:        l,
	})
	cancel()
	if err != nil {
		panic(err.Error())
	}
	defer repo.Close(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := repo.Ping(ctx); err != nil {
		panic(err.Error())
	}
	l.Info("dat(A)way service is up")
}

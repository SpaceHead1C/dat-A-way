package test

import (
	"context"
	"dataway/internal/adapter/pg"
	apg "dataway/internal/adapter/pg"
	"dataway/internal/api"
	pkgpg "dataway/pkg/db/pg"
	"dataway/pkg/log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/subosito/gotenv"
)

func newPgRepo(t *testing.T) *apg.Repository {
	if err := gotenv.Load(); err != nil {
		t.Fatal(err)
	}
	l, err := log.NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	port, _ := strconv.Atoi(os.Getenv("TEST_POSTGRES_PORT"))
	dbCC, err := pkgpg.NewConnConfig(pkgpg.Config{
		Address:      os.Getenv("TEST_POSTGRES_HOST"),
		Port:         uint(port),
		User:         os.Getenv("TEST_POSTGRES_USER"),
		Password:     os.Getenv("TEST_POSTGRES_PASSWORD"),
		DatabaseName: os.Getenv("TEST_POSTGRES_DB"),
	})
	if err != nil {
		panic(err.Error())
	}
	dbCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	out, err := pg.NewRepository(dbCtx, pg.Config{
		ConnectConfig: dbCC,
		Logger:        l,
	})
	cancel()
	if err != nil {
		panic(err.Error())
	}
	return out
}

func newTestConsumerManager(t *testing.T) *api.ConsumerManager {
	repo := newPgRepo(t)
	out, err := api.NewConsumerManager(api.ConsumerConfig{
		Repository: repo,
		Timeout:    time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

package main

import (
	"context"
	"dataway/grpc"
	"dataway/internal/adapter/pg"
	"dataway/internal/api"
	"dataway/internal/migrations"
	pkgpg "dataway/pkg/db/pg"
	pkglog "dataway/pkg/log"
	"dataway/rest"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	c := newConfig()
	if err := parse(os.Args[1:], c); err != nil {
		log.Fatalf("arguments parse error: %s", err)
	}
	l, err := pkglog.NewLogger()
	if err != nil {
		log.Fatalf("logger constructor error: %s", err)
	}
	dbCC, err := pkgpg.NewConnConfig(pkgpg.Config{
		Address:      c.PostgresAddress,
		Port:         c.PostgresPort,
		User:         c.PostgresUser,
		Password:     c.PostgresPassword,
		DatabaseName: c.PostgresDBName,
	})
	if err != nil {
		l.Fatal(err.Error())
	}
	if err := migrations.UpMigrations(dbCC); err != nil {
		l.Fatal(err.Error())
	}
	dbCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	repo, err := pg.NewRepository(dbCtx, pg.Config{
		ConnectConfig: dbCC,
		Logger:        l,
	})
	cancel()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer repo.Close(context.Background())
	l.Info("repository configured")

	consumerManager, err := api.NewConsumerManager(api.ConsumerConfig{
		Repository: repo,
		Timeout:    time.Second,
	})
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("consumers manager configured")

	tomManager, err := api.NewTomManager(api.TomConfig{
		Repository: repo,
		Timeout:    time.Second,
	})
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("toms manager configured")

	restServer, err := rest.NewServer(rest.Config{
		Logger:          l,
		Port:            c.RESTPort,
		Timeout:         time.Second * time.Duration(c.RESTTimeoutSec),
		ConsumerManager: consumerManager,
	})
	if err != nil {
		l.Fatal(err.Error())
	}

	grpcServer, err := grpc.NewServer(grpc.Config{
		Logger:     l,
		Port:       c.GRPCPort,
		TomManager: tomManager,
	})

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		err := restServer.Serve()
		return fmt.Errorf("REST server error: %w", err)
	})
	l.Infof("REST server listens at port: %d", c.RESTPort)
	g.Go(func() error {
		err := grpcServer.Serve()
		return fmt.Errorf("gRPC server error: %w", err)
	})
	l.Infof("gRPC server listens at port: %d", c.GRPCPort)

	l.Info("dat(A)way service is up")

	if err := g.Wait(); err != nil {
		l.Fatal(err.Error())
	}
}

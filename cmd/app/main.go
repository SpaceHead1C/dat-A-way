package main

import (
	"context"
	"dataway/grpc"
	"dataway/internal/adapter/pg"
	"dataway/internal/api"
	"dataway/internal/migrations"
	pkgpg "dataway/pkg/db/pg"
	"dataway/pkg/log"
	"dataway/rest"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
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
	l.Info("repository configured")

	consumerManager, err := api.NewConsumerManager(api.ConsumerConfig{
		Repository: repo,
		Timeout:    time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	l.Info("consumers manager configured")

	tomManager, err := api.NewTomManager(api.TomConfig{
		Repository: repo,
		Timeout:    time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	l.Info("toms manager configured")

	restServer, err := rest.NewServer(rest.Config{
		Logger:          l,
		Port:            c.RESTPort,
		Timeout:         time.Second * time.Duration(c.RESTTimeoutSec),
		ConsumerManager: consumerManager,
	})
	if err != nil {
		panic(err.Error())
	}

	grpcServer, err := grpc.NewServer(grpc.Config{
		Logger:     l,
		Port:       c.GRPCPort,
		TomManager: tomManager,
	})

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		err := restServer.Serve()
		l.Errorln("REST server error: ", err.Error())
		return err
	})
	l.Infof("REST server listens at port: %d", c.RESTPort)
	g.Go(func() error {
		err := grpcServer.Serve()
		l.Errorln("gRPC server error: ", err.Error())
		return err
	})
	l.Infof("gRPC server listens at port: %d", c.GRPCPort)

	l.Info("dat(A)way service is up")

	if err := g.Wait(); err != nil {
		panic(err.Error())
	}
}

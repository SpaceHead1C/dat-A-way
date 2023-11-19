package test

import (
	"context"
	"dataway/internal/adapter/pg"
	apg "dataway/internal/adapter/pg"
	"dataway/internal/api"
	"dataway/internal/pb"
	pkgpg "dataway/pkg/db/pg"
	"dataway/pkg/log"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/subosito/gotenv"
)

type grpcConfig struct {
	address string
	port    uint
}

var (
	gc grpcConfig
)

func init() {
	if err := gotenv.Load(); err != nil {
		panic(err.Error())
	}
	port, err := strconv.Atoi(os.Getenv("TEST_GRPC_PORT"))
	if err != nil {
		panic(err.Error())
	}
	gc = grpcConfig{
		address: os.Getenv("TEST_GRPC_ADDRESS"),
		port:    uint(port),
	}
}

func (c grpcConfig) dial() (*grpc.ClientConn, error) {
	return grpc.Dial(
		fmt.Sprintf("%s:%d", c.address, c.port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

func newPgRepo(t *testing.T) *apg.Repository {
	if err := gotenv.Load(); err != nil {
		t.Fatal(err)
	}
	l, err := log.New()
	if err != nil {
		t.Fatal(err)
	}
	port, err := strconv.Atoi(os.Getenv("TEST_POSTGRES_PORT"))
	if err != nil {
		t.Fatal(err)
	}
	pool, err := pkgpg.NewPoolConfig(pkgpg.Config{
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
		ConnectConfig: pool,
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

func newTestTomManager(t *testing.T) *api.TomManager {
	repo := newPgRepo(t)
	out, err := api.NewTomManager(api.TomConfig{
		Repository: repo,
		Timeout:    time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func newTestSubscriptionManager(t *testing.T) *api.SubscriptionManager {
	repo := newPgRepo(t)
	out, err := api.NewSubscriptionManager(api.SubscriptionConfig{
		Repository: repo,
		Timeout:    time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func newGRPCClient(t *testing.T) (pb.DatawayClient, *grpc.ClientConn) {
	conn, err := gc.dial()
	if err != nil {
		t.Fatal(err)
	}
	return pb.NewDatawayClient(conn), conn
}

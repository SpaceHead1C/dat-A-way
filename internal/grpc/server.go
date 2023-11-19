package grpc

import (
	"context"
	"fmt"
	"net"

	"dataway/internal/api"
	"dataway/internal/domain"
	. "dataway/internal/pb"
	"dataway/pkg/log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	UnimplementedDatawayServer
	logger   *log.Logger
	srv      *grpc.Server
	listener net.Listener
	port     uint

	tomManager          *api.TomManager
	subscriptionManager *api.SubscriptionManager
}

type Config struct {
	Logger *log.Logger
	Port   uint

	TomManager          *api.TomManager
	SubscriptionManager *api.SubscriptionManager
}

func NewServer(c Config) (domain.Server, error) {
	l := c.Logger
	if l == nil {
		l = log.GlobalLogger()
	}
	if c.TomManager == nil {
		return nil, fmt.Errorf("toms manager must be not nil")
	}
	if c.SubscriptionManager == nil {
		return nil, fmt.Errorf("subscriptions manager must be not nil")
	}
	return &server{
		logger: l,
		port:   c.Port,

		tomManager:          c.TomManager,
		subscriptionManager: c.SubscriptionManager,
	}, nil
}

func (s *server) Serve() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	s.listener = listener
	s.srv = grpc.NewServer()
	RegisterDatawayServer(s.srv, s)
	return s.srv.Serve(listener)
}

func (s *server) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) RegisterNewTom(ctx context.Context, req *RegisterTomRequest) (*UUID, error) {
	ctx = log.ContextWithLogger(ctx, s.logger)
	return RegisterNewTom(ctx, req, s.tomManager)
}

func (s *server) UpdateTom(ctx context.Context, req *UpdateTomRequest) (*emptypb.Empty, error) {
	ctx = log.ContextWithLogger(ctx, s.logger)
	return &emptypb.Empty{}, UpdateTom(ctx, req, s.tomManager)
}

func (s *server) Subscribe(ctx context.Context, req *Subscription) (*Subscription, error) {
	ctx = log.ContextWithLogger(ctx, s.logger)
	return Subscribe(ctx, s.subscriptionManager, req)
}

func (s *server) DeleteSubscription(ctx context.Context, req *Subscription) (*emptypb.Empty, error) {
	ctx = log.ContextWithLogger(ctx, s.logger)
	return &emptypb.Empty{}, DeleteSubscription(ctx, s.subscriptionManager, req)
}

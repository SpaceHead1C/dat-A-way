package grpc

import (
	"context"
	"dataway/internal/domain"
	"dataway/internal/pb"
	"dataway/pkg/log"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type server struct {
	pb.UnimplementedDatawayServer
	logger   *zap.SugaredLogger
	srv      *grpc.Server
	listener net.Listener
	port     uint
}

type Config struct {
	Logger *zap.SugaredLogger
	Port   uint
}

func NewServer(c Config) (domain.Server, error) {
	var err error
	l := c.Logger
	if l == nil {
		l, err = log.NewLogger()
		if err != nil {
			return nil, err
		}
	}
	return &server{
		logger: l,
		port:   c.Port,
	}, nil
}

func (s *server) Serve() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	s.listener = listener
	s.srv = grpc.NewServer()
	pb.RegisterDatawayServer(s.srv, s)
	return s.srv.Serve(listener)
}

func (s *server) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

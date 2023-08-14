package grpc

import (
	"context"
	"dataway/internal/api"
	"dataway/internal/pb"
	"dataway/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RegisterNewTom(ctx context.Context, man *api.TomManager) (*pb.UUID, error) {
	id, err := man.Add(ctx)
	if err != nil {
		log.LoggerFromContext(ctx).Errorf("new tom register error: %s", err)
		return nil, status.Errorf(codes.Internal, "new tom register error")
	}
	return pb.UUIDToPb(id), nil
}

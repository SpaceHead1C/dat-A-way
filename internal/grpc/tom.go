package grpc

import (
	"context"
	"errors"

	"dataway/internal/api"
	"dataway/internal/domain"
	"dataway/internal/pb"
	"dataway/pkg/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RegisterNewTom(ctx context.Context, req *pb.RegisterTomRequest, man *api.TomManager) (*pb.UUID, error) {
	id, err := man.Add(ctx, domain.RegisterTomRequest{Name: req.Name})
	if err != nil {
		log.LoggerFromContext(ctx).Errorf("new tom register error: %s", err)
		if errors.Is(err, domain.ErrDuplicate) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "new tom register error")
	}
	return pb.UUIDToPb(id), nil
}

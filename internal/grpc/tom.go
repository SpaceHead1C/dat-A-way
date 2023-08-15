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
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is empty")
	}
	id, err := man.Add(ctx, domain.RegisterTomRequest{Name: req.Name})
	if err != nil {
		if errors.Is(err, domain.ErrDuplicate) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		log.LoggerFromContext(ctx).Errorf("new tom register error: %s", err)
		return nil, status.Errorf(codes.Internal, "new tom register error")
	}
	return pb.UUIDToPb(id), nil
}

func UpdateTom(ctx context.Context, req *pb.UpdateTomRequest, man *api.TomManager) error {
	r, err := updateTomRequestFromPb(req)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	if _, err := man.Update(ctx, r); err != nil {
		switch true {
		case errors.Is(err, domain.ErrDuplicate):
			return status.Errorf(codes.AlreadyExists, err.Error())
		case errors.Is(err, domain.ErrNotFound):
			return status.Errorf(codes.NotFound, err.Error())
		default:
			log.LoggerFromContext(ctx).Errorf("tom update error: %s", err)
			return status.Errorf(codes.Internal, "tom update error")
		}
	}
	return nil
}

func updateTomRequestFromPb(in *pb.UpdateTomRequest) (domain.UpdateTomRequest, error) {
	out := domain.UpdateTomRequest{
		Name: in.Name,
	}
	if in.Name == "" {
		return out, errors.New("name is empty")
	}
	id, err := pb.UUIDFromPb(in.Id)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, "parse tom ID error: %s", err)
	}
	out.ID = id
	return out, nil
}

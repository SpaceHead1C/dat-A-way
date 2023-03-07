package pb

import (
	"github.com/google/uuid"
)

func UUIDToPb(in uuid.UUID) *UUID {
	return &UUID{Value: in[:]}
}

func UUIDFromPb(in *UUID) (uuid.UUID, error) {
	return uuid.FromBytes(in.Value)
}

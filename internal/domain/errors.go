package domain

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrConsumerNotFound = fmt.Errorf("consumer %w", ErrNotFound)

	ErrExpected = errors.New("expected")

	ErrDuplicate     = errors.New("duplicate")
	ErrNameDuplicate = fmt.Errorf("name %w", ErrDuplicate)
)

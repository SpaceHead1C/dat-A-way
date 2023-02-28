package domain

import "fmt"

var (
	ErrNotFound         = fmt.Errorf("not found")
	ErrConsumerNotFound = fmt.Errorf("consumer %w", ErrNotFound)

	ErrExpected = fmt.Errorf("expected")
)

package repository

import "errors"

var (
	// ErrMetadataNotFound is returned when metadata is not found in the repository.
	ErrNotFound = errors.New("not found")
)

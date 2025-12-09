package repository

import "errors"

// ErrNotFound is returned when no ratings are found for a
// record.
var ErrNotFound = errors.New("ratings not found for a record")

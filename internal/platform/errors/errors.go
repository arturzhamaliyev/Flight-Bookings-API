package errors

import "errors"

var (
	// ErrUniqueViolation is returned when the data already exists.
	ErrUniqueViolation = errors.New("duplicate key value violates unique constraint")
	// ErrNoRows is returned when no data found.
	ErrNoRows = errors.New("no rows in result set")
)

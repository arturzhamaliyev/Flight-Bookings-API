package errors

import "errors"

// ErrUniqueViolation is returned when the data already exists.
var ErrUniqueViolation = errors.New("duplicate key value violates unique constraint")

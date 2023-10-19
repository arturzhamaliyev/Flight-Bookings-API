package errors

import (
	"errors"
	"fmt"
)

var (
	// ErrUnknown is returned when an unexpected error occurs.
	ErrUnknown = errors.New("err_unknown: unknown error occurred")
	// ErrInvalidRequest is returned when either the parameters or the request body is invalid.
	ErrInvalidRequest = errors.New("err_invalid_request: invalid request received")
	// ErrValidation is returned when the parameters don't pass validation.
	ErrValidation = errors.New("err_validation: failed validation")
	// ErrNotFound is returned when the requested resource is not found.
	ErrNotFound = errors.New("err_not_found: not found")
)

func Wrap(err1, err2 error) error {
	return fmt.Errorf("%v: %w", err1, err2)
}

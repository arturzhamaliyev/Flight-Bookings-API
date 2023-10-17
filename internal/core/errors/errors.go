package errors

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// ErrUnknown is returned when an unexpected error occurs.
	ErrUnknown = Error("err_unknown: unknown error occurred")
	// ErrInvalidRequest is returned when either the parameters or the request body is invalid.
	ErrInvalidRequest = Error("err_invalid_request: invalid request received")
	// ErrValidation is returned when the parameters don't pass validation.
	ErrValidation = Error("err_validation: failed validation")
	// ErrNotFound is returned when the requested resource is not found.
	ErrNotFound = Error("err_not_found: not found")
)

// ErrSeperator is used to determine the boundaries of the errors in the hierarchy.
const ErrSeperator = " -- "

// Error allows errors to be defined as const errors preventing modification
// and allowing them to be evaluated against wrapped errors.
type Error string

func (s Error) Error() string {
	return string(s)
}

// Wrap allows errors to wrap an error returned from a 3rd party in
// a const service error preserving the original cause.
func (s Error) Wrap(err error) error {
	return wrappedError{cause: err, msg: string(s)}
}

// Is implements https://golang.org/pkg/errors/#Is allowing a Error
// to check it is the same even when wrapped. This implementation only
// checks the top most wrapped error.
func (s Error) Is(target error) bool {
	return s.Error() == target.Error() || strings.HasPrefix(target.Error(), s.Error()+ErrSeperator)
}

// wrappedError is an internal error type that allows the wrapping of
// underlying errors with Errors.
type wrappedError struct {
	cause error
	msg   string
}

func (w wrappedError) Error() string {
	if w.cause != nil {
		return fmt.Sprintf("%s%s%v", w.msg, ErrSeperator, w.cause)
	}
	return w.msg
}

// Is just wraps errors.Is as we don't want to alias the errors package everywhere to use it.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

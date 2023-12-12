package errors

import "errors"

var (
	// ErrUniqueViolation is returned when the data already exists.
	ErrUniqueViolation = errors.New("duplicate key value violates unique constraint")
	// ErrNoRows is returned when no data found.
	ErrNoRows = errors.New("no rows in result set")

	// ErrInvalidToken is returned when jwt token is invalid.
	ErrInvalidToken = errors.New("invalid token provided")
	// ErrNotAdminRole is returned when user is not an admin.
	ErrNotAdminRole = errors.New("not admin")
	// ErrNotCustomer is returned when user in not a customer.
	ErrNotCustomer = errors.New("not customer")
)

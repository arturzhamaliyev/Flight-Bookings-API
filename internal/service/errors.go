package service

import "errors"

var (
	ErrInvalidEmailAddress    = errors.New("invalid email address")
	ErrUserExists             = errors.New("user already exists")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)

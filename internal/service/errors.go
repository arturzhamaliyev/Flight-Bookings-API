package service

import "errors"

var (
	ErrInvalidEmailAddress = errors.New("invalid email address")
	ErrUserExists          = errors.New("user already exists")
)

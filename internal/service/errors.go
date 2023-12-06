package service

import "errors"

var (
	ErrInvalidEmailAddress = errors.New("invalid email address")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrUserExists          = errors.New("user already exists")
)

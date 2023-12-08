package service

import "errors"

var (
	ErrInvalidEmailAddress = errors.New("invalid email address")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrUserExists          = errors.New("user already exists")
	ErrUserNotFound        = errors.New("user doesn't exist")
	ErrHashPassword        = errors.New("failed to hash password")
)

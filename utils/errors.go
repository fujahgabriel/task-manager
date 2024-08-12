// utils/errors.go

package utils

import "errors"

// Define custom errors
var (
	ErrNotFound     = errors.New("resource not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized access")
)

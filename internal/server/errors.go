package server

import "errors"

var (
	ErrEmptyPackSizes  = errors.New("empty pack sizes")
	ErrInvalidPackSize = errors.New("invalid pack size")
	ErrInvalidTotal    = errors.New("invalid total")
)

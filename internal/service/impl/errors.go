package impl

import "errors"

// Các error constants
var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrInvalidProductType = errors.New("invalid product type")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrNotFound           = errors.New("resource not found")
	ErrServerError        = errors.New("internal server error")
) 
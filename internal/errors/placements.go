package errors

import "errors"

var (
	ErrProductOutOfStock      = errors.New("product out of stock")
	ErrProductReserveNotFound = errors.New("product reserve not found")
)

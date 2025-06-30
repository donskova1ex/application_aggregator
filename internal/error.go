package internal

import "errors"

var (
	ErrRecordNotFound    = errors.New("no record in result")
	UUIDValidationFailed = errors.New("uuid validation failed")
)

package internal

import "errors"

var (
	ErrRecordNotFound  = errors.New("no record in result")
	ErrUUIDValidation  = errors.New("uuid validation failed")
	ErrPhoneValidation = errors.New("phone validation failed")
	ErrRecordInsert    = errors.New("record insert failed")
)

package internal

import "errors"

var (
	ErrRecordNotFound            = errors.New("no record in result")
	ErrUUIDValidation            = errors.New("uuid validation failed")
	ErrPhoneValidation           = errors.New("phone validation failed")
	ErrOrganizationNameDuplicate = errors.New("organization name duplicating")
	ErrIncomingOrganizationUUID  = errors.New("incoming organization uuid validation failed")
	ErrIssueOrganizationUUID     = errors.New("issue organization uuid validation failed")
	ErrPositiveValue             = errors.New("the value does not correspond to the minimum")
	ErrEntityUUIDDuplicate       = errors.New("uuid duplicate")
)

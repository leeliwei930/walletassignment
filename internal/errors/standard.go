package errors

import (
	"github.com/pkg/errors"
)

func UnexpectedError(details error) error {
	return NewStandardError("UNEXPECTED_ERROR", "An unexpected error occurred", details)
}

type StandardError struct {
	Code        string `json:"errorCode,omitempty"`
	Description string `json:"description,omitempty"`
	Details     error  `json:"-"`
}

func (se *StandardError) Error() string {
	var errDetails = errors.New("No error details provided")
	if se.Details != nil {
		errDetails = se.Details
	}

	return errors.Wrapf(errDetails, "StandardError - [%s]", se.Code).Error()
}

func NewStandardError(code string, description string, details error) *StandardError {
	stdErr := &StandardError{
		Code:        code,
		Description: description,
		Details:     details,
	}
	return stdErr
}

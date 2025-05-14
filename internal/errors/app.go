package errors

import (
	"github.com/pkg/errors"
)

type ApplicationError struct {
	StandardError `json:",inline"`
}

func (ae *ApplicationError) Error() error {
	return errors.WithMessage(ae.Details, "ApplicationError:")
}

package errors

type InvalidRequestError struct {
	*StandardError `json:",inline"`
}

type InvalidRequestErrorOption = func(ve *InvalidRequestError)

func NewInvalidRequestError(code string, description string, details error, opts ...InvalidRequestErrorOption) *InvalidRequestError {
	invalidReqErr := &InvalidRequestError{
		StandardError: NewStandardError(
			code,
			description,
			details,
		),
	}
	return invalidReqErr
}

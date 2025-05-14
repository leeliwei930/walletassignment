package errors

type UnauthorizedRequestError struct {
	*StandardError `json:",inline"`
}

type UnauthorizedRequestErrorOption = func(ure *UnauthorizedRequestError)

func NewUnauthorizedRequestError(description string, details error, opts ...UnauthorizedRequestError) *UnauthorizedRequestError {
	invalidReqErr := &UnauthorizedRequestError{
		StandardError: NewStandardError(
			"UNAUTHORIZED_ERR",
			description,
			details,
		),
	}
	return invalidReqErr
}

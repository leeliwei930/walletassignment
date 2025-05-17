package errors

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	*StandardError `json:",inline"`
	Errors         map[string]string `json:"errors,omitempty"`
	Message        string
}

type ValidationErrorOption = func(ve *InvalidRequestError)

func NewValidationError(errs validator.ValidationErrors, ut ut.Translator) *ValidationError {
	errorBag := make(map[string]string)
	for _, err := range errs {
		localized := err.Translate(ut)
		localizedFieldName, _ := ut.T(err.Field())
		localizedFieldKey, _ := ut.T(err.Field() + "::label")
		localized = strings.Replace(localized, err.Field(), localizedFieldName, 1)

		if strings.Contains(ut.Locale(), "en") {
			errorBag[localizedFieldKey] = strings.ToUpper(localized[:1]) + localized[1:]
		} else {
			errorBag[err.StructField()] = localized
		}
	}

	localizedMsg, _ := ut.T("errors::validation::message")

	return &ValidationError{
		NewStandardError("INVALID_REQUEST", "", errs),
		errorBag,
		localizedMsg,
	}
}

func (ve *ValidationError) Error() string {
	return ve.StandardError.Error()
}

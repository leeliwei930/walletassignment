package validation

import (
	"bytes"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ErrorBag []validator.FieldError
type FieldErrorMessages map[string]string

func FromValidationErrors(errs validator.ValidationErrors) ErrorBag {
	eb := make(ErrorBag, len(errs))
	copy(eb, errs)
	return eb
}

func (eb ErrorBag) Error() string {

	buff := bytes.NewBufferString("")

	for i := 0; i < len(eb); i++ {

		buff.WriteString(eb[i].Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

func (eb ErrorBag) Translate(ut ut.Translator) FieldErrorMessages {
	trans := make(FieldErrorMessages)

	for i, err := range eb {
		fe := eb[i]
		trans[fe.Field()] = err.Translate(ut)
	}

	return trans
}

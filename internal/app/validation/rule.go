package validation

import (
	"context"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Rule interface {
	ValidationFunc(ctx context.Context, fl validator.FieldLevel) bool
	Name() string
	TranslationFunc(ut ut.Translator, fe validator.FieldError) string
}

package interfaces

import (
	"context"
	"net/http"

	universalTranslator "github.com/go-playground/universal-translator"
)

type Locale interface {
	GetTranslatorFromRequest(r *http.Request) universalTranslator.Translator
	GetTranslatorFromContext(ctx context.Context) universalTranslator.Translator
	GetUT() *universalTranslator.UniversalTranslator
}

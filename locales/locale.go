package locales

import (
	"context"
	"net/http"

	universalTranslator "github.com/go-playground/universal-translator"
	pkgappcontext "github.com/leeliwei930/walletassignment/internal/app/context"
)

type localeTranslator struct {
	Translator *universalTranslator.UniversalTranslator
}

func NewLocaleTranslator(ut *universalTranslator.UniversalTranslator) *localeTranslator {
	return &localeTranslator{
		Translator: ut,
	}
}

func (locale *localeTranslator) GetTranslatorFromRequest(r *http.Request) universalTranslator.Translator {
	lang := r.Header.Get("Accept-Language")
	translator, found := locale.Translator.GetTranslator(lang)
	if !found || lang == "" {
		return locale.Translator.GetFallback()
	}

	return translator
}

const AcceptLang string = "accept_lang"

func (locale *localeTranslator) GetTranslatorFromContext(ctx context.Context) universalTranslator.Translator {
	lang := ""
	if appCtx, err := pkgappcontext.GetApplicationContext(ctx); err == nil {
		lang = appCtx.GetLanguage()
	}

	translator, found := locale.Translator.GetTranslator(lang)
	if !found || lang == "" {
		return locale.Translator.GetFallback()
	}

	return translator
}

func (locale *localeTranslator) GetUT() *universalTranslator.UniversalTranslator {
	return locale.Translator
}

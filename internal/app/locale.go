package app

import (
	"github.com/go-playground/locales/en"
	zhCN "github.com/go-playground/locales/zh"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/locales"
	pkgerrors "github.com/pkg/errors"
)

func (app *application) InitLocale() error {

	enLocale := en.New()
	zhLocale := zhCN.New()
	ut := universalTranslator.New(enLocale, enLocale, zhLocale)

	err := ut.Import(universalTranslator.FormatJSON, "locales")
	if err != nil {
		errMsg := "Failed to load locale files when initialize localization"
		return errors.NewStandardError(
			"APPLICATION_INIT_LOCALE_ERR",
			errMsg,
			pkgerrors.Wrap(err, errMsg),
		)
	}

	err = ut.VerifyTranslations()
	if err != nil {
		errMsg := "Locale files structure validation errors"
		return errors.NewStandardError(
			"APPLICATION_INIT_LOCALE_ERR",
			errMsg,
			pkgerrors.Wrap(err, errMsg),
		)
	}

	app.locale = locales.NewLocaleTranslator(ut)
	return nil
}

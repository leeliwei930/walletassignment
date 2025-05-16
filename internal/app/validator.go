package app

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/leeliwei930/walletassignment/internal/app/validation"

	pkgut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	pkgerrors "github.com/pkg/errors"
)

func (app *application) InitValidator() error {

	if app.locale == nil {
		return pkgerrors.New("Unable to initialize validator without initialize localization, please call app.InitLocale() before app.InitValidator()")
	}
	locale := app.GetLocale()
	ut := locale.GetUT()
	enTrans, _ := ut.GetTranslator("en")
	zhTrans, _ := ut.GetTranslator("zh")

	v := validator.New()
	enTransErr := en_translations.RegisterDefaultTranslations(v, enTrans)
	if enTransErr != nil {
		return pkgerrors.Wrap(enTransErr, "Error when register translation error message on locale: en")
	}
	zhTransErr := zh_translations.RegisterDefaultTranslations(v, zhTrans)
	if zhTransErr != nil {
		return pkgerrors.Wrap(zhTransErr, "Error when register translation error message on locale: zh")
	}

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("localeKey")
		// skip if tag key says it should be ignored
		if name == "" {
			return field.Name
		}

		return name
	})

	rules := []validation.Rule{}

	for _, rule := range rules {
		registerRule(v, rule, []pkgut.Translator{enTrans, zhTrans})
	}

	app.validator = v
	return nil
}

func registerRule(v *validator.Validate, rule validation.Rule, trans []pkgut.Translator) error {

	err := v.RegisterValidationCtx(rule.Name(), rule.ValidationFunc)
	if err != nil {
		return pkgerrors.Wrapf(err, "Error when register validation rule for %s", rule.Name())
	}

	for _, trans := range trans {
		err = v.RegisterTranslation(rule.Name(), trans, func(ut pkgut.Translator) error {
			return nil
		}, func(ut pkgut.Translator, fe validator.FieldError) string {
			return rule.TranslationFunc(ut, fe)
		})
		if err != nil {
			return pkgerrors.Wrap(err, "Error when register translation error message on locale")
		}
	}

	return nil
}

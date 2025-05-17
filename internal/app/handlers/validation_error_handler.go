package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app/response"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
)

func ApplicationErrorHandler(app interfaces.Application) echo.HTTPErrorHandler {

	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		responder := response.NewResponder(c, app)

		if verr, isValidationErr := err.(validator.ValidationErrors); isValidationErr {
			locale := app.GetLocale()
			ut := locale.GetTranslatorFromRequest(c.Request())
			validationErr := errors.NewValidationError(verr, ut)

			_ = responder.ValidationError(c, *validationErr)
		} else {
			unexpectedError := errors.UnexpectedError(err)
			_ = responder.UnexpectedError(c, unexpectedError)
		}
	}
}

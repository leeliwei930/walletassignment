package handlers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app/response"
	"github.com/leeliwei930/walletassignment/internal/errors"

	"github.com/leeliwei930/walletassignment/internal/interfaces"
)

func ApplicationErrorHandler(app interfaces.Application) echo.HTTPErrorHandler {

	return func(err error, ec echo.Context) {
		if ec.Response().Committed {
			return
		}
		responder := response.NewResponder(ec, app)

		if echoHttpErr, isEchoHttpErr := err.(*echo.HTTPError); isEchoHttpErr {
			_ = responder.ErrorJSON(echoHttpErr.Code, response.ErrorResponse{
				ErrorCode: strconv.Itoa(echoHttpErr.Code),
				Message:   echoHttpErr.Error(),
			})
			return
		}

		if verr, isValidationErr := err.(validator.ValidationErrors); isValidationErr {
			locale := app.GetLocale()
			ut := locale.GetTranslatorFromRequest(ec.Request())
			validationErr := errors.NewValidationError(verr, ut)

			_ = responder.ValidationError(ec, *validationErr)
		} else if invalidReqErr, ok := err.(*errors.InvalidRequestError); ok {
			_ = responder.BadRequestError(ec, invalidReqErr)
		} else {
			_ = responder.UnexpectedError(ec, err)
		}
	}
}

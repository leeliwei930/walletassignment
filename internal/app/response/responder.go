package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	pkgerrors "github.com/pkg/errors"
)

const JSON_DEFAULT_INDENT_FORMAT = "  "

type Responder interface {
	AbortIfIncorrectJsonPayload(ec echo.Context, err error) error
	UnexpectedError(ec echo.Context, err error) error
	BadRequestError(ec echo.Context, err error) error
	UnauthorizedError(ec echo.Context, err error) error
	JSON(ec echo.Context, statusCode int, data any, indent string) error
}

type JSONResponse struct {
	StatusCode string `json:"statusCode,omitempty"`
	Data       any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	StatusCode string            `json:"statusCode,omitempty"`
	Message    string            `json:"message"`
	StackTrace []string          `json:"stackTrace,omitempty"`
	Fields     map[string]string `json:"fields,omitempty"`
}

type responder struct {
	Context echo.Context
	App     interfaces.Application
}

func NewResponder(ec echo.Context, app interfaces.Application) *responder {
	return &responder{Context: ec, App: app}
}

func (h *responder) AbortIfIncorrectJsonPayload(ec echo.Context, err error) error {
	if err != nil {
		err = pkgerrors.WithStack(err)
		stackTraces := fmt.Sprintf("%+v", err)
		stackTraces = strings.ReplaceAll(stackTraces, "\t", "  ")

		return h.ErrorJSON(
			http.StatusUnprocessableEntity, ErrorResponse{
				StatusCode: errorCodes[JSONEmptyBodyError],
				Message:    err.Error(),
				StackTrace: strings.Split(stackTraces, "\n"),
			},
		)
	}
	return nil
}

func (h *responder) UnexpectedError(ec echo.Context, err error) error {
	if err != nil {
		err = pkgerrors.WithStack(err)
		stackTraces := fmt.Sprintf("%+v", err)
		stackTraces = strings.ReplaceAll(stackTraces, "\t", "  ")

		return h.ErrorJSON(
			http.StatusInternalServerError, ErrorResponse{
				StatusCode: errorCodes[Unexpected],
				Message:    err.Error(),
				StackTrace: strings.Split(stackTraces, "\n"),
			},
		)
	}
	return nil
}

func (h *responder) BadRequestError(ec echo.Context, err error) error {
	if err != nil {
		if _, ok := err.(*errors.InvalidRequestError); ok {
			req := ec.Request()
			locale := h.App.GetLocale()
			ut := locale.GetTranslatorFromRequest(req)
			invalidReqErr := err.(*errors.InvalidRequestError)
			message, _ := ut.T(invalidReqErr.Description)
			return h.ErrorJSON(
				http.StatusBadRequest, ErrorResponse{
					StatusCode: errorCodes[BadRequestError],
					Message:    message,
				},
			)
		}

		return h.ErrorJSON(
			http.StatusBadRequest, ErrorResponse{
				StatusCode: errorCodes[BadRequestError],
				Message:    err.Error(),
			},
		)
	}
	return nil
}

func (h *responder) ValidationError(ec echo.Context, verr errors.ValidationError) error {

	return h.ErrorJSON(
		http.StatusBadRequest, ErrorResponse{
			StatusCode: errorCodes[ValidationError],
			Message:    verr.Message,
			Fields:     verr.Errors,
		},
	)
}

func (h *responder) UnauthorizedError(ec echo.Context) error {
	req := ec.Request()
	app := h.App
	locale := app.GetLocale().GetTranslatorFromRequest(req)
	message, _ := locale.T("errors::unauthorized")
	return h.ErrorJSON(
		http.StatusUnauthorized, ErrorResponse{
			StatusCode: errorCodes[UnauthorizedError],
			Message:    message,
		},
	)
}

func (h *responder) ErrorJSON(code int, errorRes ErrorResponse) error {
	return h.JSON(code, errorRes)
}

func (h *responder) JSON(code int, data interface{}) error {
	return h.Context.JSONPretty(code, data, JSON_DEFAULT_INDENT_FORMAT)
}

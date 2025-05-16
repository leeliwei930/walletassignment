package context

import (
	"context"

	"github.com/leeliwei930/walletassignment/constant"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
)

type applicationContext struct {
	Language      string
	AuthUserPhone string
}

var ApplicationCtxNotSetError = errors.NewStandardError(
	"APP_CTX_ERROR_100001",
	"Application context is not set",
	nil,
)

func New() interfaces.ApplicationContext {
	appCtx := &applicationContext{}
	return appCtx
}

func (actx *applicationContext) SetLanguage(language string) {
	actx.Language = language
}

func (actx *applicationContext) GetLanguage() string {
	return actx.Language
}

func (actx *applicationContext) GetAuthUserPhone() string {
	return actx.AuthUserPhone
}

func (actx *applicationContext) SetAuthUserPhone(phone string) {
	actx.AuthUserPhone = phone
}

func GetApplicationContext(ctx context.Context) (interfaces.ApplicationContext, error) {
	appCtx, ok := ctx.Value(constant.ApplicationContextKey).(*applicationContext)
	if !ok {
		return nil, ApplicationCtxNotSetError
	}

	return appCtx, nil
}

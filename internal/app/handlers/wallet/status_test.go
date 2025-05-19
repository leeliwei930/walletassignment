package wallet_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	svcinterfaces "github.com/leeliwei930/walletassignment/internal/app/services/interfaces"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WalletStatusHandlerTestSuite struct {
	suite.Suite
	client *http.Client
	srv    *httptest.Server
	app    interfaces.Application

	walletSvc *svcinterfaces.MockWalletService
	userSvc   *svcinterfaces.MockUserService
}

func (s *WalletStatusHandlerTestSuite) SetupTest() {
	app, err := app.InitializeFromEnv()
	s.NoError(err)

	s.client = http.DefaultClient

	s.walletSvc = svcinterfaces.NewMockWalletService(s.T())
	s.userSvc = svcinterfaces.NewMockUserService(s.T())
	app.SetWalletService(s.walletSvc)
	app.SetUserService(s.userSvc)

	ec := echo.New()
	ec = app.Routes(ec)
	ec = app.SetupMiddlewares(ec)
	s.srv = httptest.NewServer(ec.Server.Handler)
	s.app = app
}

func (s *WalletStatusHandlerTestSuite) TearDownTest() {
	s.srv.Close()
}

func (s *WalletStatusHandlerTestSuite) TestStatus_ShouldReturnWalletStatus() {
	userUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	walletUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")

	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60182119233").Return(userUUID, nil).Maybe()
	s.walletSvc.EXPECT().Status(mock.Anything, models.WalletStatusParams{
		UserID: userUUID,
	}).Return(&models.WalletStatus{
		ID:               walletUUID,
		Balance:          100000,
		Currency:         "USD",
		FormattedBalance: "USD 100.00",
	}, nil).Once()

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/wallet/status", s.srv.URL),
		nil,
	)
	req.Header.Set("X-USER-PHONE", "+60182119233")
	res, err := s.client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusOK, res.StatusCode)
}

func TestWalletStatusHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WalletStatusHandlerTestSuite))
}

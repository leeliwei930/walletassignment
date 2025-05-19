package wallet_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/app/handlers/wallet"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	svcinterfaces "github.com/leeliwei930/walletassignment/internal/app/services/interfaces"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	"github.com/leeliwei930/walletassignment/mocks/code.cloudfoundry.org/clock"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WalletDepositHandlerTestSuite struct {
	suite.Suite
	client *http.Client
	srv    *httptest.Server
	app    interfaces.Application

	clock     *clock.MockClock
	fixedTime time.Time

	walletSvc *svcinterfaces.MockWalletService
	userSvc   *svcinterfaces.MockUserService
}

func (s *WalletDepositHandlerTestSuite) SetupTest() {
	app, err := app.InitializeFromEnv()
	s.NoError(err)

	s.client = http.DefaultClient

	s.clock = clock.NewMockClock(s.T())
	s.fixedTime, err = time.Parse(time.RFC3339, "2025-05-19T10:00:00Z")
	s.NoError(err)

	s.clock.On("Now").Return(s.fixedTime).Maybe()
	app.SetClock(s.clock)

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

func (s *WalletDepositHandlerTestSuite) TearDownTest() {
	s.srv.Close()
}

func (s *WalletDepositHandlerTestSuite) TestDeposit_ShouldCallTheWalletServiceDeposit() {

	userUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	walletUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")
	transactionUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174002")

	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60182119233").Return(userUUID, nil).Maybe()
	s.walletSvc.EXPECT().Deposit(mock.Anything, models.WalletDepositParams{
		UserID: userUUID,
		Amount: 100,
	}).Return(&models.WalletDeposit{
		Wallet: models.WalletStatus{
			ID:               walletUUID,
			Balance:          100000,
			Currency:         "USD",
			FormattedBalance: "USD 100.00",
		},
		Transaction: models.WalletTransaction{
			ID:        transactionUUID,
			Amount:    100,
			Type:      "deposit",
			Timestamp: s.fixedTime,
		},
	}, nil).Once()

	payload := wallet.DepositRequest{
		Amount: 100,
	}
	payloadBytes, err := json.Marshal(payload)
	s.NoError(err)

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/wallet/deposit", s.srv.URL),
		bytes.NewBuffer(payloadBytes),
	)
	req.Header.Set("X-USER-PHONE", "+60182119233")
	req.Header.Set("Content-Type", "application/json")
	res, err := s.client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusOK, res.StatusCode)
}

func (s *WalletDepositHandlerTestSuite) TestDeposit_ShouldReturnErrorWhenAmountIsLessThanMinimum() {
	userUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	locale := s.app.GetLocale()
	ut := locale.GetUT().GetFallback()
	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60182119233").Return(userUUID, nil).Maybe()
	s.walletSvc.EXPECT().Deposit(mock.Anything, models.WalletDepositParams{
		UserID: userUUID,
		Amount: 50,
	}).Return(nil, errors.MinimumDepositAmountRequiredErr(ut, "USD 1.00")).Once()

	payload := wallet.DepositRequest{
		Amount: 50,
	}
	payloadBytes, err := json.Marshal(payload)
	s.NoError(err)

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/wallet/deposit", s.srv.URL),
		bytes.NewBuffer(payloadBytes),
	)
	req.Header.Set("X-USER-PHONE", "+60182119233")
	req.Header.Set("Content-Type", "application/json")
	res, err := s.client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusBadRequest, res.StatusCode)
}

func TestWalletDepositHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WalletDepositHandlerTestSuite))
}

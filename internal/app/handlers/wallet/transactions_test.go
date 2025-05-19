package wallet_test

import (
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
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	"github.com/leeliwei930/walletassignment/mocks/code.cloudfoundry.org/clock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WalletTransactionsHandlerTestSuite struct {
	suite.Suite
	client *http.Client
	srv    *httptest.Server
	app    interfaces.Application

	clock     *clock.MockClock
	fixedTime time.Time

	walletSvc *svcinterfaces.MockWalletService
	userSvc   *svcinterfaces.MockUserService
}

func (s *WalletTransactionsHandlerTestSuite) SetupTest() {
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

func (s *WalletTransactionsHandlerTestSuite) TearDownTest() {
	s.srv.Close()
}

func (s *WalletTransactionsHandlerTestSuite) TestGetTransactions_ShouldReturnTransactions() {
	userUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	transactionUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")

	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60182119233").Return(userUUID, nil).Maybe()

	expectedTransactions := &models.WalletTransactions{
		Transactions: []*models.WalletTransaction{
			{
				ID:        transactionUUID,
				Amount:    100,
				Type:      "transfer_out",
				Timestamp: s.fixedTime,
			},
		},
		Pagination: &models.Pagination{
			TotalItems:   1,
			ItemsPerPage: 20,
			CurrentPage:  2,
			TotalPages:   2,
		},
	}

	s.walletSvc.EXPECT().Transactions(mock.Anything, models.WalletTransactionsParams{
		UserID: userUUID,
		Limit:  20,
		Page:   2,
	}).Return(expectedTransactions, nil).Once()

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/wallet/transactions?limit=20&page=2", s.srv.URL),
		nil,
	)
	req.Header.Set("X-USER-PHONE", "+60182119233")
	res, err := s.client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusOK, res.StatusCode)

	var response wallet.WalletTransactionsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	s.NoError(err)

	s.Equal(expectedTransactions.Transactions, response.Data)
	s.Equal(expectedTransactions.Pagination, response.Pagination)
}

func (s *WalletTransactionsHandlerTestSuite) TestGetTransactions_ShouldUseDefaultPaginationParamsIfNotProvided() {
	userUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	transactionUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")

	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60182119233").Return(userUUID, nil).Maybe()

	expectedTransactions := &models.WalletTransactions{
		Transactions: []*models.WalletTransaction{
			{
				ID:        transactionUUID,
				Amount:    100,
				Type:      "transfer_out",
				Timestamp: s.fixedTime,
			},
		},
		Pagination: &models.Pagination{
			TotalItems:   1,
			ItemsPerPage: 10,
			CurrentPage:  1,
			TotalPages:   1,
		},
	}

	s.walletSvc.EXPECT().Transactions(mock.Anything, models.WalletTransactionsParams{
		UserID: userUUID,
		Limit:  10,
		Page:   1,
	}).Return(expectedTransactions, nil).Once()

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/wallet/transactions", s.srv.URL),
		nil,
	)
	req.Header.Set("X-USER-PHONE", "+60182119233")
	res, err := s.client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusOK, res.StatusCode)

	var response wallet.WalletTransactionsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	s.NoError(err)

	s.Equal(expectedTransactions.Transactions, response.Data)
	s.Equal(expectedTransactions.Pagination, response.Pagination)
}

func TestWalletTransactionsHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTransactionsHandlerTestSuite))
}

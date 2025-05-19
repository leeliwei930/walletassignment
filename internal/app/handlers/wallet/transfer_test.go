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

type WalletTransferHandlerTestSuite struct {
	suite.Suite
	client *http.Client
	srv    *httptest.Server
	app    interfaces.Application

	clock     *clock.MockClock
	fixedTime time.Time

	walletSvc *svcinterfaces.MockWalletService
	userSvc   *svcinterfaces.MockUserService
}

func (s *WalletTransferHandlerTestSuite) SetupTest() {
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

func (s *WalletTransferHandlerTestSuite) TearDownTest() {
	s.srv.Close()
}

func (s *WalletTransferHandlerTestSuite) TestTransfer_ShouldCallTheWalletServiceTransfer() {
	senderUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	recipientUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")
	walletUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174002")
	transactionUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174003")

	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60182119233").Return(senderUUID, nil).Maybe()
	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60123456789").Return(recipientUUID, nil).Once()
	s.walletSvc.EXPECT().Transfer(mock.Anything, models.WalletTransferParams{
		SenderUserID:           senderUUID,
		RecipientUserID:        recipientUUID,
		Amount:                 100,
		RecipientReferenceNote: "Payment for lunch",
	}).Return(&models.WalletTransfer{
		Wallet: models.WalletStatus{
			ID:               walletUUID,
			Balance:          100000,
			Currency:         "USD",
			FormattedBalance: "USD 100.00",
		},
		Transaction: models.WalletTransaction{
			ID:        transactionUUID,
			Amount:    100,
			Type:      "transfer_out",
			Timestamp: s.fixedTime,
		},
	}, nil).Once()

	payload := wallet.TransferRequest{
		Amount:                 100,
		RecipientPhoneNumber:   "+60123456789",
		RecipientReferenceNote: "Payment for lunch",
	}
	payloadBytes, err := json.Marshal(payload)
	s.NoError(err)

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/wallet/transfer", s.srv.URL),
		bytes.NewBuffer(payloadBytes),
	)
	req.Header.Set("X-USER-PHONE", "+60182119233")
	req.Header.Set("Content-Type", "application/json")
	res, err := s.client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusOK, res.StatusCode)
}

func (s *WalletTransferHandlerTestSuite) TestTransfer_ShouldReturnErrorWhenRecipientNotFound() {
	senderUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60182119233").Return(senderUUID, nil).Maybe()
	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60123456789").Return(uuid.Nil, errors.InvalidTransferRecipientPhoneNumberErr).Once()

	payload := wallet.TransferRequest{
		Amount:               100,
		RecipientPhoneNumber: "+60123456789",
	}
	payloadBytes, err := json.Marshal(payload)
	s.NoError(err)

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/wallet/transfer", s.srv.URL),
		bytes.NewBuffer(payloadBytes),
	)
	req.Header.Set("X-USER-PHONE", "+60182119233")
	req.Header.Set("Content-Type", "application/json")
	res, err := s.client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusBadRequest, res.StatusCode)
}

func (s *WalletTransferHandlerTestSuite) TestTransfer_ShouldReturnErrorWhenRequiredFieldsAreMissing() {
	senderUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	s.userSvc.EXPECT().GetUserIDByPhone(mock.Anything, "+60182119233").Return(senderUUID, nil).Maybe()

	payload := wallet.TransferRequest{}
	payloadBytes, err := json.Marshal(payload)
	s.NoError(err)

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/wallet/transfer", s.srv.URL),
		bytes.NewBuffer(payloadBytes),
	)
	req.Header.Set("X-USER-PHONE", "+60182119233")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en")
	res, err := s.client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusUnprocessableEntity, res.StatusCode)

	responseBody := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&responseBody)
	s.NoError(err)

	s.Equal("ERR_VALIDATION_422", responseBody["errorCode"])
	s.Equal("The information you provided contains errors. Please review and correct it.", responseBody["message"])

	errorFields := responseBody["fields"].(map[string]interface{})
	s.Equal("Transfer amount is a required field", errorFields["amount"])
	s.Equal("Recipient phone number is a required field", errorFields["recipientPhoneNumber"])
}

func TestWalletTransferHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTransferHandlerTestSuite))
}

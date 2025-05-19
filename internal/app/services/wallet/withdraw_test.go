package wallet_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent"
	pkgapp "github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	"github.com/leeliwei930/walletassignment/mocks/code.cloudfoundry.org/clock"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/suite"
)

type WalletWithdrawTestSuite struct {
	suite.Suite
	app       interfaces.Application
	clock     *clock.MockClock
	fixedTime time.Time
}

func (s *WalletWithdrawTestSuite) SetupTest() {
	app, err := pkgapp.InitializeFromEnv()
	s.NoError(err)

	s.clock = clock.NewMockClock(s.T())
	s.fixedTime, err = time.Parse(time.RFC3339, "2025-05-19T10:00:00Z")
	s.NoError(err)

	s.clock.On("Now").Return(s.fixedTime).Maybe()
	app.SetClock(s.clock)

	s.app = app
}

func (s *WalletWithdrawTestSuite) TestWalletWithdraw_ShouldWithdrawCorrectAmountFromUserWallet() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()
		// simulate initial deposit of USD 100.00
		_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: userRec.ID,
			Amount: 10000,
		})
		s.NoError(err)

		// simulate withdrawal of USD 10.00
		withdrawResponse, err := walletSvc.Withdraw(ctx, models.WalletWithdrawalParams{
			UserID: userRec.ID,
			Amount: 1000,
		})
		s.NoError(err)

		s.Equal(9000, withdrawResponse.Wallet.Balance)
		s.Equal("USD 90.00", withdrawResponse.Wallet.FormattedBalance)
		s.Equal(1000, withdrawResponse.Transaction.Amount)
		s.Equal("withdrawal", withdrawResponse.Transaction.Type)
		s.Empty(withdrawResponse.Transaction.RecipientReferenceNote)
		s.Equal(s.fixedTime, withdrawResponse.Transaction.Timestamp)
	})
	s.NoError(refreshDBErr)
}

func (s *WalletWithdrawTestSuite) TestWalletWithdraw_ShouldReturnError_WhenWithdrawAmountIsOverWalletBalance() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()
		// simulate initial deposit of USD 100.00
		_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: userRec.ID,
			Amount: 10000,
		})
		s.NoError(err)

		// simulate withdrawal of USD 10.00
		withdrawResponse, err := walletSvc.Withdraw(ctx, models.WalletWithdrawalParams{
			UserID: userRec.ID,
			Amount: 90000,
		})
		s.Nil(withdrawResponse)
		s.EqualError(err, errors.InsuficcientBalanceWithdrawalErr.Error())
	})
	s.NoError(refreshDBErr)
}

func (s *WalletWithdrawTestSuite) TestWalletWithdraw_ShouldReturnNotFoundError_WhenGivenUserIDIsInvalid() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()

		walletSvc := s.app.GetWalletService()
		withdrawResponse, err := walletSvc.Withdraw(ctx, models.WalletWithdrawalParams{
			UserID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Amount: 1000,
		})
		s.True(ent.IsNotFound(err))
		s.Nil(withdrawResponse)
	})
	s.NoError(refreshDBErr)
}

func (s *WalletWithdrawTestSuite) TestWalletWithdraw_ShouldReturnError_WhenWithdrawAmountIsBelowRequiredMinimumAmount() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		locale := s.app.GetLocale()
		ut := locale.GetUT().GetFallback()

		userSvc := s.app.GetUserService()
		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()
		// simulate initial deposit of USD 100.00
		_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: userRec.ID,
			Amount: 10000,
		})
		s.NoError(err)

		// simulate withdrawal of USD 1.00
		withdrawResponse, err := walletSvc.Withdraw(ctx, models.WalletWithdrawalParams{
			UserID: userRec.ID,
			Amount: 1,
		})
		s.Nil(withdrawResponse)
		s.EqualError(err, errors.MinimumWithdrawalAmountRequiredErr(
			ut,
			"USD 1.00",
		).Error())
	})
	s.NoError(refreshDBErr)
}

func TestWalletWithdrawTestSuite(t *testing.T) {
	suite.Run(t, new(WalletWithdrawTestSuite))
}

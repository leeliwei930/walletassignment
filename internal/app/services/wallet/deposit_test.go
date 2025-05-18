package wallet_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent"
	pkgapp "github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/suite"
)

type WalletDepositTestSuite struct {
	suite.Suite
	app interfaces.Application
}

func (s *WalletDepositTestSuite) SetupTest() {
	app, err := pkgapp.InitializeFromEnv()
	s.NoError(err)

	s.app = app
}

func (s *WalletDepositTestSuite) TestWalletDeposit_ShouldDepositCorrectAmountToUserWallet() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()
		depositResponse, err := walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: userRec.ID,
			Amount: 10000,
		})
		s.NoError(err)

		wallet := depositResponse.Wallet
		transaction := depositResponse.Transaction

		s.Equal(10000, wallet.Balance)
		s.Equal("USD", wallet.Currency)
		s.Equal("USD 100.00", wallet.FormattedBalance)

		s.Equal(10000, transaction.Amount)
		s.Equal("deposit", transaction.Type)
		s.Empty(transaction.RecipientReferenceNote)

	})
	s.NoError(refreshDBErr)
}

func (s *WalletDepositTestSuite) TestWalletDeposit_ShouldReturnError_WhenDepositAmountIsBelowRequiredMinimumAmount() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		locale := s.app.GetLocale()
		ut := locale.GetUT().GetFallback()

		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()
		depositResponse, err := walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: userRec.ID,
			Amount: 1,
		})
		s.Nil(depositResponse)
		s.EqualError(err, errors.MinimumDepositAmountRequiredErr(
			ut,
			"USD 1.00",
		).Error())
	})
	s.NoError(refreshDBErr)
}

func (s *WalletDepositTestSuite) TestWalletDeposit_ShouldReturnNotFoundError_WhenGivenUserIDIsInvalid() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()

		walletSvc := s.app.GetWalletService()
		depositResponse, err := walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Amount: 1,
		})
		s.True(ent.IsNotFound(err))
		s.Nil(depositResponse)
	})
	s.NoError(refreshDBErr)
}

func TestWalletDepositTestSuite(t *testing.T) {
	suite.Run(t, new(WalletDepositTestSuite))
}

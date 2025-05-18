package wallet_test

import (
	"context"
	"testing"

	pkgapp "github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/suite"
)

type WalletTransferTestSuite struct {
	suite.Suite
	app interfaces.Application
}

func (s *WalletTransferTestSuite) SetupTest() {
	app, err := pkgapp.InitializeFromEnv()
	s.NoError(err)

	s.app = app
}

func (s *WalletTransferTestSuite) TestWalletTransfer_ShouldTransferCorrectAmountBetweenWallets() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()

		// Setup source user
		sourceUser, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		// Setup destination user
		destUser, err := userSvc.SetupUser(ctx, "9876543210", "Jane", "Smith")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()

		// Initial deposit to source wallet
		_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: sourceUser.ID,
			Amount: 10000,
		})
		s.NoError(err)

		// Perform transfer
		transferResponse, err := walletSvc.Transfer(ctx, models.WalletTransferParams{
			SenderUserID:           sourceUser.ID,
			RecipientUserID:        destUser.ID,
			Amount:                 5000,
			RecipientReferenceNote: "Test transfer",
		})
		s.NoError(err)

		// Verify source wallet
		s.Equal(5000, transferResponse.Wallet.Balance)
		s.Equal("USD 50.00", transferResponse.Wallet.FormattedBalance)
		s.Equal(5000, transferResponse.Transaction.Amount)
		s.Equal("transfer_out", transferResponse.Transaction.Type)
		s.Equal("Test transfer", *transferResponse.Transaction.RecipientReferenceNote)

		// Verify destination wallet

		destWallet, err := walletSvc.Status(ctx, models.WalletStatusParams{
			UserID: destUser.ID,
		})
		s.NoError(err)
		s.Equal(5000, destWallet.Balance)
		s.Equal("USD 50.00", destWallet.FormattedBalance)
	})
	s.NoError(refreshDBErr)
}

func (s *WalletTransferTestSuite) TestWalletTransfer_ShouldReturnError_WhenInsufficientBalanceOnSenderWallet() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()

		// Setup source user
		sourceUser, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		// Setup destination user
		destUser, err := userSvc.SetupUser(ctx, "9876543210", "Jane", "Smith")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()

		// Initial deposit to source wallet
		_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: sourceUser.ID,
			Amount: 10000,
		})
		s.NoError(err)

		// Attempt transfer larger than balance
		transferResponse, err := walletSvc.Transfer(ctx, models.WalletTransferParams{
			SenderUserID:    sourceUser.ID,
			RecipientUserID: destUser.ID,
			Amount:          15000,
		})
		s.Nil(transferResponse)
		s.EqualError(err, errors.InsufficientBalanceTransferErr.Error())
	})
	s.NoError(refreshDBErr)
}

func (s *WalletTransferTestSuite) TestWalletTransfer_ShouldReturnError_WhenTransferAmountIsBelowRequiredMinimumAmount() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		locale := s.app.GetLocale()
		ut := locale.GetUT().GetFallback()

		userSvc := s.app.GetUserService()

		// Setup source user
		sourceUser, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		// Setup destination user
		destUser, err := userSvc.SetupUser(ctx, "9876543210", "Jane", "Smith")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()

		// Initial deposit to source wallet
		_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: sourceUser.ID,
			Amount: 10000,
		})
		s.NoError(err)

		// Attempt transfer below minimum amount
		transferResponse, err := walletSvc.Transfer(ctx, models.WalletTransferParams{
			SenderUserID:    sourceUser.ID,
			RecipientUserID: destUser.ID,
			Amount:          50,
		})
		s.Nil(transferResponse)
		s.EqualError(err, errors.MinimumTransferAmountRequiredErr(
			ut,
			"USD 1.00",
		).Error())
	})
	s.NoError(refreshDBErr)
}

func (s *WalletTransferTestSuite) TestWalletTransfer_ShouldReturnError_WhenSourceAndDestinationAreIdentical() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()

		// Setup user
		user, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()

		// Initial deposit
		_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: user.ID,
			Amount: 10000,
		})
		s.NoError(err)

		// Attempt transfer to self
		transferResponse, err := walletSvc.Transfer(ctx, models.WalletTransferParams{
			SenderUserID:    user.ID,
			RecipientUserID: user.ID,
			Amount:          1000,
		})
		s.Nil(transferResponse)
		s.EqualError(err, errors.IdenticalSourceAndDestinationTransferErr.Error())
	})
	s.NoError(refreshDBErr)
}

func TestWalletTransferTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTransferTestSuite))
}

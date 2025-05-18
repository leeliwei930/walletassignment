package wallet_test

import (
	"context"
	"testing"

	pkgapp "github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/suite"
)

type WalletTransactionsTestSuite struct {
	suite.Suite
	app interfaces.Application
}

func (s *WalletTransactionsTestSuite) SetupTest() {
	app, err := pkgapp.InitializeFromEnv()
	s.NoError(err)

	s.app = app
}

func (s *WalletTransactionsTestSuite) TestWalletTransactions_ShouldReturnCorrectTransactions() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		walletSvc := s.app.GetWalletService()

		// Setup user
		user, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		// Initial deposit
		_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
			UserID: user.ID,
			Amount: 10000,
		})
		s.NoError(err)

		// Withdraw some amount
		_, err = walletSvc.Withdraw(ctx, models.WalletWithdrawalParams{
			UserID: user.ID,
			Amount: 2000,
		})
		s.NoError(err)

		// Get transactions
		transactionsResponse, err := walletSvc.Transactions(ctx, models.WalletTransactionsParams{
			UserID: user.ID,
			Page:   1,
			Limit:  10,
		})
		s.NoError(err)

		pagination := transactionsResponse.Pagination
		transactions := transactionsResponse.Transactions

		// Verify transactions
		s.Len(transactions, 2)
		s.Equal(1, pagination.CurrentPage)
		s.Equal(2, pagination.ItemsPerPage)
		s.Equal(false, pagination.HasNext)
		s.Equal(false, pagination.HasPrev)
		s.Equal(2, pagination.TotalItems)

		// Verify first transaction (withdrawal)
		s.Equal(2000, transactions[0].Amount)
		s.Equal("withdrawal", transactions[0].Type)

		// Verify second transaction (deposit)
		s.Equal(10000, transactions[1].Amount)
		s.Equal("deposit", transactions[1].Type)
	})
	s.NoError(refreshDBErr)
}

func (s *WalletTransactionsTestSuite) TestWalletTransactions_ShouldReturnEmptyList_WhenNoTransactions() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		walletSvc := s.app.GetWalletService()

		// Setup user
		user, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		// Get transactions
		transactionsResponse, err := walletSvc.Transactions(ctx, models.WalletTransactionsParams{
			UserID: user.ID,
			Page:   1,
			Limit:  10,
		})
		s.NoError(err)

		pagination := transactionsResponse.Pagination
		transactions := transactionsResponse.Transactions
		// Verify empty transactions
		s.Len(transactions, 0)
		s.Equal(0, pagination.TotalItems)
		s.Equal(0, pagination.ItemsPerPage)
		s.Equal(false, pagination.HasNext)
		s.Equal(false, pagination.HasPrev)
		s.Equal(0, pagination.TotalPages)
		s.Equal(1, pagination.CurrentPage)
	})
	s.NoError(refreshDBErr)
}

func (s *WalletTransactionsTestSuite) TestWalletTransactions_ShouldRespectPagination() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		walletSvc := s.app.GetWalletService()

		// Setup user
		user, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		// Create multiple transactions
		for i := 0; i < 15; i++ {
			_, err = walletSvc.Deposit(ctx, models.WalletDepositParams{
				UserID: user.ID,
				Amount: 1000,
			})
			s.NoError(err)
		}

		// Get first page
		transactionsResponse, err := walletSvc.Transactions(ctx, models.WalletTransactionsParams{
			UserID: user.ID,
			Page:   1,
			Limit:  10,
		})
		s.NoError(err)

		pagination := transactionsResponse.Pagination
		transactions := transactionsResponse.Transactions

		// Verify first page
		s.Len(transactions, 10)
		s.Equal(15, pagination.TotalItems)
		s.Equal(10, pagination.ItemsPerPage)
		s.Equal(1, pagination.CurrentPage)
		s.Equal(true, pagination.HasNext)
		s.Equal(false, pagination.HasPrev)

		// Get second page
		transactionsResponse, err = walletSvc.Transactions(ctx, models.WalletTransactionsParams{
			UserID: user.ID,
			Page:   2,
			Limit:  10,
		})
		s.NoError(err)

		pagination = transactionsResponse.Pagination
		transactions = transactionsResponse.Transactions

		// Verify second page
		s.Len(transactions, 5)
		s.Equal(15, pagination.TotalItems)
		s.Equal(5, pagination.ItemsPerPage)
		s.Equal(2, pagination.CurrentPage)
		s.Equal(false, pagination.HasNext)
		s.Equal(true, pagination.HasPrev)
	})
	s.NoError(refreshDBErr)
}

func TestWalletTransactionsTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTransactionsTestSuite))
}

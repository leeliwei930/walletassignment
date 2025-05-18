package wallet_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent"
	pkgapp "github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/suite"
)

type WalletStatusTestSuite struct {
	suite.Suite
	app interfaces.Application
}

func (s *WalletStatusTestSuite) SetupTest() {
	app, err := pkgapp.InitializeFromEnv()
	if err != nil {
		s.T().Fatal(err)
	}
	s.app = app
}

func (s *WalletStatusTestSuite) TestWalletStatus_ShouldReturnCorrectWalletBalance() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		walletSvc := s.app.GetWalletService()
		walletRec, err := walletSvc.Status(ctx, models.WalletStatusParams{
			UserID: userRec.ID,
		})
		s.NoError(err)

		s.Equal(walletRec.Balance, 0)
		s.Equal(walletRec.Currency, "USD")
		s.Equal(walletRec.FormattedBalance, "USD 0.00")
		s.NoError(err)
	})
	s.NoError(refreshDBErr)
}

func (s *WalletStatusTestSuite) TestWalletStatus_ShouldReturnError_WhenUserIDIsInvalid() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()

		walletSvc := s.app.GetWalletService()
		walletRec, err := walletSvc.Status(ctx, models.WalletStatusParams{
			UserID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		})
		s.True(ent.IsNotFound(err))
		s.Nil(walletRec)

	})
	s.NoError(refreshDBErr)
}

func TestWalletStatusTestSuite(t *testing.T) {
	suite.Run(t, new(WalletStatusTestSuite))
}

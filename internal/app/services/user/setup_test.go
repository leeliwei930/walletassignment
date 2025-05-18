package user_test

import (
	"context"
	"testing"

	"github.com/leeliwei930/walletassignment/ent"
	pkgapp "github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/suite"
)

type SetupUserTestSuite struct {
	suite.Suite
	app interfaces.Application
}

func (s *SetupUserTestSuite) SetupTest() {
	app, err := pkgapp.InitializeFromEnv()
	s.NoError(err)

	s.app = app
}

func (s *SetupUserTestSuite) TestSetupUser_ShouldCreateUser_WhenValidInput() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()

		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)
		s.NotNil(userRec)
		s.Equal("0123456789", userRec.PhoneNumber)
		s.Equal("John", userRec.FirstName)
		s.Equal("Doe", userRec.LastName)
	})
	s.NoError(refreshDBErr)
}

func (s *SetupUserTestSuite) TestSetupUser_ShouldReturnError_WhenPhoneNumberAlreadyExists() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()

		// First user creation
		_, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		// Second user with same phone number
		_, err = userSvc.SetupUser(ctx, "0123456789", "Jane", "Smith")
		s.True(ent.IsConstraintError(err))
	})
	s.NoError(refreshDBErr)
}

func TestSetupUserTestSuite(t *testing.T) {
	suite.Run(t, new(SetupUserTestSuite))
}

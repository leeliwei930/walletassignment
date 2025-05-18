package user_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent"
	pkgapp "github.com/leeliwei930/walletassignment/internal/app"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/suite"
)

type GetUserIDByPhoneTestSuite struct {
	suite.Suite
	app interfaces.Application
}

func (s *GetUserIDByPhoneTestSuite) SetupTest() {
	app, err := pkgapp.InitializeFromEnv()
	s.NoError(err)

	s.app = app
}

func (s *GetUserIDByPhoneTestSuite) TestGetUserIDByPhone_ShouldReturnUserID_WhenUserExists() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		foundUserID, err := userSvc.GetUserIDByPhone(ctx, userRec.PhoneNumber)
		s.NoError(err)
		s.Equal(userRec.ID, foundUserID)
	})
	s.NoError(refreshDBErr)
}

func (s *GetUserIDByPhoneTestSuite) TestGetUserIDByPhone_ShouldReturnError_WhenUserDoesNotExist() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()

		foundUserID, err := userSvc.GetUserIDByPhone(ctx, "9999999999")
		s.True(ent.IsNotFound(err))
		s.Equal(uuid.Nil, foundUserID)
	})
	s.NoError(refreshDBErr)
}

func TestGetUserIDByPhoneTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserIDByPhoneTestSuite))
}

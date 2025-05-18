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

type GetUserByIDTestSuite struct {
	suite.Suite
	app interfaces.Application
}

func (s *GetUserByIDTestSuite) SetupTest() {
	app, err := pkgapp.InitializeFromEnv()
	s.NoError(err)

	s.app = app
}

func (s *GetUserByIDTestSuite) TestGetUserByID_ShouldReturnUser_WhenUserExists() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()
		userRec, err := userSvc.SetupUser(ctx, "0123456789", "John", "Doe")
		s.NoError(err)

		foundUser, err := userSvc.GetUserByID(ctx, userRec.ID)
		s.NoError(err)
		s.NotNil(foundUser)
		s.Equal(userRec.ID, foundUser.ID)
		s.Equal(userRec.PhoneNumber, foundUser.PhoneNumber)
		s.Equal(userRec.FirstName, foundUser.FirstName)
		s.Equal(userRec.LastName, foundUser.LastName)
	})
	s.NoError(refreshDBErr)
}

func (s *GetUserByIDTestSuite) TestGetUserByID_ShouldReturnError_WhenUserDoesNotExist() {
	refreshDBErr := s.app.UseRefreshDB(s.T(), func() {
		ctx := context.Background()
		userSvc := s.app.GetUserService()

		nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
		foundUser, err := userSvc.GetUserByID(ctx, nonExistentID)
		s.True(ent.IsNotFound(err))
		s.Nil(foundUser)
	})
	s.NoError(refreshDBErr)
}

func TestGetUserByIDTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserByIDTestSuite))
}

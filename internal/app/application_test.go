package app_test

import (
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/database"
	"github.com/leeliwei930/walletassignment/internal/app"
	_ "github.com/leeliwei930/walletassignment/tests"
	"github.com/stretchr/testify/suite"
)

type ApplicationTestSuite struct {
	suite.Suite
}

func (ats *ApplicationTestSuite) TestInitializeFromEnv() {

	err := godotenv.Load(".env.testing")
	ats.NoError(err)

	_app, err := app.InitializeFromEnv()
	ats.NoError(err)

	_config := _app.GetConfig()
	ats.NotNil(_config)

	ats.Equal(&config.Config{
		DBConfig: &config.DBConnectionConfig{
			Connection: &database.PostgresDBConnection{
				Host:         "127.0.0.1",
				Port:         5432,
				Username:     "wallet_app_db_test",
				Password:     "wallet_app_db_test",
				DatabaseName: "wallet_app_db_test",
			},
			MaxIdleConnections:    10,
			MaxOpenConns:          30,
			ConnectionMaxLifeTime: 10 * time.Minute,
		},
		DevDBConfig: &config.DBConnectionConfig{
			Connection: &database.PostgresDBConnection{
				Host:         "127.0.0.1",
				Username:     "wallet_app_db_test",
				DatabaseName: "wallet_app_db_test",
				Password:     "wallet_app_db_test",
				Port:         5432,
			},
			MaxIdleConnections:    10,
			MaxOpenConns:          30,
			ConnectionMaxLifeTime: 10 * time.Minute,
		},

		Server: &config.Server{
			Host:         "localhost",
			Port:         8009,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}, _config)

	_db := _app.GetDB()
	ats.NotNil(_db)

	_ent := _app.GetEnt()
	ats.NotNil(_ent)

	_log := _app.GetLog()
	ats.NotNil(_log)

}

func (ats *ApplicationTestSuite) TestInitialize_WithMissingConfigOnEnt() {
	app, err := app.InitializeFromEnv(
		app.WithConfig(&config.Config{
			DBConfig: nil,
		},
		),
	)
	ats.Nil(app)
	ats.EqualError(err, "DBConfig is not set in configuration")
}

func TestApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuite))
}

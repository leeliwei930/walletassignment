package app

import (
	"time"

	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/constant"
	"github.com/leeliwei930/walletassignment/database"
	"github.com/spf13/viper"
)

func (app *application) InitConfig() error {
	appPort := viper.GetInt(constant.APP_PORT)
	host := viper.GetString(constant.APP_HOST)

	_config := &config.Config{
		DevDBConfig: &config.DBConnectionConfig{
			Connection: &database.PostgresDBConnection{
				Host:         viper.GetString(constant.DEV_DB_HOST),
				Username:     viper.GetString(constant.DEV_DB_USER),
				Password:     viper.GetString(constant.DEV_DB_PASSWORD),
				DatabaseName: viper.GetString(constant.DEV_DB_NAME),
				Port:         viper.GetInt(constant.DEV_DB_PORT),
			},
			MaxIdleConnections:    10,
			MaxOpenConns:          30,
			ConnectionMaxLifeTime: 10 * time.Minute,
		},
		DBConfig: &config.DBConnectionConfig{
			Connection: &database.PostgresDBConnection{
				Host:         viper.GetString(constant.DB_HOST),
				Username:     viper.GetString(constant.DB_USER),
				Password:     viper.GetString(constant.DB_PASSWORD),
				DatabaseName: viper.GetString(constant.DB_NAME),
				Port:         viper.GetInt(constant.DB_PORT),
			},
			MaxIdleConnections:    10,
			MaxOpenConns:          30,
			ConnectionMaxLifeTime: 10 * time.Minute,
		},
		Server: &config.Server{
			Host:         host,
			Port:         appPort,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}

	app.config = _config
	return nil
}

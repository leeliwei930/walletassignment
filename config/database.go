package config

import "time"

type DBConnectionConfig struct {
	Connection            DBConnection
	MaxIdleConnections    int
	MaxOpenConns          int
	ConnectionMaxLifeTime time.Duration
}

type DBConnection interface {
	Driver() string
	DSN() string
	AtlasDSN() string
	EntDialect() string
}

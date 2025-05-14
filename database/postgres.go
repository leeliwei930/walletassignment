package database

import (
	"fmt"

	"entgo.io/ent/dialect"
)

type PostgresDBConnection struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

func (m PostgresDBConnection) Driver() string {
	return dialect.Postgres
}

func (m PostgresDBConnection) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DatabaseName,
	)
}

func (m PostgresDBConnection) AtlasDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?search_path=public&sslmode=disable",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DatabaseName,
	)

}

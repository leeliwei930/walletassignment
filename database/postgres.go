package database

import (
	"fmt"

	"entgo.io/ent/dialect"

	// For ent client to use pgx as driver
	_ "github.com/jackc/pgx/v5/stdlib"
	// For atlas migration to use lib/pq as driver
	_ "github.com/lib/pq"
)

type PostgresDBConnection struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

func (m PostgresDBConnection) Driver() string {
	return "pgx"
}

func (m PostgresDBConnection) EntDialect() string {
	return dialect.Postgres
}

func (m PostgresDBConnection) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
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

package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func Init(config *Config) (*PostgresDB, error) {
	db, err := sql.Open("postgres", config.Postgres_URI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse db config: %w", err)
	}

	if err := db.Ping(); err != nil{
		return nil, fmt.Errorf("failed to connect db: %w", err)
	}

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) GetDB() *sql.DB {
	return p.db
}

func (p *PostgresDB) Close() error {
	return p.db.Close()
}
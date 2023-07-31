package postgres

import (
	"fmt"
	"grey/config"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func NewPostgres(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres", 
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.DBName, cfg.Postgres.Password, cfg.Postgres.SSLMode),
		)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

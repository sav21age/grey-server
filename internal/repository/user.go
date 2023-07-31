package repository

import (
	"context"
	"fmt"

	"grey/internal/domain"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserInterface interface {
	CreateUser(ctx context.Context, user domain.User) error
}

func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, firstname, lastname, fullname, age, is_married, password) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", tableUser)
	
	row := r.db.QueryRow(
		query, 
		user.Username, user.Firstname, user.Lastname, user.Fullname, user.Age, user.IsMarried, user.Password,
	)
	
	err := row.Scan(&id)
	// if  err != nil {
	// 	log.Error().Err(err).Msg("")
	// 	return err
	// }

	// TODO: Fix test, he doesn't see *pq.Error
	if err, ok := err.(*pq.Error); ok {
		log.Error().Err(err).Msg("")

		if err.Code.Name() == "unique_violation" {
			return domain.ErrRecordAlreadyExists
		}

		// if err.Code == "23505" {
		// 	return domain.ErrRecordAlreadyExists
		// }

		// if err == sql.ErrNoRows {
		// }
		
		return err
	}

	return nil
}


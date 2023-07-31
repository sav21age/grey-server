package service

import (
	"context"
	"fmt"
	"grey/config"
	"grey/internal/domain"
	"grey/internal/repository"
	"grey/pkg/cipher"
)

type UserService struct {
	r   repository.UserInterface
	cfg *config.Config
}

func NewUserService(r repository.UserInterface, config *config.Config) *UserService {
	return &UserService{
		r:   r,
		cfg: config,
	}
}

//go:generate mockgen -source=user.go -destination=mock/user.go

type UserInterface interface {
	SignUp(ctx context.Context, input domain.UserSignUpInput) error
}

func (s *UserService) SignUp(ctx context.Context, input domain.UserSignUpInput) error {
	passwordHash := cipher.GeneratePassword(input.Password, s.cfg.Cipher.Salt)

	user := domain.User{
		Username:    input.Username,
		Firstname:   input.Firstname,
		Lastname:    input.Lastname,
		Fullname:    fmt.Sprintf("%s %s", input.Firstname, input.Lastname),
		Age:    	 input.Age,
		IsMarried:   input.IsMarried,
		Password:    passwordHash,
	}
	
	if err := s.r.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
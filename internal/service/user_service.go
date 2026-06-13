package service

import (
	"context"
	"errors"

	"taskflow/internal/model"
	"taskflow/internal/repository"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(ctx context.Context, name, email, password string) (*model.User, error) {
	_, err := s.userRepo.GetByEmail(ctx, email)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	if err != pgx.ErrNoRows {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

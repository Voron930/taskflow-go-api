package repository

import (
	"context"
	"time"

	"taskflow/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `
	INSERT INTO users (name, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`
	now := time.Now()

	err := r.db.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
		now,
		now,
	).Scan(&user.ID)

	return err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
	SELECT id, name, email, password_hash, created_at, updated_at
	FROM users
	WHERE email = $1
	`
	var user model.User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

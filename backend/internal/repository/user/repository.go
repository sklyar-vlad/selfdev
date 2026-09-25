package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	appErrors "github.com/sklyar-vlad/selfDev/internal/errors"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
	"go.uber.org/zap"
)

type repository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewRepository(pool *pgxpool.Pool, logger *zap.Logger) *repository {
	return &repository{pool: pool, logger: logger}
}
func (r *repository) Create(ctx context.Context, u *model.User) error {
	_, err := r.pool.Exec(ctx, `INSERT INTO users (user_id,username,email,password_hash,avatar_url) VALUES ($1,$2,$3,$4,$5)`, u.UserId, u.Username, u.Email, u.PasswordHash, u.AvatarURL)
	return err
}
func (r *repository) GetByLogin(ctx context.Context, login string) (model.User, error) {
	var u model.User
	err := r.pool.QueryRow(ctx, `SELECT user_id,username,email,password_hash,avatar_url FROM users WHERE lower(email)=lower($1) OR lower(username)=lower($1) LIMIT 1`, login).Scan(&u.UserId, &u.Username, &u.Email, &u.PasswordHash, &u.AvatarURL)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, appErrors.ErrUserNotFound
	}
	if err != nil {
		return model.User{}, fmt.Errorf("failed get user: %w", err)
	}
	return u, nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	var u model.User
	err := r.pool.QueryRow(ctx, `SELECT user_id, username, email, password_hash, avatar_url FROM users WHERE user_id = $1`, id).Scan(&u.UserId, &u.Username, &u.Email, &u.PasswordHash, &u.AvatarURL)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, appErrors.ErrUserNotFound
	}
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

package auth

import (
	"context"
	"github.com/google/uuid"
	appErrors "github.com/sklyar-vlad/selfDev/internal/errors"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserService interface {
	GetUserByLogin(context.Context, string) (model.User, error)
	CreateUser(context.Context, model.User) (model.User, error)
}
type AuthRepository interface {
	CreateSession(context.Context, string, uuid.UUID) error
	DeleteSession(context.Context, string) error
}
type Service struct {
	users    UserService
	sessions AuthRepository
	logger   *zap.Logger
}

func NewService(u UserService, r AuthRepository, l *zap.Logger) *Service {
	return &Service{users: u, sessions: r, logger: l}
}
func (s *Service) Register(ctx context.Context, username, email, password, avatar string) (model.User, string, error) {
	username = strings.TrimSpace(username)
	email = strings.ToLower(strings.TrimSpace(email))
	if len(username) < 2 || len(username) > 30 || !strings.Contains(email, "@") {
		return model.User{}, "", appErrors.ErrInvalidEmail
	}
	if len(password) < 8 {
		return model.User{}, "", appErrors.ErrInvalidPassword
	}
	if _, err := s.users.GetUserByLogin(ctx, email); err == nil {
		return model.User{}, "", appErrors.ErrEmailAlreadyExists
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, "", err
	}
	u := model.NewUser(username, email, string(hash), avatar)
	if _, err = s.users.CreateUser(ctx, u); err != nil {
		return model.User{}, "", err
	}
	sid, err := s.newSession(ctx, u.UserId)
	return u, sid, err
}
func (s *Service) Login(ctx context.Context, login, password string) (model.User, string, error) {
	u, err := s.users.GetUserByLogin(ctx, strings.TrimSpace(login))
	if err != nil || bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return model.User{}, "", appErrors.ErrUnauthorized
	}
	sid, err := s.newSession(ctx, u.UserId)
	return u, sid, err
}
func (s *Service) Logout(ctx context.Context, sid string) error {
	return s.sessions.DeleteSession(ctx, sid)
}
func (s *Service) newSession(ctx context.Context, id uuid.UUID) (string, error) {
	sid := uuid.NewString()
	return sid, s.sessions.CreateSession(ctx, sid, id)
}
func PublicUser(u model.User) map[string]any {
	return map[string]any{"user_id": u.UserId, "username": u.Username, "email": u.Email, "avatar_url": u.AvatarURL}
}

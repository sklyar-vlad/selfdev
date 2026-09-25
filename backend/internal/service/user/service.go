package user

import (
	"context"
	"github.com/google/uuid"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
	"go.uber.org/zap"
)

type Repository interface {
	Create(context.Context, *model.User) error
	GetByLogin(context.Context, string) (model.User, error)
	GetByID(context.Context, uuid.UUID) (model.User, error)
}
type Service struct {
	repo   Repository
	logger *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}
func (s *Service) CreateUser(ctx context.Context, u model.User) (model.User, error) {
	return u, s.repo.Create(ctx, &u)
}
func (s *Service) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *Service) GetUserByLogin(ctx context.Context, l string) (model.User, error) {
	return s.repo.GetByLogin(ctx, l)
}

package user

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
	"github.com/sklyar-vlad/selfDev/middleware"
	"go.uber.org/zap"
	"net/http"
)

type Service interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error)
}
type handler struct {
	service Service
	logger  *zap.Logger
}

func NewHandler(s Service, l *zap.Logger) *handler { return &handler{service: s, logger: l} }
func (h *handler) GetCurrent(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(middleware.UserIDKey{}).(uuid.UUID)
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}
	u, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", 404)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"user_id": u.UserId, "username": u.Username, "email": u.Email, "avatar_url": u.AvatarURL})
}

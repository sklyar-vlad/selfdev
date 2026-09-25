package auth

import (
	"context"
	"encoding/json"
	model "github.com/sklyar-vlad/selfDev/internal/model/user"
	svc "github.com/sklyar-vlad/selfDev/internal/service/auth"
	"go.uber.org/zap"
	"net/http"
)

type Service interface {
	Register(context.Context, string, string, string, string) (model.User, string, error)
	Login(context.Context, string, string) (model.User, string, error)
	Logout(context.Context, string) error
}
type handler struct {
	service      Service
	cookieDomain string
	cookieSecure bool
	logger       *zap.Logger
}

func NewHandler(s Service, domain string, secure bool, l *zap.Logger) *handler {
	return &handler{service: s, cookieDomain: domain, cookieSecure: secure, logger: l}
}

type credentials struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var in credentials
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "invalid json", 400)
		return
	}
	u, sid, err := h.service.Register(r.Context(), in.Username, in.Email, in.Password, in.AvatarURL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	h.setCookie(w, sid)
	json.NewEncoder(w).Encode(svc.PublicUser(u))
}
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var in credentials
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "invalid json", 400)
		return
	}
	u, sid, err := h.service.Login(r.Context(), in.Login, in.Password)
	if err != nil {
		http.Error(w, "invalid credentials", 401)
		return
	}
	h.setCookie(w, sid)
	json.NewEncoder(w).Encode(svc.PublicUser(u))
}
func (h *handler) Logout(w http.ResponseWriter, r *http.Request) {
	if c, e := r.Cookie("session"); e == nil {
		_ = h.service.Logout(r.Context(), c.Value)
	}
	http.SetCookie(w, &http.Cookie{Name: "session", Path: "/", MaxAge: -1})
	w.WriteHeader(http.StatusNoContent)
}
func (h *handler) setCookie(w http.ResponseWriter, sid string) {
	http.SetCookie(w, &http.Cookie{Name: "session", Value: sid, Path: "/", Domain: h.cookieDomain, HttpOnly: true, Secure: h.cookieSecure, SameSite: http.SameSiteLaxMode, MaxAge: 30 * 24 * 3600})
}

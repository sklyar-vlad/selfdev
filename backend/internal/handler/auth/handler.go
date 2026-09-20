package auth

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type AuthService interface {
	Login(ctx context.Context, code string) (string, error)
}

type handler struct {
	service      AuthService
	redirectURI  string
	cookieDomain string
	cookieSecure bool
	logger       *zap.Logger
}

func NewHandler(
	service AuthService,
	redirectURI, cookieDomain string,
	cookieSecure bool,
	logger *zap.Logger,
) *handler {
	return &handler{
		service:      service,
		redirectURI:  redirectURI,
		cookieDomain: cookieDomain,
		cookieSecure: cookieSecure,
		logger:       logger,
	}
}

func (h *handler) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	session, err := h.service.Login(r.Context(), code)
	if err != nil {
		h.logger.Error("failed login", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session,
		Path:     "/",
		Domain:   h.cookieDomain,
		HttpOnly: true,
		Secure:   h.cookieSecure,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   30 * 24 * 3600,
	})

	http.Redirect(w, r, h.redirectURI+"/me/profile", http.StatusFound)
}

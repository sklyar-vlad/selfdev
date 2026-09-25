package handler

import (
	"net/http"
)

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

// TODO: GetUsers(w http.ResponseWriter, r *http.Request)
// TODO: DeleteUser(w http.ResponseWriter, r *http.Request)
// TODO: UpdateUser(w http.ResponseWriter, r *http.Request)
type UserHandler interface {
	GetCurrent(w http.ResponseWriter, r *http.Request)
	// CreateUser(w http.ResponseWriter, r *http.Request)
	// GetUser(w http.ResponseWriter, r *http.Request)
}

// TODO: UpdateHabit(w http.ResponseWriter, r *http.Request)
type HabitHandler interface {
	GetHabits(w http.ResponseWriter, r *http.Request)
	CreateHabit(w http.ResponseWriter, r *http.Request)
	UpdateHabit(w http.ResponseWriter, r *http.Request)
	DeleteHabit(w http.ResponseWriter, r *http.Request)
	ConfirmHabit(w http.ResponseWriter, r *http.Request)
	CancelHabit(w http.ResponseWriter, r *http.Request)
	GetHabitConfirmDates(w http.ResponseWriter, r *http.Request)
}

func RegisterPublicRoutes(mux *http.ServeMux, authHandler AuthHandler) {
	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)
	mux.HandleFunc("POST /auth/logout", authHandler.Logout)
}

// TODO: mux.HandleFunc("GET /api/users", userHandler.GetUsers)
// TODO: mux.HandleFunc("PATCH /api/users/{id}", userHandler.UpdateUser)
// TODO: mux.HandleFunc("DELETE /api/users/{id}", userHandler.DeleteUser)
// TODO: mux.HandleFunc("PATCH /api/habit/{id}", habitHandler.UpdateHabit)

// mux.HandleFunc("POST /api/users", userHandler.CreateUser)
// mux.HandleFunc("GET /api/users/{id}", userHandler.GetUser)

func RegisterProtectedRoutes(mux *http.ServeMux, userHandler UserHandler, habitHandler HabitHandler) {
	mux.HandleFunc("GET /api/me", userHandler.GetCurrent)
	mux.HandleFunc("GET /api/habits", habitHandler.GetHabits)
	mux.HandleFunc("POST /api/habit", habitHandler.CreateHabit)
	mux.HandleFunc("PUT /api/habit/{id}", habitHandler.UpdateHabit)
	mux.HandleFunc("DELETE /api/habit/{id}", habitHandler.DeleteHabit)
	mux.HandleFunc("POST /api/habit/{id}/confirm", habitHandler.ConfirmHabit)
	mux.HandleFunc("DELETE /api/habit/{id}/confirm", habitHandler.CancelHabit)
	mux.HandleFunc("GET /api/habit/{id}/confirm", habitHandler.GetHabitConfirmDates)
}

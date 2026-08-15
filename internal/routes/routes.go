package routes

import (
	"database/sql"
	"net/http"

	"gradivus-core-backend/internal/config"
	"gradivus-core-backend/internal/handlers"
	"gradivus-core-backend/internal/middleware"
)

func New(db *sql.DB, cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	authHandler := handlers.NewAuthHandler(db, cfg)
	userHandler := handlers.NewUserHandler(db)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// Auth — dipakai semua modul lain buat login & validasi token (SSO)
	mux.HandleFunc("POST /api/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/auth/login", authHandler.Login)
	mux.HandleFunc("GET /api/auth/me", authHandler.Me)

	// Users
	mux.HandleFunc("GET /api/users", userHandler.List)
	mux.HandleFunc("GET /api/users/{id}", userHandler.Get)

	return middleware.CORS(mux)
}

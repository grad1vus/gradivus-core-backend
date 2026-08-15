package handlers

import (
	"database/sql"
	"net/http"

	"gradivus-core-backend/internal/models"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GET /api/users
// List semua user dalam 1 organization (buat admin/HR lihat daftar karyawan)
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	writeJSON(w, http.StatusOK, users)
}

// GET /api/users/{id}
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, models.User{})
}

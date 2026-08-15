package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"gradivus-core-backend/internal/config"
	"gradivus-core-backend/internal/middleware"
	"gradivus-core-backend/internal/models"
)

type AuthHandler struct {
	DB  *sql.DB
	Cfg *config.Config
}

func NewAuthHandler(db *sql.DB, cfg *config.Config) *AuthHandler {
	return &AuthHandler{DB: db, Cfg: cfg}
}

type registerInput struct {
	OrganizationName string `json:"organization_name"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
}

// POST /api/auth/register
// Dipakai saat bikin akun baru + workspace/organization baru (misal client SaaS baru daftar).
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input registerInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to process password")
		return
	}

	// TODO: 1) insert row ke tabel organizations
	//       2) insert row ke tabel users dengan password_hash = hash, role = "owner"
	//       3) generate JWT dan kembalikan ke client
	_ = hash

	writeJSON(w, http.StatusCreated, map[string]string{
		"message": "register endpoint — belum terhubung ke database, isi query di sini",
	})
}

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST /api/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input loginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// TODO: 1) query user by email
	//       2) bcrypt.CompareHashAndPassword
	//       3) generate token pakai middleware.GenerateToken(h.Cfg.JWTSecret, user.ID, user.OrganizationID, user.Role)

	token, _ := middleware.GenerateToken(h.Cfg.JWTSecret, "dummy-user-id", "dummy-org-id", "member")

	writeJSON(w, http.StatusOK, map[string]string{
		"token":   token,
		"message": "login endpoint — belum terhubung ke database, isi query di sini",
	})
}

// GET /api/auth/me
// Dipanggil modul lain (Board, HR, dll) buat validasi token & ambil identitas user.
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	tokenString := middleware.ExtractBearerToken(r)
	if tokenString == "" {
		writeError(w, http.StatusUnauthorized, "missing token")
		return
	}

	claims, err := middleware.ParseToken(h.Cfg.JWTSecret, tokenString)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid or expired token")
		return
	}

	writeJSON(w, http.StatusOK, models.User{
		ID:             claims.UserID,
		OrganizationID: claims.OrganizationID,
		Role:           claims.Role,
	})
}

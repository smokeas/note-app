package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"note-app/internal/auth"
	"note-app/internal/service"
)

// AuthHandler отвечает за регистрацию и логин
type AuthHandler struct {
	userService *service.UserService
	jwtSecret   string
	jwtTTL      time.Duration
}

func NewAuthHandler(us *service.UserService, jwtSecret string, jwtTTL time.Duration) *AuthHandler {
	return &AuthHandler{userService: us, jwtSecret: jwtSecret, jwtTTL: jwtTTL}
}

type registerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "use POST", http.StatusMethodNotAllowed)
		return
	}

	var req registerReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		if err == service.ErrEmailTaken {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{"id": user.ID, "email": user.Email})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "use POST", http.StatusMethodNotAllowed)
		return
	}

	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// сгенерировать JWT (если у тебя пакет internal/auth предоставляет GenerateToken)
	token, err := auth.GenerateToken(user.ID, h.jwtSecret, h.jwtTTL)
	if err != nil {
		http.Error(w, "cannot create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
		"user": map[string]any{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

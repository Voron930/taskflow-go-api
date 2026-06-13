package handler

import (
	"encoding/json"
	"net/http"
	"taskflow/internal/middleware"
)

type ProfileResponse struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

func Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int64)
	email := r.Context().Value(middleware.EmailKey).(string)

	resp := ProfileResponse{
		UserID: userID,
		Email:  email,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

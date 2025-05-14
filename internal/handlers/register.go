package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email string `json:"email"`
	Pswd  string `json:"pswd"`
}

var DB *pgxpool.Pool

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Pswd == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	hashed, errs := bcrypt.GenerateFromPassword([]byte(req.Pswd), bcrypt.DefaultCost)
	if errs != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users (email, password VALUES ($1, $2)`
	_, err = DB.Exec(context.Background(), query, req.Email, hashed)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully!",
	})
}

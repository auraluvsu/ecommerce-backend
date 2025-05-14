package services

import (
	"github.com/auraluvsu/internal/handlers"
	"net/http"
)

func Router() {
	http.HandleFunc("/register", handlers.RegisterHandler)
}

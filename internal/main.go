package main

import (
	"log"
	"net/http"

	"github.com/auraluvsu/internal/db"
	"github.com/auraluvsu/internal/services"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	services.Router()
	log.Println("Server running on http://localhost:5432")
	if err := http.ListenAndServe(":5432", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

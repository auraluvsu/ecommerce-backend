package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://admin:admin123@localhost:5432/ecommerce"
	}
	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v", err)
	}
	log.Println("Connected to Postgres!")
	return nil
}

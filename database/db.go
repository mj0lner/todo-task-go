package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var DB *pgxpool.Pool

func InitDatabase() {
	var err error
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Создание таблицы задач
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id          SERIAL PRIMARY KEY,
		title       TEXT NOT NULL,
		description TEXT,
		status      TEXT NOT NULL,
		created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	_, err = DB.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create tasks table: %v", err)
	}

	fmt.Println("Successfully connected and initialized database")
}

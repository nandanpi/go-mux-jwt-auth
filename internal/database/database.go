package database

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DbServices interface {
	CreateAccount(username, hashedPassword string) error
	GetUser(w http.ResponseWriter, username string) *User
}

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, err

}

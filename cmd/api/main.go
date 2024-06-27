package main

import (
	"log"

	"github.com/nandanpi/go-mux-jwt-auth/internal/database"
	"github.com/nandanpi/go-mux-jwt-auth/internal/server"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	err = db.SchemaPush()
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(":1234", db)
	server.Run()

}

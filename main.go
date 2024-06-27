package main

import (
	"log"
)

func main() {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	err = db.SchemaPush()
	if err != nil {
		log.Fatal(err)
	}
	server := NewServer(":1234", db)
	server.Run()

}

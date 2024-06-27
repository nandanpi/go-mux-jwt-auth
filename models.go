package main

import "log"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d *DB) SchemaPush() error {
	userTableSchema := `
		create table if not exists Users (
			id serial primary key,
			username varchar(50) not null unique,
			password varchar(255) not null
    )
`
	_, err := d.db.Exec(userTableSchema)
	if err != nil {
		log.Println("Schema push failed")
		return err
	}
	log.Println("Schema updated")
	return nil
}

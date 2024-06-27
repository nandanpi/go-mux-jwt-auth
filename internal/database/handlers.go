package database

import (
	"net/http"
)

func (d *DB) CreateAccount(username, hashedPassword string) error {
	_, err := d.db.Exec(`
	insert into Users (username, password)
	values ($1, $2)
	`, username, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetUser(w http.ResponseWriter, username string) *User {
	dbUser := &User{}
	_ = d.db.QueryRow(`
		select * from Users where username=$1
	`, username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	return dbUser

}

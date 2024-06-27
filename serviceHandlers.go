package main

import (
	"net/http"
)

func (d *DB) createAccount(username, hashedPassword string) error {
	_, err := d.db.Exec(`
	insert into Users (username, password)
	values ($1, $2)
	`, username, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (d *DB) getUser(w http.ResponseWriter, username string) *User {
	// dbUser := &User{}
	// _ = d.db.QueryRow(`
	// 	select * from Users where username=$1
	// `, username).Scan(dbUser.ID, dbUser.Username, dbUser.Password)
	dbUser := &User{}
	_ = d.db.QueryRow(`
		select * from Users where username=$1
	`, username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	return dbUser

}

func (d *DB) getAllUsers(w http.ResponseWriter) ([]*User, error) {
	rows, err := d.db.Query(`
		select * from users
	`)
	if err != nil {
		return nil, err
	}
	users := []*User{}
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

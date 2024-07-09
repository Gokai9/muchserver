package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthDB struct {
	db *sql.DB
}

const createauth string = `
  CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY,
  username TEXT,
  password TEXT
  );`

func OpenAuthDb() (*AuthDB, error) {
	db, err := sql.Open("sqlite3", "./auth.db")
	if err != nil {
		return nil, err
	}
	if _, err = db.Exec(createauth); err != nil {
		return nil, err
	}

	return &AuthDB{db: db}, nil
}

func (a *AuthDB) AddUser(username, password string) (int, error) {
	result, err := a.db.Exec("INSERT INTO users VALUES(NULL,?,?);", username,
		password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (a *AuthDB) GetAllUser() ([]User, error) {
	rows, err := a.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (a *AuthDB) GetById(id int) User {
	row := a.db.QueryRow("SELECT * FROM users WHERE id = ?;", id)
	var user User
	row.Scan(&user.Id, &user.Username, &user.Password)
	return user
}

func (a *AuthDB) UpdateById(upuser User) User {
	row := a.db.QueryRow("UPDATE anime SET username=?, password=? WHERE id = ?;", upuser.Username, upuser.Password, upuser.Id)
	var user User
	row.Scan(&user.Id, &user.Username, &user.Password)
	return user
}

func (a *AuthDB) DeleteById(id int) {
	a.db.Exec("DELETE FROM users WHERE id = ?;", id)
}

package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// id title eps
type Anime struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Episode int    `json:"episode"`
}
type AnimeDetail struct {
	Title   string `json:"title"`
	Episode int    `json:"episode"`
}

type AnimeDB struct {
	db *sql.DB
}

const create string = `
  CREATE TABLE IF NOT EXISTS anime (
  id INTEGER NOT NULL PRIMARY KEY,
  title TEXT,
  episode TEXT
  );`

func OpenDb() (*AnimeDB, error) {
	db, err := sql.Open("sqlite3", "./anime.db")
	if err != nil {
		return nil, err
	}
	if _, err = db.Exec(create); err != nil {
		return nil, err
	}

	return &AnimeDB{db: db}, nil
}

func (a *AnimeDB) AddAnime(title string, episode int) (int, error) {
	fmt.Println(title, episode)
	result, err := a.db.Exec("INSERT INTO anime VALUES(NULL,?,?);", title,
		episode)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (a *AnimeDB) GetAll() ([]Anime, error) {
	rows, err := a.db.Query("SELECT * FROM anime;")
	if err != nil {
		return nil, err
	}
	var animes []Anime
	for rows.Next() {
		var anime Anime
		err = rows.Scan(&anime.Id, &anime.Title, &anime.Episode)
		if err != nil {
			return nil, err
		}
		animes = append(animes, anime)
	}
	return animes, nil
}

func (a *AnimeDB) GetById(id int) Anime {
	row := a.db.QueryRow("SELECT * FROM anime WHERE id = ?;", id)
	var anime Anime
	row.Scan(&anime.Id, &anime.Title, &anime.Episode)
	return anime
}

func (a *AnimeDB) UpdateById(anime AnimeDetail, id int) Anime {
	row := a.db.QueryRow("UPDATE anime SET title=?, episode=? WHERE id = ?;", anime.Title, anime.Episode, id)
	var animes Anime
	row.Scan(&animes.Id, &animes.Title, &animes.Episode)
	return animes
}

func (a *AnimeDB) DeleteById(id int) {
	a.db.Exec("DELETE FROM anime WHERE id = ?;", id)
}

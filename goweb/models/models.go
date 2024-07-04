package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// id title eps
type Anime struct {
	id      int
	title   string
	episode int
}

type AnimeDB struct {
	db *sql.DB
}

const create string = `
  CREATE TABLE IF NOT EXISTS activities (
  id INTEGER NOT NULL PRIMARY KEY,
  title TEXT,
  episode TEXT
  );`

func OpenDb() (*AnimeDB, error) {
	db, err := sql.Open("sqlite3", "anime.db")
	if err != nil {
		return nil, err
	}
	if _, err = db.Exec(create); err != nil {
		return nil, err
	}

	return &AnimeDB{db: db}, nil
}

func (a *AnimeDB) AddAnime(anime Anime) (int, error) {
	result, err := a.db.Exec("INSERT INTO anime VALUES(NULL, ?, ?)", anime.title, anime.episode)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

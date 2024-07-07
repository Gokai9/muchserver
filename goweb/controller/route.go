package controller

import (
	"encoding/json"
	"goweb/models"
	"io"
	"log"
	"net/http"
	"regexp"
)

type AnimeDetail struct {
	Title   string
	Episode int
}

var (
	AnimeRe       = regexp.MustCompile(`^/anime/*$`)
	AnimeReWithID = regexp.MustCompile(`^/anime/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

type AnimeHandler struct{}

type animes []models.Anime

func (anime *AnimeHandler) CreateAnime(w http.ResponseWriter, r *http.Request) {
	var db, err = models.OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("There no file"))
		log.Fatal(err)
	}
	var ani AnimeDetail
	err = json.Unmarshal(b, &ani)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal(err)
	}
	_, err = db.AddAnime(ani.Title, ani.Episode)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal(err)
	}
	w.WriteHeader(200)
	w.Write([]byte("Succes"))
}
func (anime *AnimeHandler) GetAllAnime(w http.ResponseWriter, r *http.Request) {}
func (anime *AnimeHandler) GetAnime(w http.ResponseWriter, r *http.Request)    {}
func (anime *AnimeHandler) UpdateAnime(w http.ResponseWriter, r *http.Request) {}
func (anime *AnimeHandler) DeleteAnime(w http.ResponseWriter, r *http.Request) {}

func (anime *AnimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case http.MethodGet == r.Method && AnimeRe.MatchString(r.URL.Path):
		anime.GetAllAnime(w, r)
		return
	case http.MethodGet == r.Method && AnimeReWithID.MatchString(r.URL.Path):
		anime.GetAnime(w, r)
		return
	case http.MethodPost == r.Method && AnimeRe.MatchString(r.URL.Path):
		anime.CreateAnime(w, r)
		return
	case http.MethodPut == r.Method && AnimeReWithID.MatchString(r.URL.Path):
		anime.UpdateAnime(w, r)
		return
	case http.MethodDelete == r.Method && AnimeReWithID.MatchString(r.URL.Path):
		anime.DeleteAnime(w, r)
		return
	default:
		return
	}
}

type HomeHandler struct{}

func (home *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("haha"))
}

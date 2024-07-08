package controller

import (
	"encoding/json"
	"goweb/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

type AnimeHandler struct {
	*models.AnimeDB
}

func New() *AnimeHandler {
	db, err := models.OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	return &AnimeHandler{db}
}

func logError(w http.ResponseWriter, err error, status int) {
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
}

func (anime *AnimeHandler) CreateAnime(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	logError(w, err, http.StatusBadRequest)
	var ani models.AnimeDetail
	err = json.Unmarshal(b, &ani)
	logError(w, err, http.StatusBadGateway)
	_, err = anime.AddAnime(ani.Title, ani.Episode)
	logError(w, err, http.StatusBadGateway)
	w.WriteHeader(200)
	w.Write([]byte("Succes"))
}
func (anime *AnimeHandler) GetAllAnime(w http.ResponseWriter, r *http.Request) {
	ani, err := anime.GetAll()
	logError(w, err, http.StatusBadRequest)
	js, err := json.Marshal(ani)
	logError(w, err, http.StatusBadRequest)

	w.Write(js)
}
func (anime *AnimeHandler) GetAnime(w http.ResponseWriter, r *http.Request) {
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)
	ani := anime.GetById(id)
	js, err := json.Marshal(ani)
	logError(w, err, http.StatusBadRequest)

	w.Write(js)

}
func (anime *AnimeHandler) UpdateAnime(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	logError(w, err, http.StatusBadRequest)
	var ani models.AnimeDetail
	err = json.Unmarshal(b, &ani)
	logError(w, err, http.StatusBadGateway)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)
	ani := anime.UpdateById(ani, id)
	js, err := json.Marshal(ani)
	logError(w, err, http.StatusBadRequest)

	w.Write(js)
}
func (anime *AnimeHandler) DeleteAnime(w http.ResponseWriter, r *http.Request) {
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)
	anime.DeleteById(id)
	w.Write([]byte("ok"))
}

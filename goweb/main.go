package main

import (
	"goweb/controller"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := controller.New()
	mux.HandleFunc("GET /anime", server.GetAllAnime)
	mux.HandleFunc("POST /anime", server.CreateAnime)
	mux.HandleFunc("GET /anime/{id}", server.GetAnime)
	mux.HandleFunc("PUT /anime/{id}", server.UpdateAnime)
	mux.HandleFunc("DELETE /anime/{id}", server.DeleteAnime)

	http.ListenAndServe(":8080", mux)
}

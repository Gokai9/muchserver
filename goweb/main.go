package main

import (
	"goweb/controller"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := controller.New()
	auth := controller.Newauth()
	mux.HandleFunc("POST /register", auth.Register)
	mux.HandleFunc("GET /login", auth.Login)
	mux.HandleFunc("GET /anime", server.GetAllAnime)
	mux.HandleFunc("POST /anime", server.CreateAnime)
	mux.HandleFunc("GET /anime/{id}", server.GetAnime)
	mux.HandleFunc("PUT /anime/{id}", server.UpdateAnime)
	mux.HandleFunc("DELETE /anime/{id}", server.DeleteAnime)

	http.ListenAndServe(":8080", mux)
}

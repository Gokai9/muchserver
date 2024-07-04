package main

import (
	"goweb/controller"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", &controller.HomeHandler{})
	mux.Handle("/anime", &controller.AnimeHandler{})
	mux.Handle("/anime/", &controller.AnimeHandler{})
	http.ListenAndServe(":8080", mux)
}

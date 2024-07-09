package controller

import (
	"goweb/models"
	"log"
	"net/http"
)

type AuthHandler struct {
	*models.AuthDB
}

func Newauth() *AuthHandler {
	db, err := models.OpenAuthDb()
	if err != nil {
		log.Fatal(err)
	}
	return &AuthHandler{db}
}

// make jwtauth and send to client
func (auth *AuthHandler) Register(w http.ResponseWriter, req *http.Request)

// get token from client and check either valid or not
func (auth *AuthHandler) Login(w http.ResponseWriter, req *http.Request)

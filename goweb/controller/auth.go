package controller

import (
	"encoding/json"
	"fmt"
	"goweb/models"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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

var (
	key []byte
	t   *jwt.Token
	s   string
)

func GenerateTimestampID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(username string) (string, error) {
	key = []byte("hard")
	t = jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": username,
			"exp": time.Now().Add(time.Hour).Unix(),
		})
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, err
}

func verifyToken(tok string) (*jwt.Token, error) {
	key = []byte("hard")
	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return token, nil
}

// make jwtauth and send to client
func (auth *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	logError(w, err, http.StatusBadRequest)
	var user models.User
	user.Id = GenerateTimestampID()
	err = json.Unmarshal(b, &user)
	logError(w, err, http.StatusBadGateway)
	hashpass, _ := HashPassword(user.Password)
	_, err = auth.AddUser(user.Username, hashpass)
	logError(w, err, http.StatusBadGateway)
}

// get token from client and check either valid or not
func (auth *AuthHandler) Login(w http.ResponseWriter, r *http.Request)

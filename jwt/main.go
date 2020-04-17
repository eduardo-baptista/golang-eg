package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// User ...
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtToken ...
type JwtToken struct {
	Token string `json:"token"`
}

// Execption ...
type Execption struct {
	Message string `json:"message"`
}

// SignIn ...
func SignIn(w http.ResponseWriter, req *http.Request) {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	// set expired time to 5 hours after generated token
	expirationTime := time.Now().Add(time.Hour * 5)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     expirationTime.Unix(),
		"usename": user.Username,
	})

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		log.Fatal(error)
	}

	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

// ProtectedEndpoint ...
func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
	token := strings.Split(req.Header["Authorization"][0], " ")[1]

	validatedToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	isValid := validatedToken.Valid

	json.NewEncoder(w).Encode(isValid)
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/session", SignIn).Methods(http.MethodPost)
	api.HandleFunc("/protected", ProtectedEndpoint).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", r))
}

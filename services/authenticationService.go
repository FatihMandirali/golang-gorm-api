package services

import (
	. "appword-api/dbconnect"
	. "appword-api/dto/request"
	. "appword-api/models"
	. "appword-api/utils"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

func Login(W http.ResponseWriter, r *http.Request) {
	connection := DbInit()
	defer CloseDatabase(connection)

	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	CheckError(err)
	var user User
	connection.First(&user)
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		W.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(user)
	W.WriteHeader(http.StatusOK)
	W.Write([]byte(tokenString))
}

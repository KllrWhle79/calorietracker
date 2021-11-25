package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/KllrWhle79/calorietracker/db"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"time"
)

const MySecretKey = "20036BUTTERFLY*JESSROSE0405*"

type Authentication struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Token struct {
	Admin       string `json:"admin"`
	Username    string `json:"user_name"`
	TokenString string `json:"token"`
}

// swagger:operation POST /login root loginUser
// ---
// summary: Login user
// description: "Logs in a user based on password and loads applicable calorie information. Loads admin console for admins"
// responses:
//  "200": "User logged in, JWT token created"
//  "400": "Bad request"
//  "401": "Bad username or password"
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var authdetails Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := db.LoginUser(authdetails.UserName, authdetails.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	validToken, err := generateJWT(user.UserName, user.Admin)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var token Token
	token.TokenString = validToken
	token.Username = user.UserName
	token.Admin = strconv.FormatBool(user.Admin)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func generateJWT(username string, admin bool) (string, error) {
	var mySigningKey = []byte(MySecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Error creating JWT token: %v", err))
	}

	return tokenString, nil
}

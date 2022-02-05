package api

import (
	"encoding/json"
	"github.com/KllrWhle79/calorietracker/db"
	"net/http"
	"strconv"
)

type authentication struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type token struct {
	Admin       string `json:"admin"`
	FirstName   string `json:"first_name"`
	AcctId      int    `json:"acct_id"`
	TokenString string `json:"token"`
}

// swagger:operation POST /login root loginUser
// ---
// summary: Login user
// description: "Logs in a user based on password and loads applicable calorie information. Loads admin console for admins"
// responses:
//  "200": "user logged in, JWT token created"
//  "400": "Bad request"
//  "401": "Bad username or password"
var loginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var authdetails authentication
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

	validToken, err := generateJWT(user.UserName, user.Admin, int(user.ID))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var token token
	token.TokenString = validToken
	token.AcctId = int(user.ID)
	token.FirstName = user.FirstName
	token.Admin = strconv.FormatBool(user.Admin)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
})

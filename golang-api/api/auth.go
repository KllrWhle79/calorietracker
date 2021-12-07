package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const mySecretKey = "918459fb-2eca-464e-97fe-ba0742b525e3"

type idStruct struct {
	ID int `json:"id"`
}

func authMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqId, foundId := mux.Vars(r)["id"]
		reqUsername, foundUsername := mux.Vars(r)["username"]

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		claimUsername := claims.(jwt.MapClaims)["username"].(string)
		claimUserId := claims.(jwt.MapClaims)["id"].(float64)
		admin := claims.(jwt.MapClaims)["admin"].(bool)
		exp := int64(claims.(jwt.MapClaims)["exp"].(float64))

		if time.Now().Unix() > exp {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token expired"))
			return
		}

		var idJson idStruct
		err = json.NewDecoder(r.Body).Decode(&idJson)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !admin {
			if (foundId && reqId != strconv.Itoa(int(claimUserId))) || idJson.ID != int(claimUserId) || (foundUsername && reqUsername != claimUsername) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Operation Not Allowed"))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func generateJWT(username string, admin bool, id int) (string, error) {
	var mySigningKey = []byte(mySecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["id"] = id
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Error creating JWT token: %v", err))
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(mySecretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

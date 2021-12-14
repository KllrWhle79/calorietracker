package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const mySecretKey = "918459fb-2eca-464e-97fe-ba0742b525e3"

type idStruct struct {
	ID int `json:"acct_id"`
}

func authMiddleWare(next http.Handler, adminOnly bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqId, foundId := mux.Vars(r)["acct_id"]
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
		claimUserId := claims.(jwt.MapClaims)["acct_id"].(float64)
		admin := claims.(jwt.MapClaims)["admin"].(bool)
		exp := int64(claims.(jwt.MapClaims)["exp"].(float64))

		if time.Now().Unix() > exp {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token expired"))
			return
		}

		if adminOnly && !admin {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !admin {
			idJson := idStruct{ID: -1}
			if r.Body != nil {
				bodyStr, err := ioutil.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				err = r.Body.Close()
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				r.Body = ioutil.NopCloser(bytes.NewReader(bodyStr))
				err = json.Unmarshal(bodyStr, &idJson)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}

			if (foundId && reqId != strconv.Itoa(int(claimUserId))) ||
				(idJson.ID > 0 && idJson.ID != int(claimUserId)) ||
				(foundUsername && reqUsername != claimUsername) {
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
	claims["acct_id"] = id
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

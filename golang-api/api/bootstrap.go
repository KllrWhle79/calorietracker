package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Start() {
	router := MakeRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:4000"},
		AllowedMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders: []string{"X-Requested-With", "Content-Type"},
	})

	handler := corsHandler.Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func MakeRouter() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/", rootHandler)
	router.Handle("/ping", pingHandler)
	router.Handle("/login", loginHandler).Methods(http.MethodPost)

	router.Handle("/user", createUser).Methods(http.MethodPut)
	router.Handle("/user", authMiddleWare(getUser, false)).Queries("acct_id", "{acct_id}").Methods(http.MethodGet)
	router.Handle("/user", authMiddleWare(getUser, false)).Queries("username", "{username}").Methods(http.MethodGet)
	router.Handle("/user", authMiddleWare(deleteUser, false)).Queries("acct_id", "{acct_id}").Methods(http.MethodDelete)
	router.Handle("/user", authMiddleWare(deleteUser, false)).Queries("username", "{username}").Methods(http.MethodDelete)
	router.Handle("/user", authMiddleWare(updateUser, false)).Queries("acct_id", "{acct_id}").Methods(http.MethodPost)
	router.Handle("/user", authMiddleWare(updateUser, false)).Queries("username", "{username}").Methods(http.MethodPost)
	router.Handle("/users", authMiddleWare(getAllUsers, true)).Methods(http.MethodGet)

	router.Handle("/calorie", authMiddleWare(createCalorieEntry, false)).Methods(http.MethodPut)
	router.Handle("/calorie", authMiddleWare(getCalorieEntry, false)).Queries("acct_id", "{acct_id}", "cal_id", "{cal_id}").Methods(http.MethodGet)
	router.Handle("/calorie", authMiddleWare(deleteCalorieEntry, false)).Queries("acct_id", "{acct_id}", "cal_id", "{cal_id}").Methods(http.MethodDelete)
	router.Handle("/calorie", authMiddleWare(updateCalorieEntry, false)).Methods(http.MethodPost)
	router.Handle("/calories", authMiddleWare(getAllCaloriesForUser, false)).Queries("acct_id", "{acct_id}").Methods(http.MethodGet)

	return router
}

// swagger:operation GET / root mainPage
// ---
// summary: Return a message for the root call
// description: Will return a canned message when the root URL is hit
// responses:
//	"200": "root handled"
var rootHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "root handled"}
	json.NewEncoder(w).Encode(response)
})

// swagger:operation GET /ping root pingPong
// ---
// summary: Return a message for the ping call
// description: Returns a pong signifying the server is up
// responses:
//	"200": "root handled"
var pingHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "pong ping pong"}
	json.NewEncoder(w).Encode(response)
})

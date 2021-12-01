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
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"X-Requested-With", "Content-Type"},
	})

	handler := corsHandler.Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func MakeRouter() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/", rootHandler)
	router.Handle("/ping", pingHandler)
	router.Handle("/login", loginHandler).Methods("POST")

	router.Handle("/user", createUser).Methods("PUT")
	router.Handle("/user", authMiddleWare(getUser)).Queries("id", "{id}").Methods("GET")
	router.Handle("/user", authMiddleWare(getUser)).Queries("username", "{username}").Methods("GET")
	router.Handle("/user", authMiddleWare(deleteUser)).Queries("id", "{id}").Methods("DELETE")
	router.Handle("/user", authMiddleWare(deleteUser)).Queries("username", "{username}").Methods("DELETE")
	router.Handle("/user", authMiddleWare(updateUser)).Queries("id", "{id}").Methods("POST")
	router.Handle("/user", authMiddleWare(updateUser)).Queries("username", "{username}").Methods("POST")

	router.Handle("/calories", authMiddleWare(createCalorieEntry)).Methods("PUT")
	router.Handle("/calories", authMiddleWare(getCalorieEntry)).Queries("id", "{id}").Methods("GET")
	router.Handle("/calories", authMiddleWare(deleteCalorieEntry)).Queries("id", "{id}").Methods("DELETE")
	router.Handle("/calories", authMiddleWare(updateCalorieEntry)).Queries("id", "{id}").Methods("POST")

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

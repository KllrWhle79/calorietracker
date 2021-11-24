package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Start() {
	router := makeRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"X-Requested-With", "Content-Type"},
	})

	handler := corsHandler.Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func makeRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/ping", PingHandler)
	router.HandleFunc("/login", LoginHandler).Methods("POST")

	router.HandleFunc("/user", CreateUser).Methods("PUT")
	router.HandleFunc("/user", GetUser).Queries("id", "{id}").Methods("GET")
	router.HandleFunc("/user", DeleteUser).Queries("id", "{id}").Methods("DELETE")
	router.HandleFunc("/user", UpdateUser).Queries("id", "{id}").Methods("POST")

	router.HandleFunc("/calories", CreateCalorieEntry).Methods("PUT")
	router.HandleFunc("/calories", GetCaloriesForUser).Queries("id", "{id}").Methods("GET")
	router.HandleFunc("/calories", DeleteCalorieEntry).Queries("id", "{id}").Methods("DELETE")
	router.HandleFunc("/calories", UpdateCalorieEntry).Queries("id", "{id}").Methods("POST")

	return router
}

// swagger:operation GET / root mainPage
// ---
// summary: Return a message for the root call
// description: Will return a canned message when the root URL is hit
// responses:
//	"200": "root handled"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "root handled"}
	json.NewEncoder(w).Encode(response)
}

// swagger:operation GET /ping root pingPong
// ---
// summary: Return a message for the ping call
// description: Returns a pong signifying the server is up
// responses:
//	"200": "root handled"
func PingHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "pong ping pong"}
	json.NewEncoder(w).Encode(response)
}

// swagger:operation POST /login root loginUser
// ---
// summary: Login user
// description: "Logs in a user based on password and loads applicable calorie information. Loads admin console for admins"
// responses:
//  "200": "User logged in, JWT token created"
//  "401": "Bad username or password"
func LoginHandler(w http.ResponseWriter, r *http.Request) {

}

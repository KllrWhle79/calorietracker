package api

import (
	"encoding/json"
	"github.com/KllrWhle79/calorietracker/db"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type User struct {
	Id        string `json:"id"`
	UserName  string `json:"user_name"`
	EmailAddr string `json:"email_addr"`
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
}

//swagger:response userResponse
type UserResponse struct {
	//in:body
	Body User
}

// swagger:operation PUT /user user createUser
// ---
// summary: Creates new user for API
// description: "Creates a new user that must be of type: admin, or user."
// responses:
// 	 "200": "New user created"
//   "400": "Bad request"
// 	 "401": "Unauthorized Request"
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 14)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUserId, err := db.CreateNewUser(userData.UserName, userData.EmailAddr, string(hashedPassword), userData.Admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userData.Id = strconv.Itoa(newUserId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}

// swagger:operation GET /user/{id} user getUser
// ---
// summary: Returns a user based on id
// description: If user does not exist, throws an exception. Will only return yourself if no admin permissions.
// parameters:
// - name: id
//   in: path
//   description: id of the user
//   type: number
//   required: false
// - name: username
//   in: path
//   description: username of the user
//   type: string
//   required: false
// responses:
//   "200":
//     "$ref": "#/responses/userResponse"
//   "400": "Bad request"
//   "401": "Unauthorized Request"
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, foundId := mux.Vars(r)["id"]
	username, foundUsername := mux.Vars(r)["username"]

	if foundId {
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userRow, err := db.GetUserById(intId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userData := User{
			Id:        strconv.Itoa(userRow.Id),
			UserName:  userRow.UserName,
			EmailAddr: userRow.EmailAddr,
			Password:  userRow.Password,
			Admin:     userRow.Admin,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userData)
	} else if foundUsername {
		userRow, err := db.GetUserByUsername(username)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userData := User{
			Id:        strconv.Itoa(userRow.Id),
			UserName:  userRow.UserName,
			EmailAddr: userRow.EmailAddr,
			Password:  userRow.Password,
			Admin:     userRow.Admin,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userData)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// swagger:operation POST /user/{id} user updateUser
// ---
// summary: Updates a user based on id
// description: If Admin, can update any user. Otherwise can only update self.
// parameters:
// - name: id
//   in: path
//   description: id of the user
//   type: string
//   required: false
// - name: username
//   in: path
//   description: username of the user
//   type: string
//   required: false
// responses:
//   "200": "User updated"
//   "400": "Bad request"
//   "401": "Unauthorized Request"
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userData User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 14)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, foundId := mux.Vars(r)["id"]
	username, foundUsername := mux.Vars(r)["username"]

	if foundId {
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateUserById(intId, userData.UserName, userData.EmailAddr, string(hashedPassword), userData.Admin)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else if foundUsername {
		err := db.UpdateUserByUsername(username, userData.UserName, userData.EmailAddr, string(hashedPassword), userData.Admin)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	return
}

// swagger:operation DELETE /user/{id} user deleteUser
// ---
// summary: Deletes a user based on id
// description: If Admin, can delete any user. Otherwise, can only delete self.
// parameters:
// - name: id
//   in: path
//   description: id of the user
//   type: string
//   required: false
// - name: username
//   in: path
//   description: username of the user
//   type: string
//   required: false
// responses:
//   "200": "User deleted"
//   "400": "Bad request"
//   "401": "Unauthorized Request"
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, foundId := mux.Vars(r)["id"]
	username, foundUsername := mux.Vars(r)["username"]

	if foundId {
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.DeleteUserById(intId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		return
	} else if foundUsername {
		err := db.DeleteUserByUsername(username)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

package api

import (
	"encoding/json"
	"github.com/KllrWhle79/calorietracker/db"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type user struct {
	Id        int    `json:"acct_id"`
	UserName  string `json:"user_name"`
	EmailAddr string `json:"email_addr"`
	FirstName string `json:"first_name"`
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
	CalMax    int    `json:"calorie_max"`
}

//swagger:response userResponse
type userResponse struct {
	//in:body
	Body []user
}

func sendUserResp(w http.ResponseWriter, userData []user) {
	userResp := userResponse{
		Body: userData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResp)
}

// swagger:operation PUT /user user createUser
// ---
// summary: Creates new user for API
// description: "Creates a new user that must be of type: admin, or user."
// responses:
// 	 "200": "New user created"
//   "400": "Bad request"
// 	 "401": "Unauthorized Request"
var createUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var userData user
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

	userData.Password = string(hashedPassword)

	newUserId, err := db.CreateNewUser(userData.UserName, userData.EmailAddr, userData.Password, userData.FirstName, userData.Admin, userData.CalMax)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userData.Id = int(newUserId)

	sendUserResp(w, []user{userData})
})

// swagger:operation GET /user/{acct_id} user getUser
// ---
// summary: Returns a user based on id
// description: If user does not exist, throws an exception. Will only return yourself if no admin permissions.
// parameters:
// - name: acct_id
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
var getUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	id, foundId := mux.Vars(r)["acct_id"]
	username, foundUsername := mux.Vars(r)["username"]

	if foundId {
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userRow, err := db.GetUserById(uint(intId))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userData := user{
			Id:        int(userRow.ID),
			UserName:  userRow.UserName,
			EmailAddr: userRow.EmailAddr,
			Password:  userRow.Password,
			Admin:     userRow.Admin,
			CalMax:    userRow.CalMax,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userData)
	} else if foundUsername {
		userRow, err := db.GetUserByUsername(username)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sendUserResp(w, []user{{
			Id:        int(userRow.ID),
			UserName:  userRow.UserName,
			EmailAddr: userRow.EmailAddr,
			Password:  userRow.Password,
			Admin:     userRow.Admin,
			CalMax:    userRow.CalMax,
		}})
	}
})

// swagger:operation GET /users users getAllUsers
// ---
// summary: Retrieves all users from database
// description: If Admin, gets a list of all the users in the database
// responses:
//   "200":
//     "$ref": "#/responses/userResponse"
//   "400": "Bad request"
//   "401": "Unauthorized Request"
var getAllUsers = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var users []user
	usersData, err := db.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, userData := range *usersData {
		users = append(users, user{
			Id:        int(userData.ID),
			UserName:  userData.UserName,
			EmailAddr: userData.EmailAddr,
			Password:  userData.Password,
			Admin:     userData.Admin,
			CalMax:    userData.CalMax,
		})
	}

	sendUserResp(w, users)
})

// swagger:operation POST /user/{acct_id} user updateUser
// ---
// summary: Updates a user based on id
// description: If Admin, can update any user. Otherwise, can only update self.
// parameters:
// - name: acct_id
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
//   "200":
//     "$ref": "#/responses/userResponse"
//   "400": "Bad request"
//   "401": "Unauthorized Request"
var updateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var userData user
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, foundId := mux.Vars(r)["acct_id"]
	username, foundUsername := mux.Vars(r)["username"]

	var userToUpdate *db.Users

	if foundId {
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userToUpdate, err = db.GetUserById(uint(intId))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else if foundUsername {
		userToUpdate, err = db.GetUserByUsername(username)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if userData.Password != userToUpdate.Password {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 14)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userData.Password = string(hashedPassword)
	}

	err = db.UpdateUserById(userToUpdate.ID, userData.UserName, userData.EmailAddr, userData.Password, userData.FirstName, userData.Admin, userData.CalMax)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	sendUserResp(w, []user{userData})
})

// swagger:operation DELETE /user/{acct_id} user deleteUser
// ---
// summary: Deletes a user based on id
// description: If Admin, can delete any user. Otherwise, can only delete self.
// parameters:
// - name: acct_id
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
//   "200": "user deleted"
//   "400": "Bad request"
//   "401": "Unauthorized Request"
var deleteUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	id, foundId := mux.Vars(r)["acct_id"]
	username, foundUsername := mux.Vars(r)["username"]

	if foundId {
		intId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.DeleteUserById(uint(intId))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else if foundUsername {
		err := db.DeleteUserByUsername(username)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
})

//TODO: Add password update endpoint

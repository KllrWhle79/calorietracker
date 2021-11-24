package api

import "net/http"

type User struct {
	ID        string `json:"id"`
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	EmailAddr string `json:"email_addr"`
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
// 	 "401": "Unauthorized Request"
func CreateUser(w http.ResponseWriter, r *http.Request) {
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
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/userResponse"
//   "401": "Unauthorized Request"
func GetUser(w http.ResponseWriter, r *http.Request) {
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
//   required: true
// responses:
//   "200": "User updated"
//   "401": "Unauthorized Request"
func UpdateUser(w http.ResponseWriter, r *http.Request) {
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
//   required: true
// responses:
//   "200": "User deleted"
//   "401": "Unauthorized Request"
func DeleteUser(w http.ResponseWriter, r *http.Request) {
}

package api

import "net/http"

type calorie struct {
	ID       string `json:"id"`
	AcctID   string `json:"acct_id"`
	Date     string `json:"date"`
	Calories string `json:"calories"`
}

//swagger:response calorieResponse
type calorieResponse struct {
	//in:body
	Body calorie
}

// swagger:operation PUT /calories calories createCalorieEntry
// ---
// summary: Create new calorie entry
// description: "Creates a calorie entry for a specific user."
// responses:
//  "200": "New calorie entry created"
//  "401": "Unauthorized Request"
var createCalorieEntry = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
})

// swagger:operation GET /calories/{id} calories getCalorieEntries
// ---
// summary: Returns a calories entry based on the acct id
// description: "Retrieves all the calorie entries for a specified user."
// parameters:
// - name: id
//   in: path
//   description: id of the user
//   type: number
//   required: true
// responses:
//  "200":
//    "$ref": "#/responses/calorieResponse
//  "401": "Unauthorized Request"
var getCalorieEntry = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
})

// swagger:operation POST /calories/{id} calories updateCalorieEntry
// ---
// summary: Updates a calorie entry
// description: "Update a calorie entry based on the row id."
// parameters:
// - name: id
//   in: path
//   description: id of the user
//   type: number
//   required: true
// responses:
//  "200": "calorie entry updated"
//  "401": "Unauthorized Request"
var updateCalorieEntry = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
})

// swagger:operation DELETE /calories/{id} calories deleteCalorieEntry
// ---
// summary: Deletes a calorie entry
// description: "Delete a calorie entry based on the row id."
// parameters:
// - name: id
//   in: path
//   description: id of the user
//   type: number
//   required: true
// responses:
//  "200": "calorie entry deleted"
//  "401": "Unauthorized Request"
var deleteCalorieEntry = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
})

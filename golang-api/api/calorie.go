package api

import (
	"encoding/json"
	"github.com/KllrWhle79/calorietracker/db"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type calorie struct {
	ID       int    `json:"id"`
	AcctID   int    `json:"acct_id"`
	Date     string `json:"date"`
	Calories int    `json:"calories"`
}

//swagger:response calorieResponse
type calorieResponse struct {
	//in:body
	Body []calorie
}

func sendCalResp(w http.ResponseWriter, calData []calorie) {
	calorieResp := calorieResponse{
		Body: calData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calorieResp)
}

// swagger:operation PUT /calories calories createCalorieEntry
// ---
// summary: Create new calorie entry
// description: "Creates a calorie entry for a specific user."
// responses:
//  "200": "New calorie entry created"
//  "401": "Unauthorized Request"
var createCalorieEntry = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var calorieData calorie
	err := json.NewDecoder(r.Body).Decode(&calorieData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calTime, err := time.Parse("2006-01-02 15:04:05", calorieData.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calId, err := db.CreateNewCalorieRow(calorieData.Calories, calorieData.AcctID, calTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calorieData.ID = int(calId)

	sendCalResp(w, []calorie{calorieData})
})

// swagger:operation GET /calories/id={id}&calId={calId} calories getCalorieEntries
// ---
// summary: Returns a calories entry based on the acct id
// description: "Retrieves a single calorie entry for a specified user."
// parameters:
// - name: id
//   in: path
//   description: id of the user
//   type: number
//   required: true
// - name: calId
//   in: path
//   description: id of the calorie row
//   type: number
//   required: true
// responses:
//  "200":
//    "$ref": "#/responses/calorieResponse
//  "401": "Unauthorized Request"
var getCalorieEntry = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	userId, foundUserId := mux.Vars(r)["id"]
	userIdInt, err := strconv.Atoi(userId)
	if !foundUserId || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calorieRowId, foundCalRowId := mux.Vars(r)["calId"]
	calRowIdInt, err := strconv.Atoi(calorieRowId)
	if !foundCalRowId || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calorieRow, err := db.GetCalorieRowById(userIdInt, uint(calRowIdInt))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sendCalResp(w, []calorie{{
		ID:       calorieRow.Calories,
		AcctID:   calorieRow.AcctId,
		Date:     calorieRow.Date.String(),
		Calories: calorieRow.Calories,
	}})
})

// swagger:operation GET /calories/id={id} calories getAllCaloriesForUser
// ---
// summary: Returns all the calories entries based on the acct id
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
var getAllCaloriesForUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	userId, foundUserId := mux.Vars(r)["id"]
	userIdInt, err := strconv.Atoi(userId)
	if !foundUserId || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calorieRows, err := db.GetAllCalorieRowsByAcctId(userIdInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var calories []calorie
	for _, calorieRow := range *calorieRows {
		calories = append(calories, calorie{
			ID:       calorieRow.Calories,
			AcctID:   calorieRow.AcctId,
			Date:     calorieRow.Date.String(),
			Calories: calorieRow.Calories,
		})
	}

	sendCalResp(w, calories)
})

// swagger:operation POST /calories/id={id}&calId={calId} calories updateCalorieEntry
// ---
// summary: Updates a calorie entry
// description: "Update a calorie entry based on the row id."
// parameters:
// - name: id
//   in: path
//   description: id of the user
//   type: number
//   required: true
// - name: calId
//   in: path
//   description: id of the calorie row
//   type: number
//   required: true
// responses:
//  "200": "calorie entry updated"
//  "401": "Unauthorized Request"
var updateCalorieEntry = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var calorieData calorie
	err := json.NewDecoder(r.Body).Decode(&calorieData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	calTime, err := time.Parse("2006-01-02 15:04:05", calorieData.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.UpdateCalorieRow(uint(calorieData.ID), calorieData.AcctID, calorieData.Calories, calTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sendCalResp(w, []calorie{calorieData})
})

// swagger:operation DELETE /calories/id={id}&calId={calId} calories deleteCalorieEntry
// ---
// summary: Deletes a calorie entry
// description: "Delete a calorie entry based on the row id."
// parameters:
// - name: id
//   in: path
//   description: id of the user, used for auth purposes only
//   type: number
//   required: true
// - name: calId
//   in: path
//   description: id of the calorie row
//   type: number
//   required: true
// responses:
//  "200": "calorie entry deleted"
//  "401": "Unauthorized Request"
var deleteCalorieEntry = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	calorieRowId, foundCalRowId := mux.Vars(r)["calId"]
	calRowIdInt, err := strconv.Atoi(calorieRowId)
	if !foundCalRowId || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.DeleteCalorieRow(uint(calRowIdInt))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
})

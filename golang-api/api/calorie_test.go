package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var testCalorieJson = `{"acct_id": %d,"calories": %d,"date":"%s"}`

func TestCreateNewCalorieEntry(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))

	response := createTestCalorie(t, []byte(calorieJson))
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestCreateNewCalorieEntryBadJson(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	calorieJson := fmt.Sprintf(`{"acct_id": %d,"calories": "BadCalorie","date":"%s"}`, testUserData.Body[0].Id, time.Now().Format("2006-01-02 15:04:05"))

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/calorie"), bytes.NewBuffer([]byte(calorieJson)))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

func TestCreateNewCalorieEntryFailBadDate(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	calorieJson := fmt.Sprintf(`{"acctId": %d,"calorie": %d,"time":"%s"}`, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/calorie"), bytes.NewBuffer([]byte(calorieJson)))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

func TestGetAllCaloriesForUser(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	createTestCalorie(t, []byte(calorieJson))
	calorieJson = fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(3000), time.Now().Format("2006-01-02 15:04:05"))
	createTestCalorie(t, []byte(calorieJson))

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/calories?acct_id=%d", testUserData.Body[0].Id), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var calResp calorieResponse
	json.NewDecoder(response.Body).Decode(&calResp)
	if len(calResp.Body) != 2 {
		t.Error("Not enough users retrieved")
		t.Fail()
		return
	}

	cleanUp()
}

func TestGetAllCaloriesForUserBadAcctId(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	createTestCalorie(t, []byte(calorieJson))
	calorieJson = fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(3000), time.Now().Format("2006-01-02 15:04:05"))
	createTestCalorie(t, []byte(calorieJson))

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/calories?acct_id=BadAcctId"), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

func TestGetAllCaloriesForUserNoRows(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/calories?acct_id=%d", testUserData.Body[0].Id), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var calResp calorieResponse
	json.NewDecoder(response.Body).Decode(&calResp)
	if len(calResp.Body) != 0 {
		t.Error("Too many calories returned")
		t.Fail()
		return
	}

	cleanUp()
}

func TestGetCalorieByCalorieID(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)
	var calData2 calorieResponse
	calorieJson = fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(3000), time.Now().Format("2006-01-02 15:04:05"))
	response = createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData2)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/calorie?acct_id=%d&cal_id=%d", testUserData.Body[0].Id, calData1.Body[0].ID), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var calResp calorieResponse
	json.NewDecoder(response.Body).Decode(&calResp)
	if calResp.Body[0].Calories != calData1.Body[0].Calories {
		t.Fail()
	}

	cleanUp()
}

func TestGetCalorieByCalorieIDBadAcctId(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)
	var calData2 calorieResponse
	calorieJson = fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(3000), time.Now().Format("2006-01-02 15:04:05"))
	response = createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData2)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/calorie?acct_id=BadAcctId&cal_id=%d", calData1.Body[0].ID), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

func TestGetCalorieByCalorieIDBadCalRowId(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)
	var calData2 calorieResponse
	calorieJson = fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(3000), time.Now().Format("2006-01-02 15:04:05"))
	response = createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData2)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/calorie?acct_id=%d&cal_id=BadCalRowId", testUserData.Body[0].Id), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

func TestGetCalorieByCalorieIDNoRow(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)
	var calData2 calorieResponse
	calorieJson = fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(3000), time.Now().Format("2006-01-02 15:04:05"))
	response = createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData2)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/calorie?acct_id=%d&cal_id=%d", testUserData.Body[0].Id, calData1.Body[0].ID+10), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var calResp calorieResponse
	json.NewDecoder(response.Body).Decode(&calResp)
	if len(calResp.Body) != 0 {
		t.Error("Not enough users retrieved")
		t.Fail()
	}

	cleanUp()
}

func TestDeleteCalorie(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/calorie?acct_id=%d&cal_id=%d", testUserData.Body[0].Id, calData1.Body[0].ID), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestDeleteCalorieBadCalRowId(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/calorie?acct_id=%d&cal_id=BadCalRowId", testUserData.Body[0].Id), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

func TestUpdateCalorie(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)

	newCal := calData1.Body[0]
	newCal.Calories = 10000
	newCalStr, _ := json.Marshal(newCal)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/calorie"), bytes.NewBuffer(newCalStr))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/calorie?acct_id=%d&cal_id=%d", testUserData.Body[0].Id, calData1.Body[0].ID), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var calResp calorieResponse
	json.NewDecoder(response.Body).Decode(&calResp)
	if calResp.Body[0].Calories != 10000 {
		t.Fail()
	}

	cleanUp()
}

func TestUpdateCalorieBadJson(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)

	newCalStr := fmt.Sprintf(`{"acct_id": %d,"calories": "BadCalories","time":"BadTime"}`, calData1.Body[0].ID)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/calorie"), bytes.NewBuffer([]byte(newCalStr)))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

func TestUpdateCalorieBadDate(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)

	newCalStr := fmt.Sprintf(`{"acct_id": %d,"calories": %d,"time":"BadTime"}`, calData1.Body[0].ID, 10000)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/calorie"), bytes.NewBuffer([]byte(newCalStr)))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

func TestUpdateCalorieNoSuchRow(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	var calData1 calorieResponse
	calorieJson := fmt.Sprintf(testCalorieJson, testUserData.Body[0].Id, rand.Intn(4000), time.Now().Format("2006-01-02 15:04:05"))
	response := createTestCalorie(t, []byte(calorieJson))
	json.NewDecoder(response.Body).Decode(&calData1)

	newCalStr := fmt.Sprintf(`{"acct_id": %d, "cal_id":%d, "calories": %d,"time":"%s"}`, calData1.Body[0].AcctID, calData1.Body[0].ID+10, 10000, calData1.Body[0].Date)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/calorie"), bytes.NewBuffer([]byte(newCalStr)))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	cleanUp()
}

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/KllrWhle79/calorietracker/config"
	"github.com/KllrWhle79/calorietracker/db"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testUser       = []byte(`{"user_name": "test_user1","password": "password","email_addr": "test_user1@email.com","first_name": "User1","admin": false}`)
	testUser2      = []byte(`{"user_name": "test_user2","password": "password","email_addr": "test_user2@email.com","first_name": "User2","admin": false}`)
	testUserUpdate = []byte(`{"user_name": "test_user1","password": "password1","email_addr": "test_user1@email.com","first_name": "UpdatedUser","admin": false}`)
	testAdminUser  = []byte(`{"user_name": "admin_test_user1","password": "password","email_addr": "admin_test_user1@email.com","first_name": "Admin","admin": true}`)
)

var testRouter *mux.Router
var testTokenData token
var testUserData userResponse

func testSetup() error {
	testRouter = MakeRouter()

	config.InitConfig()

	return db.InitTables(true)
}

func cleanUp() {
	db.DropTables()
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	testRouter.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func createUserForTest(t *testing.T, userStr []byte) *httptest.ResponseRecorder {
	req, err := http.NewRequest(http.MethodPut, "/user", bytes.NewBuffer(userStr))
	if err != nil {
		t.Error(err)
		t.Fail()
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	json.NewDecoder(response.Body).Decode(&testUserData)

	return response
}

func loginTestUser(t *testing.T, admin bool) {
	var loginJson []byte
	if admin {
		loginJson = []byte(`{"user_name": "admin_test_user1","password": "password"}`)
	} else {
		loginJson = []byte(`{"user_name": "test_user1","password": "password"}`)
	}
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(loginJson))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	json.NewDecoder(response.Body).Decode(&testTokenData)
}

func createTestCalorie(t *testing.T, caloreStr []byte) *httptest.ResponseRecorder {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/calorie"), bytes.NewBuffer(caloreStr))
	if err != nil {
		t.Error(err)
		t.Fail()
		return nil
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	return response
}

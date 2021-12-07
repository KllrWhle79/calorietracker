package api

import (
	"bytes"
	"encoding/json"
	"github.com/KllrWhle79/calorietracker/config"
	"github.com/KllrWhle79/calorietracker/db"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testUser       = []byte(`{"user_name": "test_user1","password": "password","email_addr": "test_user1@email.com","admin": false}`)
	testUserUpdate = []byte(`{"user_name": "test_user1","password": "password1","email_addr": "test_user1@email.com","admin": false}`)
	testAdminUser  = []byte(`{"user_name": "admin_test_user1","password": "password","email_addr": "admin_test_user1@email.com","admin": true}`)
)

var testRouter *mux.Router
var testTokenData token
var testUserData user

func testSetup() error {
	testRouter = MakeRouter()

	config.InitConfig()

	return db.InitTables(true)
}

func cleanUp() {
	dropTables()
	dropSequences()
}

func dropTables() {
	for _, tbl := range db.TableList {
		if db.TableExists(tbl) {
			db.DropTable(tbl)
		}
	}
}

func dropSequences() {
	for _, seq := range db.TableList {
		if db.SequeneceExists(seq + "_seq") {
			db.DropSequence(seq + "_seq")
		}
	}
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

func createUserForTest(t *testing.T, admin bool) *httptest.ResponseRecorder {
	var userStr []byte
	if admin {
		userStr = testAdminUser
	} else {
		userStr = testUser
	}
	req, err := http.NewRequest("PUT", "/user", bytes.NewBuffer(userStr))
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
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJson))
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

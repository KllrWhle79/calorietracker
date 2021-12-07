package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateUserAndLogin(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	req, err := http.NewRequest("PUT", "/user", bytes.NewBuffer(testAdminUser))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var userData user
	json.NewDecoder(response.Body).Decode(&userData)

	loginJson := []byte(`{"user_name": "admin_test_user1","password": "password"}`)
	req, err = http.NewRequest("POST", "/login", bytes.NewBuffer(loginJson))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var tokenData token
	json.NewDecoder(response.Body).Decode(&tokenData)

	req, err = http.NewRequest("GET", fmt.Sprintf("/user?id=%s", userData.Id), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenData.TokenString))

	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestCreateUserAndLoginFail(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	req, err := http.NewRequest("PUT", "/user", bytes.NewBuffer(testAdminUser))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var userData user
	json.NewDecoder(response.Body).Decode(&userData)

	loginJson := []byte(`{"user_name": "admin_test_user1","password": "badpassword"}`)
	req, err = http.NewRequest("POST", "/login", bytes.NewBuffer(loginJson))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)
	checkResponseCode(t, http.StatusUnauthorized, response.Code)

	cleanUp()
}
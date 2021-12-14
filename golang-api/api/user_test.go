package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestGetSameUserById(t *testing.T) {
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

	req, err := http.NewRequest("GET", fmt.Sprintf("/user?acct_id=%d", testUserData.Body[0].Id), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestGetDifferentUserNotAdmin(t *testing.T) {
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

	req, err := http.NewRequest("GET", fmt.Sprintf("/user?acct_id=100"), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusUnauthorized, response.Code)

	cleanUp()
}

func TestGetDifferentUserAdmin(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	nonAdminId := testUserData.Body[0].Id
	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/user?acct_id=%d", nonAdminId), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestGetUserByIdBadId(t *testing.T) {
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

	req, err := http.NewRequest("GET", fmt.Sprintf("/user?acct_id=badId"), nil)
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

func TestGetUserByIdNoUser(t *testing.T) {
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

	req, err := http.NewRequest("GET", fmt.Sprintf("/user?acct_id=100"), nil)
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

func TestGetSameUserByUsername(t *testing.T) {
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

	req, err := http.NewRequest("GET", fmt.Sprintf("/user?username=%s", testUserData.Body[0].UserName), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestGetUserByUsernameNoUser(t *testing.T) {
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

	req, err := http.NewRequest("GET", fmt.Sprintf("/user?username=badusername"), nil)
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

func TestDeleteUserById(t *testing.T) {
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

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/user?acct_id=%d", testUserData.Body[0].Id), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestDeleteUserByIdBadId(t *testing.T) {
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

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/user?acct_id=badid"), nil)
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

func TestDeleteUserByIdNoUser(t *testing.T) {
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

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/user?acct_id=100"), nil)
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

func TestDeleteUserByUsername(t *testing.T) {
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

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/user?username=%s", testUserData.Body[0].UserName), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestDeleteUserByUsernameNoUser(t *testing.T) {
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

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/user?username=badusername"), nil)
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

func TestUpdateUserById(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	userId := testUserData.Body[0].Id
	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("/user?acct_id=%d", userId), bytes.NewBuffer(testUserUpdate))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestUpdateUserByIdFailNoUser(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("/user?acct_id=100"), bytes.NewBuffer(testUserUpdate))
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

func TestUpdateUserByIdFailBadId(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("/user?acct_id=badid"), bytes.NewBuffer(testUserUpdate))
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

func TestUpdateUserByUsername(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	userName := testUserData.Body[0].UserName
	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("/user?username=%s", userName), bytes.NewBuffer(testUserUpdate))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestUpdateUserByUsernameFailNoUser(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("/user?username=nouser"), bytes.NewBuffer(testUserUpdate))
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

func TestGetAllUsers(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, testUser)
	createUserForTest(t, testUser2)
	createUserForTest(t, testAdminUser)
	loginTestUser(t, true)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/users"), nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testTokenData.TokenString))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var userResponse userResponse
	json.NewDecoder(response.Body).Decode(&userResponse)
	if len(userResponse.Body) != 3 {
		t.Error("Not enough users retrieved")
		t.Fail()
		return
	}

	cleanUp()
}

package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCreateNewCalorieEntry(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	createUserForTest(t, false)
	loginTestUser(t, false)

	if testTokenData.TokenString == "" {
		t.Error("Error creating user or token data")
		t.Fail()
		return
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("/calorie?id=%s", testUserData.Id), nil)
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

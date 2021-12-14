package api

import (
	"bytes"
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

	req, err := http.NewRequest("PUT", fmt.Sprintf("/calories"), bytes.NewBuffer([]byte(calorieJson)))
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

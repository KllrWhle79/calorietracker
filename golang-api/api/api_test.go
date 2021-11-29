package api

import (
	"bytes"
	"net/http"
	"testing"
)

func TestCreateUserAndLogin(t *testing.T) {
	err := ApiTestSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	req, err := http.NewRequest("PUT", "/user", bytes.NewBuffer(TestAdminUser))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/json")

	response := ExecuteRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)

	ApiCleanUp()
}

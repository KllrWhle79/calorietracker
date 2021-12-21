package api

import (
	"net/http"
	"testing"
)

func TestRootHandler(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

func TestPingHandler(t *testing.T) {
	err := testSetup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	cleanUp()
}

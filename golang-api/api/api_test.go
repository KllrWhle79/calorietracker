package api

import (
	"bytes"
	"net/http"
	"testing"
)

var (
	testUser      = []byte(`{"user_name": "test_user1","password": "password","email_addr": "test_user1@email.com","admin": "false"}`)
	testAdminUser = []byte(`{"user_name": "admin_test_user1","password": "password","email_addr": "admin_test_user1@email.com","admin": "true"}`)
)

func TestCreateUserAndLogin(t *testing.T) {
	req, _ := http.NewRequest("/PUT", "/user", bytes.NewBuffer(testAdminUser))
	req.Header.Set("Content-Type", "application/json")
}

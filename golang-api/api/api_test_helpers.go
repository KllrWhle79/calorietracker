package api

import (
	"github.com/KllrWhle79/calorietracker/config"
	"github.com/KllrWhle79/calorietracker/db"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	TestUser      = []byte(`{"user_name": "test_user1","password": "password","email_addr": "test_user1@email.com","admin": false}`)
	TestAdminUser = []byte(`{"user_name": "admin_test_user1","password": "password","email_addr": "admin_test_user1@email.com","admin": true}`)
)

var Router *mux.Router

func TestSetup() error {
	Router = MakeRouter()

	config.InitConfig()

	return db.InitTables(true)
}

func CleanUp() {
	DropTables()
	DropSequences()
}

func DropTables() {
	for _, tbl := range db.TableList {
		if db.TableExists(tbl) {
			db.DropTable(tbl)
		}
	}
}

func DropSequences() {
	for _, seq := range db.TableList {
		if db.SequeneceExists(seq + "_seq") {
			db.DropSequence(seq + "_seq")
		}
	}
}

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	Router.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

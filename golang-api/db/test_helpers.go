package db

import (
	"github.com/KllrWhle79/calorietracker/config"
)

var (
	testUser = UsersDBRow{
		UserName:  "test_user1",
		Password:  "password",
		EmailAddr: "test_user1@email.com",
		Admin:     false,
	}
	testAdminUser = UsersDBRow{
		UserName:  "admin_test_user",
		Password:  "password",
		EmailAddr: "admin_test_user@email.com",
		Admin:     true,
	}
)

func Setup() error {
	config.InitConfig()

	err := InitTables(true)
	return err
}

func CleanUp() {
	DropTables()
	DropSequences()
}

func DropTables() {
	for _, tbl := range TableList {
		if TableExists(tbl) {
			DropTable(tbl)
		}
	}
}

func DropSequences() {
	for _, seq := range TableList {
		if SequeneceExists(seq + "_seq") {
			DropSequence(seq + "_seq")
		}
	}
}

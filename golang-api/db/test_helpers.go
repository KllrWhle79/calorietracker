package db

import (
	"github.com/KllrWhle79/calorietracker/config"
)

var (
	testUser = Users{
		UserName:  "test_user1",
		Password:  "password",
		EmailAddr: "test_user1@email.com",
		Admin:     false,
	}
	testAdminUser = Users{
		UserName:  "admin_test_user",
		Password:  "password",
		EmailAddr: "admin_test_user@email.com",
		Admin:     true,
	}
)

func setup() error {
	config.InitConfig()

	err := InitTables(true)
	return err
}

func cleanUp() {
	DropTables()
}

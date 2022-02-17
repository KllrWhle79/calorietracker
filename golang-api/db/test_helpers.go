package db

import (
	"github.com/KllrWhle79/calorietracker/config"
)

var (
	testUser = Users{
		UserName:  "test_user1",
		Password:  "password",
		FirstName: "Test",
		EmailAddr: "test_user1@email.com",
		Admin:     false,
		CalMax:    2500,
	}
	testAdminUser = Users{
		UserName:  "admin_test_user",
		Password:  "password",
		FirstName: "Admin",
		EmailAddr: "admin_test_user@email.com",
		Admin:     true,
		CalMax:    3000,
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

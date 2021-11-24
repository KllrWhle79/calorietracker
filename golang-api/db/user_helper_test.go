package db

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var (
	testUser = UsersDBRow{
		UserName:  "test_user2",
		Password:  "password",
		EmailAddr: "test_user2@email.com",
		Admin:     false,
	}
	adminUser = UsersDBRow{
		UserName:  "test_user1",
		Password:  "password",
		EmailAddr: "test_user1@email.com",
		Admin:     true,
	}
)

func TestCreateNewUserAndLogin(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	user, err := LoginUser(adminUser.UserName, adminUser.Password)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if id != user.Id {
		t.Error(err)
		t.Fail()
	}

	CleanUp()
}

func TestGetMultipleUsers(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id1, err := CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(testUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id2, err := CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	users, err := GetUsersByIds([]int{id1, id2})
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		if len(*users) != 2 {
			t.Error(err)
			t.Fail()
		}
	}

	CleanUp()
}

func TestCreateNewUserFail(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.Admin)
	if err == nil {
		t.Error(err)
		t.Fail()
	}

	CleanUp()
}

func TestGetUserById(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	user1, err := GetUserById(id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if user1.Id != id || user1.UserName != testUser.UserName || user1.EmailAddr != testUser.EmailAddr {
		t.Error("The wrong user1 was retrieved")
		t.Fail()
	}

	user2, _ := GetUserById(id + 1)
	if user2 != nil {
		t.Error("Should not have found user")
		t.Fail()
	}

	CleanUp()
}

func TestUserLoginBadUsername(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = LoginUser("bad_username", adminUser.Password)
	if err == nil {
		t.Error(err)
		t.Fail()
	}

	CleanUp()
}

func TestUserLoginBadPassword(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = LoginUser(adminUser.UserName, "bad_password")
	if err == nil {
		t.Error(err)
		t.Fail()
	}

	CleanUp()
}

func TestDeleteUserById(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	id, err := CreateNewUser(adminUser.UserName, adminUser.EmailAddr, adminUser.Password, adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = DeleteUserById(id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	CleanUp()
}

func TestCreateAndDeleteUserById(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
	}

	deleted, err := DeleteUserById(id)
	if err != nil || !deleted {
		t.Error(t)
		t.Fail()
	}

	user, err := GetUserById(id)
	if err == nil || user != nil {
		t.Error(t)
		t.Fail()
	}
}

func TestCreateAndDeleteUserByUsername(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
	}

	deleted, err := DeleteUserByUsername(adminUser.UserName)
	if err != nil || !deleted {
		t.Error(t)
		t.Fail()
	}

	user, err := GetUserById(id)
	if err == nil || user != nil {
		t.Error(t)
		t.Fail()
	}
}

func TestDeleteUserByIdFail(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
	}

	deleted, err := DeleteUserById(10000)
	if err == nil || deleted {
		t.Error(t)
		t.Fail()
	}

	CleanUp()
}

func TestDeleteUserByUsernameFail(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
	}

	deleted, err := DeleteUserByUsername("testUser")
	if err == nil || deleted {
		t.Error(t)
		t.Fail()
	}

	CleanUp()
}

func TestUpdateUserById(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte("new_password"), 14)

	err = UpdateUserById(id, "updatedUser", "updatedUser@gmail.com", string(hashedPassword), false)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	CleanUp()
}

func TestUpdateUserByUsername(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte("new_password"), 14)

	err = UpdateUserByUsername(adminUser.UserName, "updatedUser", "updatedUser@gmail.com", string(hashedPassword), false)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	CleanUp()
}

func TestUpdateUserByUsernameFail(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(adminUser.UserName, adminUser.EmailAddr, string(hashedPassword), adminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte("new_password"), 14)

	err = UpdateUserByUsername("updatedUser", "updatedUser", "updatedUser@gmail.com", string(hashedPassword), false)
	if err == nil {
		t.Error(err)
		t.Fail()
	}

	CleanUp()
}

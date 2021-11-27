package db

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCreateNewUserAndLogin(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	user, err := LoginUser(testAdminUser.UserName, testAdminUser.Password)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id1, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = LoginUser("bad_username", testAdminUser.Password)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = LoginUser(testAdminUser.UserName, "bad_password")
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

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, testAdminUser.Password, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	err = DeleteUserById(id)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
	}

	err = DeleteUserById(id)
	if err != nil {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
	}

	err = DeleteUserByUsername(testAdminUser.UserName)
	if err != nil {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
	}

	err = DeleteUserById(10000)
	if err == nil {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
	}

	err = DeleteUserByUsername("testUser")
	if err == nil {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte("new_password"), 14)

	err = UpdateUserByUsername(testAdminUser.UserName, "updatedUser", "updatedUser@gmail.com", string(hashedPassword), false)
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
	}

	_, err = CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.Admin)
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

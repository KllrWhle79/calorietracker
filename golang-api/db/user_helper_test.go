package db

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCreateNewUserAndLogin(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	user, err := LoginUser(testAdminUser.UserName, testAdminUser.Password)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if id != user.ID {
		t.Error(err)
		t.Fail()
		return
	}

	cleanUp()
}

func TestGetMultipleUsers(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id1, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(testUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id2, err := CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.FirstName, testUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	users, err := GetUsersByIds([]uint{id1, id2})
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	} else {
		if len(*users) != 2 {
			t.Error(err)
			t.Fail()
			return
		}
	}

	cleanUp()
}

func TestGetAllUsers(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	_, err = CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(testUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	_, err = CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.FirstName, testUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	users, err := GetAllUsers()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	} else {
		if len(*users) != 2 {
			t.Error(err)
			t.Fail()
			return
		}
	}

	cleanUp()
}

func TestCreateNewUserFail(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	_, err = CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.FirstName, testUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.FirstName, testUser.Admin)
	if err == nil {
		t.Error(err)
		t.Fail()
		return
	}

	cleanUp()
}

func TestGetUserById(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id, err := CreateNewUser(testUser.UserName, testUser.EmailAddr, string(hashedPassword), testUser.FirstName, testUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	user1, err := GetUserById(id)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if user1.ID != id || user1.UserName != testUser.UserName || user1.EmailAddr != testUser.EmailAddr {
		t.Error("The wrong user1 was retrieved")
		t.Fail()
		return
	}

	user2, _ := GetUserById(id + 1)
	if user2 != nil {
		t.Error("Should not have found user")
		t.Fail()
		return
	}

	cleanUp()
}

func TestUserLoginBadUsername(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	_, err = CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = LoginUser("bad_username", testAdminUser.Password)
	if err == nil {
		t.Error(err)
		t.Fail()
		return
	}

	cleanUp()
}

func TestUserLoginBadPassword(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	_, err = CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = LoginUser(testAdminUser.UserName, "bad_password")
	if err == nil {
		t.Error(err)
		t.Fail()
		return
	}

	cleanUp()
}

func TestDeleteUserById(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, testAdminUser.Password, testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	err = DeleteUserById(id)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	cleanUp()
}

func TestCreateAndDeleteUserById(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
		return
	}

	err = DeleteUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
		return
	}

	user, err := GetUserById(id)
	if err == nil || user != nil {
		t.Error(t)
		t.Fail()
		return
	}

	cleanUp()
}

func TestCreateAndDeleteUserByUsername(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
		return
	}

	err = DeleteUserByUsername(testAdminUser.UserName)
	if err != nil {
		t.Error(t)
		t.Fail()
		return
	}

	user, err := GetUserById(id)
	if err == nil || user != nil {
		t.Error(t)
		t.Fail()
		return
	}

	cleanUp()
}

func TestDeleteUserByIdFail(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
		return
	}

	err = DeleteUserById(10000)
	if err == nil {
		t.Error(t)
		t.Fail()
		return
	}

	cleanUp()
}

func TestDeleteUserByUsernameFail(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	_, err = GetUserById(id)
	if err != nil {
		t.Error(t)
		t.Fail()
		return
	}

	err = DeleteUserByUsername("testUser")
	if err == nil {
		t.Error(t)
		t.Fail()
		return
	}

	cleanUp()
}

func TestUpdateUserById(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testAdminUser.Password), 14)
	if err != nil {
		t.Error(fmt.Sprintf("Error hashing password on creation: %v", err))
		t.Fail()
		return
	}

	id, err := CreateNewUser(testAdminUser.UserName, testAdminUser.EmailAddr, string(hashedPassword), testAdminUser.FirstName, testAdminUser.Admin)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte("new_password"), 14)

	err = UpdateUserById(id, "updatedUser", "updatedUser@gmail.com", string(hashedPassword), "Updated", false)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	cleanUp()
}

package db

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

// CreateNewUser /* Creates a new user in the database. Assumes the password has already been hashed.
func CreateNewUser(userName, emailAddress, hashedPassword string, admin bool) (uint, error) {
	user := Users{
		UserName:  userName,
		EmailAddr: emailAddress,
		Password:  hashedPassword,
		Admin:     admin,
	}

	result := DB.Create(&user)
	if result.Error != nil {
		return 0, errors.New(fmt.Sprintf("Error creating user: %v", result.Error))
	}

	return user.ID, nil
}

/*
Retrieves a user from the database based on the where clause.
*/
func getUser(whereClause string) (*Users, error) {
	var user Users
	result := DB.Where(whereClause).First(&user)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Error finding user: %v", result.Error))
	}

	return &user, nil
}

// GetUserByUsername /* Retrieves a user from the database via it's username
func GetUserByUsername(userName string) (*Users, error) {
	whereClause := fmt.Sprintf("user_name='%s'", userName)

	return getUser(whereClause)
}

// GetUserById /* Retrieves a user from the database via it's database id
func GetUserById(id uint) (*Users, error) {
	whereClause := fmt.Sprintf("id='%d'", id)

	return getUser(whereClause)
}

func GetUsersByIds(ids []uint) (*[]Users, error) {
	whereClause := fmt.Sprintf("id in (%s)", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]"))

	var users []Users
	result := DB.Where(whereClause).Find(&users)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Error finding user: %v", result.Error))
	}

	return &users, nil
}

func GetAllUsers() (*[]Users, error) {
	var userRows []Users
	result := DB.Find(&userRows)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Error finding users: %v", result.Error))
	}

	return &userRows, nil
}

// DeleteUserByUsername /* Deletes a user from the database, looks up the user by username
func DeleteUserByUsername(userName string) error {
	user, err := GetUserByUsername(userName)

	if err != nil {
		return errors.New(fmt.Sprintf("Could not find user %s to delete: %v", userName, err))
	}

	result := DB.Delete(&Users{}, user.ID)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error deleting user: %v", result.Error))
	}

	return nil
}

// DeleteUserById /* Deletes a user from the database, looks up the user by database id
func DeleteUserById(id uint) error {
	user, err := GetUserById(id)

	if err != nil {
		return errors.New(fmt.Sprintf("Could not find user to delete: %v", err))
	}

	result := DB.Delete(&Users{}, user.ID)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error deleting user: %v", result.Error))
	}

	return nil
}

// UpdateUserById /* Updates a user's entry in the database, looks up the user by database id
func UpdateUserById(id uint, userName, emailAddr, password string, admin bool) error {
	user := Users{
		Model:     gorm.Model{ID: id},
		UserName:  userName,
		EmailAddr: emailAddr,
		Password:  password,
		Admin:     admin,
	}

	result := DB.Save(&user)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error updating user %d: %v", id, result.Error))
	}

	return nil
}

// LoginUser /*
func LoginUser(userName, password string) (*Users, error) {
	user, err := GetUserByUsername(userName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error finding user %s: %s", userName, err))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return user, nil
	} else {
		return nil, errors.New("bad username or password")
	}
}

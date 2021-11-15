package db

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

/*
Retrieves a user from the database based on the where clause.
*/
func getUser(whereClause string) (*UsersDBRow, error) {
	var userRow UsersDBRow
	err := GetSingleRow("users", strings.Join(UsersColumns, ","), whereClause).
		Scan(&userRow.Id, &userRow.UserName, &userRow.EmailAddr, &userRow.Password, &userRow.Admin)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error finding user: %v", err))
	}

	return &userRow, nil
}

// CreateNewUser /* Creates a new user in the database. Assumes the password has already been hashed.
func CreateNewUser(userName, emailAddress, hashedPassword string, admin bool) (int, error) {
	valuesString := fmt.Sprintf("nextval('users_seq'), '%s','%s', '%s', '%v'", userName, emailAddress, hashedPassword, admin)
	id, err := CreateRow("users", strings.Join(UsersColumns, ","), valuesString, "id")
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Error creating user: %v", err))
	}

	return id, nil
}

// GetUserByUsername /* Retrieves a user from the database via it's username
func GetUserByUsername(userName string) (*UsersDBRow, error) {
	whereClause := fmt.Sprintf("user_name='%s'", userName)

	return getUser(whereClause)
}

// GetUsersById /* Retrieves a user from the database via it's database id
func GetUsersById(id int) (*UsersDBRow, error) {
	whereClause := fmt.Sprintf("id='%d'", id)

	return getUser(whereClause)
}

// DeleteUserByUsername /* Deletes a user from the database, looks up the user by username
func DeleteUserByUsername(userName string) (bool, error) {
	user, err := GetUserByUsername(userName)

	if err != nil {
		return false, errors.New(fmt.Sprintf("Could not find user %s to delete: %v", userName, err))
	}

	err = DeleteRow("users", "id", user.Id)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Error deleting user: %v", err))
	}

	return true, nil
}

// DeleteUserById /* Deletes a user from the database, looks up the user by database id
func DeleteUserById(id int) (bool, error) {
	user, err := GetUsersById(id)

	if err != nil {
		return false, errors.New(fmt.Sprintf("Could not find user to delete: %v", err))
	}

	err = DeleteRow("users", "id", user.Id)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Error deleting user: %v", err))
	}

	return true, nil
}

// UpdateUserByUsername /* Updates a user's entry in the database, looks up the user by username
func UpdateUserByUsername(oldUserNam, userName, emailAddr, password string, admin bool) error {
	user, err := GetUserByUsername(oldUserNam)
	if err != nil {
		return errors.New(fmt.Sprintf("Error finding user %s to update: %v", userName, err))
	}

	return UpdateUserById(user.Id, userName, emailAddr, password, admin)
}

// UpdateUserById /* Updates a user's entry in the database, looks up the user by database id
func UpdateUserById(id int, userName, emailAddr, password string, admin bool) error {
	var updateStrings []string
	for _, col := range UsersColumns {
		switch {
		case strings.HasPrefix(UsersTblCols[col], "text"):
			updateStrings = append(updateStrings, col+"='%s'")
		case strings.HasPrefix(UsersTblCols[col], "integer"):
			updateStrings = append(updateStrings, col+"=%d")
		case strings.HasPrefix(UsersTblCols[col], "boolean"):
			updateStrings = append(updateStrings, col+"=%t")
		}
	}

	updateString := fmt.Sprintf(strings.Join(updateStrings, ","), id, userName, emailAddr, password, admin)

	err := UpdateRow("users", updateString, "id", id)
	if err != nil {
		return errors.New(fmt.Sprintf("Error updating user %d: %v", id, err))
	}

	return nil
}

// LoginUser /*
func LoginUser(userName, password string) (*UsersDBRow, error) {
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

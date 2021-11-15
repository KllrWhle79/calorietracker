package db

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func CreateNewCalorieRow(calories, acctId int, timestamp time.Time) error {
	valuesString := fmt.Sprintf("nextval('users_seq'), %d, %s, %d", acctId, timestamp, calories)
	_, err := CreateRow("calories", strings.Join(CaloriesColumns, ","), valuesString, "id")
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating calorie row: %v", err))
	}

	return nil
}

func DeleteCalorieRow(rowId int) error {
	err := DeleteRow("calories", "id", rowId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error deleting calorie row: %v", err))
	}
	return nil
}

func UpdateCalorieRow(rowId, acctId, calories int, timestamp time.Time) error {
	var updateStrings []string
	for _, col := range UsersColumns {
		switch {
		case strings.HasPrefix(UsersTblCols[col], "text"):
			updateStrings = append(updateStrings, col+"='%s'")
		case strings.HasPrefix(UsersTblCols[col], "integer"):
			updateStrings = append(updateStrings, col+"=%d")
		case strings.HasPrefix(UsersTblCols[col], "timestamp"):
			updateStrings = append(updateStrings, col+"=%s")
		}
	}

	updateString := fmt.Sprintf(strings.Join(updateStrings, ","), rowId, acctId, timestamp, calories)
	err := UpdateRow("users", updateString, "id", rowId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error updating calories %d: %v", rowId, err))
	}

	return nil
}

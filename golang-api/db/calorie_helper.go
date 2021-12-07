package db

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func CreateNewCalorieRow(calories, acctId int, timestamp time.Time) (int, error) {
	valuesString := fmt.Sprintf("nextval('calories_seq'), %d, '%s', %d", acctId, timestamp.Format("2006-01-02 15:04:05"), calories)
	id, err := CreateRow("calories", strings.Join(CaloriesColumns, ","), valuesString, "id")
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Error creating calorie row: %v", err))
	}

	return id, nil
}

func DeleteCalorieRow(rowId int) error {
	err := DeleteRow("calories", "id", rowId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error deleting calorie row: %v", err))
	}
	return nil
}

func GetCalorieRowById(acctId, rowId int) (*CaloriesDBRow, error) {
	whereClause := fmt.Sprintf("acctId='%d' AND id='%d'", acctId, rowId)

	var calorieRow CaloriesDBRow
	err := GetSingleRow("calories", strings.Join(CaloriesColumns, ","), whereClause).
		Scan(&calorieRow.Id, &calorieRow.AcctId, &calorieRow.Date, &calorieRow.Calories)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error finding user: %v", err))
	}

	return &calorieRow, nil
}

func GetAllCalorieRowsByAcctId(acctId int) (*[]CaloriesDBRow, error) {
	var calorieRows []CaloriesDBRow
	rows, err := GetRows("calories", strings.Join(CaloriesColumns, ","), fmt.Sprintf("acctId=%d", acctId), "id")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error getting all calories for user %d: %v", acctId, err))
	}

	for rows.Next() {
		var calorieRow CaloriesDBRow
		err = rows.Scan(&calorieRow.Id, &calorieRow.AcctId, &calorieRow.Date, &calorieRow.Calories)
		if err != nil {
			log.Debugf("Error scanning rowL : %v", err)
			continue
		}

		calorieRows = append(calorieRows, calorieRow)
	}

	return &calorieRows, nil
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

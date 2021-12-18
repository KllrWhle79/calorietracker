package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func CreateNewCalorieRow(calories, acctId int, timestamp time.Time) (uint, error) {
	calorie := Calories{
		AcctId:   acctId,
		Date:     timestamp,
		Calories: calories,
	}
	result := DB.Create(&calorie)
	if result.Error != nil {
		return 0, errors.New(fmt.Sprintf("Error creating calorie row: %v", result.Error))
	}

	return calorie.ID, nil
}

func DeleteCalorieRow(rowId uint) error {
	result := DB.Delete(&Calories{}, rowId)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error deleting calorie row: %v", result.Error))
	}
	return nil
}

func GetCalorieRowById(acctId int, rowId uint) (*Calories, error) {
	var calorie Calories
	result := DB.Where("acct_id = ?  AND id = ?", acctId, rowId).First(&calorie)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Error finding callorie: %v", result.Error))
	}

	return &calorie, nil
}

func GetAllCalorieRowsByAcctId(acctId int) (*[]Calories, error) {
	var calories []Calories
	result := DB.Where("acct_id = ?", acctId).Find(&calories)
	if result.Error != nil {
		return nil, errors.New(fmt.Sprintf("Error getting all calories for user %d: %v", acctId, result.Error))
	}

	return &calories, nil
}

func UpdateCalorieRow(rowId uint, acctId, calories int, timestamp time.Time) error {
	calorie := Calories{
		Model:    gorm.Model{ID: rowId},
		AcctId:   acctId,
		Date:     timestamp,
		Calories: calories,
	}
	result := DB.Save(&calorie)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("Error updating calories %d: %v", rowId, result.Error))
	}

	return nil
}

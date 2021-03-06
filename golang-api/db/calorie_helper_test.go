package db

import (
	"testing"
	"time"
)

var calorieRow = Calories{
	AcctId:   1,
	Date:     time.Now(),
	Calories: 1000,
}

func TestGetCalorieRowById(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	id, err := CreateNewCalorieRow(calorieRow.Calories, calorieRow.AcctId, calorieRow.Date)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	row, err := GetCalorieRowById(calorieRow.AcctId, id)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		if row.ID != id {
			t.Error(err)
			t.Fail()
		}
	}

	cleanUp()
}

func TestGetAllCaloriesByAcctId(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = CreateNewCalorieRow(calorieRow.Calories, calorieRow.AcctId, calorieRow.Date)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = CreateNewCalorieRow(calorieRow.Calories+100, calorieRow.AcctId, time.Now())
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	rows, err := GetAllCalorieRowsByAcctId(calorieRow.AcctId)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		if len(*rows) < 2 {
			t.Fail()
		}
	}

	cleanUp()
}

func TestDeleteCalorieRow(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	id, err := CreateNewCalorieRow(calorieRow.Calories, calorieRow.AcctId, calorieRow.Date)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	err = DeleteCalorieRow(id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	cleanUp()
}

func TestUpdateCalorieRow(t *testing.T) {
	err := setup()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	id, err := CreateNewCalorieRow(calorieRow.Calories, calorieRow.AcctId, calorieRow.Date)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	newRow := Calories{
		AcctId:   1,
		Date:     time.Now(),
		Calories: 2000,
	}

	err = UpdateCalorieRow(id, newRow.AcctId, newRow.Calories, newRow.Date)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	cleanUp()
}

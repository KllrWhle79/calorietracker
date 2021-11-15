package db

import (
	"time"
)

var TableList = []string{"users", "calories"}

var UsersColumns = []string{"id", "user_name", "email_addr", "password", "admin"}

type UsersDBRow struct {
	Id        int    `json:"id"`
	UserName  string `json:"user_name"`
	EmailAddr string `json:"email_addr"`
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
}

var UsersTblCols = map[string]string{
	"id":         "integer PRIMARY KEY",
	"user_name":  "text NOT NULL UNIQUE",
	"email_addr": "text NOT NULL UNIQUE",
	"password":   "text NOT NULL",
	"admin":      "boolean NOT NULL",
}

var CaloriesColumns = []string{"id", "acct_id", "date", "calories"}

type CaloriesDBRow struct {
	Id       int       `json:"id"`
	AcctId   int       `json:"acct_id"`
	Date     time.Time `json:"date"`
	Calories int       `json:"calories"`
}

var CaloriesTblCols = map[string]string{
	"id":       "integer PRIMARY KEY",
	"acct_id":  "integer NOT NULL",
	"date":     "timestamp NOT NULL",
	"calories": "integer NOT NULL",
}

var TblColumns = map[string]map[string]string{
	"users":    UsersTblCols,
	"calories": CaloriesTblCols,
}

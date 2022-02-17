package db

import (
	"gorm.io/gorm"
	"time"
)

var TableList = []string{"users", "calories"}

type Users struct {
	gorm.Model
	UserName  string `gorm:"unique" json:"user_name"`
	EmailAddr string `gorm:"unique" json:"email_addr"`
	FirstName string `json:"first_name"`
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
	CalMax    int    `json:"cal_max"`
}

type Calories struct {
	gorm.Model
	AcctId   int       `json:"acct_id"`
	Date     time.Time `json:"date"`
	Calories int       `json:"calories"`
}

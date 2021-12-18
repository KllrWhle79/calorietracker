package db

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb(test bool) error {
	dbHost := viper.GetString("psqlHost")
	dbUser := viper.GetString("psqlUser")
	dbPassword := viper.GetString("psqlPassword")
	dbDB := viper.GetString("psqlDb")
	var dbPort string
	if test {
		dbPort = viper.GetString("testPort")
	} else {
		dbPort = viper.GetString("psqlPort")
	}

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbDB)
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		return errors.New(fmt.Sprintf("Error initializing Database: %v", err))
	}

	return nil
}

func InitTables(test bool) error {
	if DB == nil {
		err := InitDb(test)
		if err != nil {
			return err
		}
	}

	err := DB.AutoMigrate(Users{}, Calories{})
	if err != nil {
		return err
	}

	return nil
}

func DropTables() {
	for _, tblName := range TableList {
		err := DB.Migrator().DropTable(tblName)
		if err != nil {
			panic(err)
		}
	}
}

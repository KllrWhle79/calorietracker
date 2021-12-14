package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"strings"
)

var DB *sql.DB

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
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("Error initializing Database: %v", err))
	}

	err = DB.Ping()
	if err != nil {
		return errors.New(fmt.Sprintf("Error pinging Database: %v", err))
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

	for _, tableName := range TableList {
		if !TableExists(tableName) {
			CreateTable(tableName)
		}

		if !SequeneceExists(tableName + "_seq") {
			CreateSequence(tableName + "_seq")
		}
	}

	return nil
}

func DropTable(tblName string) {
	sqlQuery := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tblName)

	_, err := DB.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
}

func DropSequence(seqName string) {
	sqlQuery := fmt.Sprintf("DROP SEQUENCE IF EXISTS %s CASCADE", seqName)

	_, err := DB.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
}

func TableExists(tableName string) bool {
	sqlQuery := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE  table_schema = 'public' AND table_name = '%s')", tableName)

	return CheckExists(sqlQuery)
}

func SequeneceExists(seqName string) bool {
	sqlQuery := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.sequences where sequence_schema='public' AND sequence_name='%s')", seqName)

	return CheckExists(sqlQuery)
}

func CheckExists(query string) bool {
	var exists bool
	err := DB.QueryRow(query).Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

func CreateTable(tableName string) {
	colString := ""
	if cols, ok := TblColumns[tableName]; ok {
		for key, val := range cols {
			colString += fmt.Sprintf("%s %s, ", key, val)
		}
		colString = colString[:len(colString)-2]
		sqlQuery := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, colString)

		_, err := DB.Exec(sqlQuery)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("No columns saved for table %s.\n", tableName)
	}
}

func CreateSequence(seqName string) {
	sqlQuery := fmt.Sprintf("CREATE SEQUENCE %s", seqName)

	_, err := DB.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
}

func CreateRow(tblName, columns, values, returnId string) (int, error) {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING %s", tblName, columns, values, returnId)

	var id int
	err := DB.QueryRow(sqlQuery).Scan(&id)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Error creating row : %s", err))
	}
	return id, nil
}

func UpdateRow(tblName, values, keyColumn string, id int) error {
	sqlQuery := fmt.Sprintf("UPDATE %s SET %s WHERE %s=%d", tblName, values, keyColumn, id)

	_, err := DB.Exec(sqlQuery)
	if err != nil {
		return errors.New(fmt.Sprintf("Error updating row: %s", err))
	}
	return nil
}

func DeleteRow(tblName, keyColumn string, id int) error {
	sqlQuery := fmt.Sprintf("DELETE FROM %s WHERE %s=%d", tblName, keyColumn, id)

	_, err := DB.Exec(sqlQuery)
	if err != nil {
		return errors.New(fmt.Sprintf("Error updating row: %s", err))
	}
	return nil
}

func GetSingleRow(tblName, columns, whereClause string) *sql.Row {
	sqlQuery := fmt.Sprintf("SELECT %s FROM %s WHERE %s", columns, tblName, whereClause)

	return DB.QueryRow(sqlQuery)
}

func GetRows(tblName, columns, whereClause, orderBy string) (*sql.Rows, error) {
	var sqlQuery strings.Builder

	sqlQuery.WriteString(fmt.Sprintf("SELECT %s FROM %s", columns, tblName))
	if whereClause != "" {
		sqlQuery.WriteString(fmt.Sprintf(" WHERE %s", whereClause))
	}
	sqlQuery.WriteString(fmt.Sprintf(" ORDER BY %s", orderBy))

	rows, err := DB.Query(sqlQuery.String())
	if err != nil || rows.Err() != nil {
		var errorMsg error
		if err == nil {
			errorMsg = err
		} else {
			errorMsg = rows.Err()
		}
		return nil, errors.New(fmt.Sprintf("Error querying database: %s", errorMsg))
	}
	return rows, nil
}

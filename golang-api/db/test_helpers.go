package db

import "github.com/KllrWhle79/calorietracker/cmd"

func Setup() error {
	cmd.InitConfig()

	err := InitTables(true)
	return err
}

func DropTables() {
	for _, tbl := range TableList {
		if TableExists(tbl) {
			DropTable(tbl)
		}
	}
}

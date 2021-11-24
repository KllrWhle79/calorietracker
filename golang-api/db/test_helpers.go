package db

import "github.com/KllrWhle79/calorietracker/cmd"

func Setup() error {
	cmd.InitConfig()

	err := InitTables(true)
	return err
}

func CleanUp() {
	DropTables()
	DropSequences()
}

func DropTables() {
	for _, tbl := range TableList {
		if TableExists(tbl) {
			DropTable(tbl)
		}
	}
}

func DropSequences() {
	for _, seq := range TableList {
		if SequeneceExists(seq + "_seq") {
			DropSequence(seq + "_seq")
		}
	}
}

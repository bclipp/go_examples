package data_ingestion

import (
	"fmt"
	"testing"
)

func TestPG(t *testing.T) {
	CheckIntegrationTest(t)
	config := GetVariables()
	var database = Database{
		IpAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDb:       config["postgresDb"],
	}

	t.Run("loadTable", func(t *testing.T) {
		err := database.connect()
		if err != nil {fmt.Print(err.Error())}
		if err != nil {
			fmt.Print(err.Error())
		}
		defer database.close()
		err = database.loadTable("customers")
		if err != nil {fmt.Print(err.Error())}
		if len(database.table) < 1 {
			t.Errorf("Error, read customers table and no data was found.")
		}
	})
	t.Run("UpdateTable", func(t *testing.T) {
		err := database.connect()
		if err != nil {fmt.Print(err.Error())}
		defer database.close()
	})

}

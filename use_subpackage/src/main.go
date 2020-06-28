package main

import (
	"api_db_ingestion/src/data_ingestion"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	config := data_ingestion.GetVariables()
	//needs to be mocked
	var database = data_ingestion.Postgresql{
		IpAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDb:       config["postgresDb"],
	}
	tables := []string{
		"customers",
		"stores",
	}
	err := data_ingestion.UpdateTables(false, tables, &database)
	if err != nil {fmt.Print(err.Error())}

}
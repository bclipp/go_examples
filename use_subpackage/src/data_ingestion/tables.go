//This module is used for controling the lookup and update process

package data_ingestion

// update_tables is used for handling the update process
// table field
// Params:
//       concurrent: used to control if the process is done concurrently
//       tables: what tables are going to be updated
// 		 database: the struct used for handling databses.
//return:
//       the error
func UpdateTables(concurrent bool, tables []string, db database) error {
	if !concurrent {
		for _, tableName := range tables {
			err := db.connect()
			if err != nil {return err}
			defer db.close()
			table ,err := db.returnTable(table)
			if err != nil {return err}
			for _, row := range table {
				response, _, err := censusApi(row.Latitude, row.Longitude)
				if err != nil {return err}
				row.BlockId = response.Results[0].BlockId
				row.BlockPop = response.Results[0].BlockPop
				row.StateCode = response.Results[0].StateCode
				row.StateFips = response.Results[0].StateFips
			}
			err = db.updateDbTable(table, tableName)
			if err != nil {return err}
		}
	}
	return nil
}

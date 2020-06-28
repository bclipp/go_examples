//This module is used for holding reusable SQL queries

package data_ingestion

import "fmt"

// update_tables is used for handling the update process
// table field
// Params:
//       table: table to generate the update query for.
//return:
//       the error
func updateTableQuery(table string, row Row) string {
	state_fips := row.StateFips
	state_code := row.StateCode
	block_pop := row.BlockPop
	block_id := row.BlockId
	table_id := row.Id
	return fmt.Sprintf(
		"UPDATE %s SET state_fips = %d, state_code = '%s', block_pop = %d, block_id = %d WHERE ID = %d;",
		table, state_fips, state_code, block_pop, block_id, table_id)
}

// select_table is used for generating a query for selecting a table
// table field
// Params:
//       table: table to generate the update query for.
//		 limit: < 0 will cause assume you don't want a limit
//return:
//       the error
func selectTableQuery(table string, limit int) string {
	if limit < 0 {
		return fmt.Sprintf(
			"SELECT * FROM %s;",
			table)
	} else {
		return fmt.Sprintf(
			"SELECT * FROM %s LIMIT %d;",
			table, limit)
	}
}

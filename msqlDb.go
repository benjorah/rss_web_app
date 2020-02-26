package main

import (
	"database/sql"
	"fmt"
	"log"
)

//MsqlConnection struct defines methods and data properties for connecting with the mySQL databasae
type MsqlConnection struct {
	connection *sql.DB
}

//AddRecords adds new entries into the mysql database
func (msql *MsqlConnection) AddRecords(records []RSSData) (err error) {

	sqlStr := "INSERT INTO feed (title, description, link, published_at) VALUES "
	values := []interface{}{}

	//create the sql string
	for _, row := range records {
		sqlStr += "(?, ?, ?, ?),"
		values = append(values, row.Title, row.Description, row.Link, row.CreatedAt)
	}
	//trim the last `,`
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	//prepare the statement, (sanitize)
	stmt, err := msql.connection.Prepare(sqlStr)

	if err != nil {
		return fmt.Errorf("[ERROR] MsqlConnection.AddRecords() : %s", err.Error())
	}

	//format all entries at once
	_, err = stmt.Exec(values...)

	if err != nil {
		return fmt.Errorf("[ERROR] MsqlConnection.AddRecords() : %s", err.Error())
	}

	fmt.Println("[MYSQL] Sucessfully added entries")

	return err

}

//SearchRecords searches for entries in the database that matches a particular search string and returns those entries
//It returns an error if any error is encountered
func (msql *MsqlConnection) SearchRecords(searchString string) (records []RSSData, err error) {

	records = []RSSData{}
	rows, err := msql.connection.Query("SELECT * FROM feed WHERE MATCH (title, description) AGAINST ('" + searchString + "' IN NATURAL LANGUAGE MODE)")

	if err != nil {
		return nil, fmt.Errorf("[ERROR] MsqlConnection.SearchRecords() : %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		record := RSSData{}
		var id int
		var createdAt []uint8
		fmt.Println(createdAt)
		err = rows.Scan(&id, &record.Title, &record.Description, &record.Link, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("[ERROR] MsqlConnection.SearchRecords() : %s", err.Error())
		}

		parsedTime, err := convertTimeToUTC(string(createdAt) + " +0000 UTC")

		if err != nil {
			return nil, fmt.Errorf("[ERROR] MsqlConnection.SearchRecords() : %s", err.Error())
		}

		record.CreatedAt = parsedTime

		records = append(records, record)

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] MsqlConnection.SearchRecords() : %s", err.Error())
	}

	return records, nil

}

//InitDatabseConnection initializes the algolia client
func (msql *MsqlConnection) InitDatabseConnection() (err error) {

	db, err := sql.Open("mysql", goDotEnvVariable("MYSQL_USERNAME")+":"+goDotEnvVariable("MYSQL_PASSWORD")+"@"+goDotEnvVariable("MYSQL_HOST")+"/"+goDotEnvVariable("MYSQL_DBNAME"))

	if err != nil {
		return fmt.Errorf("[MYSQL] Error opening mysql connection")

	}

	msql.connection = db

	// make sure connection is available
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("[MYSQL] db is not connected, " + err.Error())
	}
	log.Println("[MYSQL] Successfully created database connection")

	return nil

}

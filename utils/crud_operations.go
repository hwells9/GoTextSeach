package utils

import (
	"database/sql"
	"fmt"
	// _ tells go that we want to import so we can use the drivers without ever referencing the library directly in code
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func DbConnect() {
	const (
		user     = "postgres"
		password = "postgrespw"
		host     = "127.0.0.1"
		port     = 5432
		dbName   = "postgres"
	)

	postgresConnectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", postgresConnectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertRow(query string) int64 {
	// Insert example
	result, err := db.Exec(query)
	if err != nil {
		return 0
	}

	// Get the new album's generated ID for the client.
	affected, err := result.RowsAffected()
	if err != nil {
		return 0
	}
	return affected
}

func SelectRows(query string) []Series {
	//Select example
	var series []Series
	rows, err := db.Query(query)
	if err != nil {
		return []Series{}
	}

	for rows.Next() {
		var seriesObject Series
		if err := rows.Scan(&seriesObject.Title, &seriesObject.Description, &seriesObject.Id); err != nil {
			return []Series{}
		}
		series = append(series, seriesObject)
	}

	return series
}

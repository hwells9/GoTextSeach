package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"strconv"
	// _ tells go that we want to import so we can use the drivers without ever referencing the library directly in code
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func DbConnect() {
	var err error
	user := getEnvironmentVariable("DATABASE_USERNAME")
	password := getEnvironmentVariable("DATABASE_PASSWORD")
	host := getEnvironmentVariable("DATABASE_ADDRESS")
	dbName := getEnvironmentVariable("DATABASE_NAME")
	port, err := strconv.Atoi(getEnvironmentVariable("DATABASE_PORT"))

	if err != nil {
		fmt.Printf("problem")
	}

	postgresConnectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err = sqlx.Connect("postgres", postgresConnectionString)
	if err != nil {
		log.Fatalln(err)
	}
}

func insertStruct(dataStruct interface{}, query string) {
	_, err := db.NamedExec(query, dataStruct)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

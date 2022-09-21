package utils

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"

	// _ tells go that we want to import so we can use the drivers without ever referencing the library directly in code
	"log"

	_ "github.com/lib/pq"
)

var Db *sqlx.DB

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

	Db, err = sqlx.Connect("postgres", postgresConnectionString)
	if err != nil {
		log.Fatalln(err)
	}
}

func insertStruct(dataStruct interface{}, query string) {
	_, err := Db.NamedExec(query, dataStruct)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

// func SelectStruct(structType string, query string) []interface{} {

// 	switch(structType) {
// 	case "comic":
// 		myData := []ComicBookDbEntry{}
// 	}

// 	err := db.Select(&myData, query)

// 	if err != nil {
// 		fmt.Printf("Error: %s", err)
// 	}

// 	return myData
// }

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

// var db *sql.DB
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

func SelectRows(query string) []SeriesDbEntry {
	//Select example
	var series []SeriesDbEntry
	rows, err := db.Query(query)
	if err != nil {
		return []SeriesDbEntry{}
	}

	for rows.Next() {
		var seriesObject SeriesDbEntry
		if err := rows.Scan(&seriesObject.Title, &seriesObject.Description, &seriesObject.Id); err != nil {
			return []SeriesDbEntry{}
		}
		series = append(series, seriesObject)
	}

	return series
	//
	//rows, _ := db.Query(query)
	//defer rows.Close()
	//
	//cols, _ := rows.Columns()
	//
	//w := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', 0)
	//defer w.Flush()
	//
	//sep := []byte("\t")
	//newLine := []byte("\n")
	//
	//w.Write([]byte(strings.Join(cols, "\t") + "\n"))
	//
	//row := make([][]byte, len(cols))
	//rowPtr := make([]any, len(cols))
	//for i := range row {
	//	rowPtr[i] = &row[i]
	//}
	//
	//for rows.Next() {
	//	_ = rows.Scan(rowPtr...)
	//
	//	w.Write(bytes.Join(row, sep))
	//	w.Write(newLine)
	//}

}

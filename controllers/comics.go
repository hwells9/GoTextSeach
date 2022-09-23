package controllers

import (
	"bench/textsearch/database"
	"bench/textsearch/tables"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// The request body for search requests
type SearchBody struct {
	SearchTable    string   `json:"searchTable"`
	ResultsColumns []string `json:"resultsColumns"`
	SearchTerm     string   `json:"searchTerm"`
	SearchColumn   string   `json:"searchColumn"`
}

// Represents body of search call
type SearchQuery struct {
	TableName      string
	ResultsColumns []string
	SearchTerm     string
	SearchColumn   string
}

// Gets all series
func ReturnAllSeries(context *gin.Context) {
	fmt.Println("Retrieving all Series")

	var series []tables.Series

	res := database.Db.Find(&series)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": series})
}

// Gets all comic books
func ReturnComicBooks(context *gin.Context) {
	fmt.Println("Retrieving all Comic Books")

	var comicBooks []tables.ComicBook

	res := database.Db.Find(&comicBooks)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": comicBooks})
}

// Gets all characters
func ReturnCharacters(context *gin.Context) {
	fmt.Println("Retrieving all Characters")

	var characters_results []tables.Character

	res := database.Db.Find(&characters_results)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": characters_results})
}

// Gets a series based off id
func ReturnSeries(context *gin.Context) {
	string_id := context.Param("id")
	id, _ := strconv.Atoi(string_id)

	fmt.Printf("Retrieving Series with id: %d", id)

	var series []tables.Series

	res := database.Db.First(&series, id)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": series})
}

// Gets a comic book hased off the id
func ReturnComicBook(context *gin.Context) {
	string_id := context.Param("id")
	id, _ := strconv.Atoi(string_id)

	fmt.Printf("Retrieving Comic Book with id: %d", id)

	var comicBook tables.ComicBook

	res := database.Db.First(&comicBook, id)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": comicBook})
}

// Gets a character based off the id
func ReturnCharacter(context *gin.Context) {
	string_id := context.Param("id")
	id, _ := strconv.Atoi(string_id)

	fmt.Printf("Retrieving Character with id: %d", id)

	var character tables.Character

	res := database.Db.First(&character, id)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": character})
}

// Handles a textbase database search
func SearchDatabase(context *gin.Context) {
	var searchBody SearchBody
	var results []map[string]interface{}

	if err := context.BindJSON(&searchBody); err != nil {
		fmt.Printf("There was an issue getting the Series Search Body: %s", err)
		return
	}

	resultColumns := strings.Join(searchBody.ResultsColumns, ",")

	fmt.Printf("Retrieving %s columns: %s with term: %s in column: %s\n",
		searchBody.SearchTable, resultColumns, searchBody.SearchTerm, searchBody.SearchColumn)

	query := fmt.Sprintf("SELECT %s FROM %s WHERE to_tsvector('english', %s) @@ to_tsquery('english', ?)", resultColumns, searchBody.SearchTable, searchBody.SearchColumn)

	res := database.Db.Raw(query, searchBody.SearchTerm).Scan(&results)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": results})

}

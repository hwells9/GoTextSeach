package main

import (
	"bench/textsearch/utils"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ComicBook struct {
	Title string
	Year  int
	Id    int
}

type SearchQuery struct {
	TableName      string
	ResultsColumns []string
	SearchTerm     string
	SearchColumn   string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllSeries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieving all Series")

	query := "select * from series s"

	results := []utils.SeriesDbEntry{}

	err := utils.Db.Select(&results, query)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(results)
}

func returnComicBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieving all Comic Books")

	query := "select * from comic_books cb"

	results := []utils.ComicBookDbEntry{}

	err := utils.Db.Select(&results, query)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(results)
}

func returnCharacters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieving all Characters")

	query := "select * from characters c"

	results := []utils.CharacterDbEntry{}

	err := utils.Db.Select(&results, query)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(results)
}

func returnSeries(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	fmt.Printf("Retrieving Series with id: %d", key)

	query := fmt.Sprintf("select * from series s where s.id=%d", key)

	result := []utils.SeriesDbEntry{}

	err := utils.Db.Select(&result, query)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(result)
}

func returnComicBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	fmt.Printf("Retrieving Comic Book with id: %d", key)

	cbQuery := fmt.Sprintf("select * from comic_books cb where cb.id=%d", key)

	result := []utils.ComicBookDbEntry{}

	err := utils.Db.Select(&result, cbQuery)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(result)
}

func returnCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	fmt.Printf("Retrieving Character with id: %d", key)

	cbQuery := fmt.Sprintf("select * from characters c where c.id=%d", key)

	result := []utils.CharacterDbEntry{}

	err := utils.Db.Select(&result, cbQuery)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(result)
}

func searchSeries(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var searchQuery SearchQuery

	json.Unmarshal(reqBody, &searchQuery)

	resultColumns := strings.Join(searchQuery.ResultsColumns, ",")
	searchTerm := searchQuery.SearchTerm
	searchColumn := searchQuery.SearchColumn

	fmt.Printf("Retrieving Series with tableNames: %s, %s, %s",
		resultColumns, searchTerm, searchColumn)

	query := "SELECT $1 FROM series WHERE to_tsvector('english', $2) @@ to_tsquery('english', $3)"

	result, err := utils.Db.Query(query, resultColumns, searchColumn, searchTerm)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(result)
}

func searchComicBooks(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var searchQuery SearchQuery

	json.Unmarshal(reqBody, &searchQuery)

	resultColumns := strings.Join(searchQuery.ResultsColumns, ",")
	searchTerm := searchQuery.SearchTerm
	searchColumn := searchQuery.SearchColumn

	fmt.Printf("Retrieving Comicbooks with : %s, %s, %s",
		resultColumns, searchTerm, searchColumn)

	query := "SELECT $1 FROM comic_books WHERE to_tsvector('english', $2) @@ to_tsquery('english', $3)"

	result, err := utils.Db.Query(query, resultColumns, searchColumn, searchTerm)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(result)
}

func searchCharacters(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var searchQuery SearchQuery

	json.Unmarshal(reqBody, &searchQuery)

	resultColumns := strings.Join(searchQuery.ResultsColumns, ",")
	searchTerm := searchQuery.SearchTerm
	searchColumn := searchQuery.SearchColumn

	fmt.Printf("Retrieving Character with tableNames: %s, %s, %s",
		resultColumns, searchTerm, searchColumn)

	query := "SELECT $1 FROM characters WHERE to_tsvector('english', $2) @@ to_tsquery('english', $3)"

	result, err := utils.Db.Query(query, resultColumns, searchColumn, searchTerm)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode(result)
}

// func getLargestComicBookId() int {
// 	currentId := 0

// 	for _, comicBook := range ComicBooks {

// 		if comicBook.Id > currentId {
// 			currentId = comicBook.Id
// 		}
// 	}

// 	return currentId

// }

// func createComicBook(w http.ResponseWriter, r *http.Request) {
// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var comicBook ComicBook

// 	currentId := getLargestComicBookId() + 1

// 	comicBook.Id = currentId

// 	json.Unmarshal(reqBody, &comicBook)

// 	fmt.Printf("Creating Comic Book with id: %d\n", currentId)

// 	ComicBooks = append(ComicBooks, comicBook)

// 	json.NewEncoder(w).Encode(comicBook)
// }

// func deleteComicBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key, _ := strconv.Atoi(vars["id"])

// 	fmt.Printf("Removing Comic Book with id: %d\n", key)

// 	for index, article := range ComicBooks {
// 		if article.Id == key {
// 			ComicBooks = append(ComicBooks[:index], ComicBooks[index+1:]...)
// 		}
// 	}
// }

// func updateComicBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key, _ := strconv.Atoi(vars["id"])

// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	fmt.Printf("Updating Comic Book with id: %d\n", key)

// 	for index, comicBook := range ComicBooks {
// 		if comicBook.Id == key {
// 			json.Unmarshal(reqBody, &ComicBooks[index])
// 		}
// 	}
// }

func handleRequests() {
	// Create router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Routes
	myRouter.HandleFunc("/", homePage)
	// Series endpoints
	myRouter.HandleFunc("/series", returnAllSeries).Methods("GET")
	myRouter.HandleFunc("/series/{id}", returnSeries).Methods("GET")

	// Comic Book endpoints
	myRouter.HandleFunc("/comic-books", returnComicBooks).Methods("GET")
	myRouter.HandleFunc("/comic-books/{id}", returnComicBook).Methods("GET")

	// Characters endpoints
	myRouter.HandleFunc("/characters", returnCharacters).Methods("GET")
	myRouter.HandleFunc("/characters/{id}", returnCharacter).Methods("GET")

	// search endpoints
	myRouter.HandleFunc("/searchSeries", searchSeries).Methods("GET")
	myRouter.HandleFunc("/searchComicBooks", searchComicBooks).Methods("GET")
	myRouter.HandleFunc("/searchCharacters", searchCharacters).Methods("GET")

	// myRouter.HandleFunc("/comic-books", createComicBook).Methods("POST")
	// myRouter.HandleFunc("/comic-books/{id}", deleteComicBook).Methods("DELETE")
	// myRouter.HandleFunc("/comic-books/{id}", updateComicBook).Methods("PUT")

	// Handle fatal errors
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	// get the executeDataMigration parameter from the cmd line
	executeDataMigrationPtr := flag.Bool("executeDataMigration", false, "bool value")
	flag.Parse()

	if *executeDataMigrationPtr {
		// run the export data from marvel api to db
		utils.ExecuteMigration()
	} else {
		// ComicBooks = []ComicBook{
		// 	{Title: "X-Men", Year: 1970, Id: 1},
		// 	{Title: "Avengers", Year: 1856, Id: 2},
		// }
		// Set the location of the environment variables file
		viper.SetConfigFile("env-vars.env")
		utils.DbConnect()
		handleRequests()
	}

}

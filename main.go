package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ComicBook struct {
	Title string
	Year  int
	Id    int
}

var ComicBooks []ComicBook

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnComicBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Retrieving all Comic Books")
	json.NewEncoder(w).Encode(ComicBooks)
}

func returnComicBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	fmt.Printf("Retrieving Comic Book with id: %d", key)

	for _, comicBook := range ComicBooks {
		if comicBook.Id == key {
			json.NewEncoder(w).Encode(comicBook)
		}
	}
}

func getLargestComicBookId() int {
	currentId := 0

	for _, comicBook := range ComicBooks {

		if comicBook.Id > currentId {
			currentId = comicBook.Id
		}
	}

	return currentId

}

func createComicBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var comicBook ComicBook

	currentId := getLargestComicBookId() + 1

	comicBook.Id = currentId

	json.Unmarshal(reqBody, &comicBook)

	fmt.Printf("Creating Comic Book with id: %d\n", currentId)

	ComicBooks = append(ComicBooks, comicBook)

	json.NewEncoder(w).Encode(comicBook)
}

func deleteComicBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	fmt.Printf("Removing Comic Book with id: %d\n", key)

	for index, article := range ComicBooks {
		if article.Id == key {
			ComicBooks = append(ComicBooks[:index], ComicBooks[index+1:]...)
		}
	}
}

func updateComicBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("Updating Comic Book with id: %d\n", key)

	for index, comicBook := range ComicBooks {
		if comicBook.Id == key {
			json.Unmarshal(reqBody, &ComicBooks[index])
		}
	}
}

func handleRequests() {
	// Create router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Routes
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/comic-books", returnComicBooks).Methods("GET")
	myRouter.HandleFunc("/comic-books/{id}", returnComicBook).Methods("GET")
	myRouter.HandleFunc("/comic-books", createComicBook).Methods("POST")
	myRouter.HandleFunc("/comic-books/{id}", deleteComicBook).Methods("DELETE")
	myRouter.HandleFunc("/comic-books/{id}", updateComicBook).Methods("PUT")

	// Handle fatal errors
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	ComicBooks = []ComicBook{
		{Title: "X-Men", Year: 1970, Id: 1},
		{Title: "Avengers", Year: 1856, Id: 2},
	}
	handleRequests()
}

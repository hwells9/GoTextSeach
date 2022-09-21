package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	// _ tells go that we want to import so we can use the drivers without ever referencing the library directly in code
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Global Vars set from environment variables
var privateAPIKey string = ""
var publicAPIKey string = ""

// The number of series to pull in from Marvel API
const seriesLimit int = 20

// Series Response from Marvel API
type SeriesResponse struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Data   SeriesData `json:"data"`
}

// Series' Comics Response from Marvel API
type SeriesComicsResponse struct {
	Code     int              `json:"code"`
	Status   string           `json:"status"`
	Data     SeriesComicsData `json:"data"`
	SeriesId int
}

// Series' Characters Response from Marvel API
type SeriesCharactersResponse struct {
	Code     int                  `json:"code"`
	Status   string               `json:"status"`
	Data     SeriesCharactersData `json:"data"`
	SeriesId int
	ComicId  int
}

// Series Data in SeriesResponse from Marvel API call
type SeriesData struct {
	Offset  int             `json:"offset"`
	Limit   int             `json:"limit"`
	Total   int             `json:"total"`
	Count   int             `json:"count"`
	Results []SeriesResults `json:"results"`
}

// Series' Comics Data in SeriesComicsResponse from Marvel API call
type SeriesComicsData struct {
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	Count   int         `json:"count"`
	Results []ComicBook `json:"results"`
}

// Series Characters Data in SeriesCharactersResponse from Marvel API call
type SeriesCharactersData struct {
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	Count   int         `json:"count"`
	Results []Character `json:"results"`
}

// Series Results in series response from Marvel API call
type SeriesResults struct {
	Id          int              `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ResourceURI string           `json:"resourceURI"`
	Urls        []Url            `json:"urls"`
	StartYear   int              `json:"startYear"`
	EndYear     int              `json:"endYear"`
	Rating      string           `json:"rating"`
	Type        string           `json:"type"`
	Modified    string           `json:"modified"`
	Creators    CreatorsResponse `json:"creators"`
}

// The Url container for Urls in SeriesResults in SeriesResponse from Marvel API call
type Url struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// The Creators container for Creators in SeriesResults in SeriesResponse from Marvel API call
type CreatorsResponse struct {
	Available     int       `json:"available"`
	CollectionURI string    `json:"collectionURI"`
	Items         []Creator `json:"items"`
	Returned      int       `json:"returned"`
}

// Represents the Creators in the Items array in SeriesResults in SeriesResponse from Marvel API call
type Creator struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

// Represents the Characters in the Items array in SeriesResults in SeriesResponse from Marvel API call
type Character struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Represents the Comics in the Items array in SeriesResults in SeriesResponse from Marvel API call
type ComicBook struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Represents the tables
type SeriesDbEntry struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type ComicBookDbEntry struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	SeriesId    int    `db:"seriesId"`
}

type CharacterDbEntry struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	SeriesId    int    `db:"seriesId"`
	ComicId     int    `db:"comicId"`
}

type DistinctCharacterDbEntry struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

// Use viper to get an environment variable
func getEnvironmentVariable(key string) string {

	// load .env file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

// Set any global variables from environment variables
func setGlobalVars() {
	privateAPIKey = getEnvironmentVariable("PRIVATE_API_KEY")
	publicAPIKey = getEnvironmentVariable("PUBLIC_API_KEY")

}

// Hashes a string
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Get and populate the SeriesResponse struct from the Marvel API
func getSeriesResponse(hash string, nowString string) SeriesResponse {
	var seriesCall string = fmt.Sprintf("http://gateway.marvel.com/v1/public/series?limit=%d&ts=%s&apikey=%s&hash=%s", seriesLimit, nowString, publicAPIKey, hash)

	res, err := http.Get(seriesCall)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	var response SeriesResponse

	responseBody, _ := ioutil.ReadAll(res.Body)

	if err = json.Unmarshal(responseBody, &response); err != nil {
		log.Fatal(err)
	}

	return response
}

// Get and populate the SeriesComicsResponse struct from the Marvel API
func getSeriesComicsResponses(seriesResponse SeriesResponse, hash string, now string) []SeriesComicsResponse {
	var seriesComicsresponses []SeriesComicsResponse

	for _, results := range seriesResponse.Data.Results {
		seriesComicsCall := fmt.Sprintf("https://gateway.marvel.com:443/v1/public/series/%d/comics?ts=%s&apikey=%s&hash=%s", results.Id, now, publicAPIKey, hash)

		res, err := http.Get(seriesComicsCall)

		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}

		var response SeriesComicsResponse

		responseBody, _ := ioutil.ReadAll(res.Body)

		if err = json.Unmarshal(responseBody, &response); err != nil {
			log.Fatal(err)

		}

		response.SeriesId = results.Id

		seriesComicsresponses = append(seriesComicsresponses, response)

	}

	return seriesComicsresponses

}

// Get and populate the SeriesCharactersResponse struct from the Marvel API
func getSeriesCharactersResponses(seriesComicsResponses []SeriesComicsResponse, hash string, now string) []SeriesCharactersResponse {
	var seriesCharactersresponses []SeriesCharactersResponse

	for _, response := range seriesComicsResponses {
		seriesId := response.SeriesId
		for _, results := range response.Data.Results {

			seriesCharactersCall := fmt.Sprintf("https://gateway.marvel.com:443/v1/public/series/%d/characters?ts=%s&apikey=%s&hash=%s", seriesId, now, publicAPIKey, hash)

			res, err := http.Get(seriesCharactersCall)

			if err != nil {
				fmt.Printf("error making http request: %s\n", err)
				os.Exit(1)
			}

			var charResponse SeriesCharactersResponse

			responseBody, _ := ioutil.ReadAll(res.Body)

			if err = json.Unmarshal(responseBody, &charResponse); err != nil {
				log.Fatal(err)

			}

			charResponse.SeriesId = seriesId
			charResponse.ComicId = results.Id

			seriesCharactersresponses = append(seriesCharactersresponses, charResponse)
		}
	}

	return seriesCharactersresponses

}

// Populate the SeriesDbEntry structs that will be stored in DB
func populateSeriesDbEntries(seriesResponse SeriesResponse) []SeriesDbEntry {
	var seriesDbEntries []SeriesDbEntry

	for _, series := range seriesResponse.Data.Results {
		var s SeriesDbEntry

		s.Id = series.Id
		s.Title = series.Title
		s.Description = series.Description

		seriesDbEntries = append(seriesDbEntries, s)
	}

	return seriesDbEntries
}

// Populate the ComicsDbEntry structs that will be stored in DB
func populateComicsDbEntries(seriesComicsResponse []SeriesComicsResponse) []ComicBookDbEntry {
	var ComicBookDbEntries []ComicBookDbEntry

	for _, response := range seriesComicsResponse {
		seriesId := response.SeriesId

		for _, comic := range response.Data.Results {
			var c ComicBookDbEntry

			c.Id = comic.Id
			c.Title = comic.Title
			c.Description = comic.Description
			c.SeriesId = seriesId

			ComicBookDbEntries = append(ComicBookDbEntries, c)
		}
	}

	return ComicBookDbEntries
}

// Populate the CharactersDbEntry structs that will be stored in DB
func populateCharactersDbEntries(seriesCharactersResponses []SeriesCharactersResponse) []CharacterDbEntry {
	var charactersDbEntries []CharacterDbEntry

	for _, response := range seriesCharactersResponses {
		seriesId := response.SeriesId
		comicId := response.ComicId
		for _, character := range response.Data.Results {
			var c CharacterDbEntry

			c.Id = character.Id
			c.Name = character.Name
			c.Description = character.Description
			c.SeriesId = seriesId
			c.ComicId = comicId

			charactersDbEntries = append(charactersDbEntries, c)
		}
	}

	return charactersDbEntries
}

func populateDistinctCharactersDbEntries(seriesCharactersResponses []SeriesCharactersResponse) []DistinctCharacterDbEntry {
	var distinctCharacterDbEntries []DistinctCharacterDbEntry
	//distinctCharacterIds := []
	for _, response := range seriesCharactersResponses {
		for _, character := range response.Data.Results {
			var c DistinctCharacterDbEntry

			c.Id = character.Id
			c.Name = character.Name
			c.Description = character.Description

			idPresent := false
			for _, i := range distinctCharacterDbEntries {
				if i.Id == c.Id {
					idPresent = true
					break
				}
			}
			if idPresent == false {
				// Add to the struct
				distinctCharacterDbEntries = append(distinctCharacterDbEntries, c)
			}
		}
	}

	return distinctCharacterDbEntries
}

func ExecuteMigration() {
	// Set the location of the environment variables file
	viper.SetConfigFile("env-vars.env")

	// Set globals using environment variables with viper
	setGlobalVars()

	now := time.Now().Unix()
	nowString := strconv.FormatInt(now, 16)
	hash_this := nowString + privateAPIKey + publicAPIKey

	// hash needed to make Marvel API call
	hash := GetMD5Hash(hash_this)

	// Get the data from Marvel API and store in response structs
	seriesResponse := getSeriesResponse(hash, nowString)
	seriesComicsResponses := getSeriesComicsResponses(seriesResponse, hash, nowString)
	seriesCharactersResponses := getSeriesCharactersResponses(seriesComicsResponses, hash, nowString)

	// Build out the db entry structs to be stored in db
	seriesDbEntries := populateSeriesDbEntries(seriesResponse)
	comicsDbEntries := populateComicsDbEntries(seriesComicsResponses)
	charactersDbEntries := populateCharactersDbEntries(seriesCharactersResponses)
	distinctCharacterDbEntries := populateDistinctCharactersDbEntries(seriesCharactersResponses)

	DbConnect()
	seriesInsertQuery := `INSERT INTO series(id, title, description)
	    VALUES(:id, :title, :description)`
	insertStruct(seriesDbEntries, seriesInsertQuery)

	comicsInsertQuery := `INSERT INTO comic_books(id, title, description, series_id)
	     VALUES(:id, :title, :description, :seriesId)`
	insertStruct(comicsDbEntries, comicsInsertQuery)

	charactersInsertQuery := `INSERT INTO characters(id, name, description)
	     VALUES(:id, :name, :description)`
	insertStruct(distinctCharacterDbEntries, charactersInsertQuery)

	charactersComicBooksInsertQuery := `INSERT INTO characters_comic_books(character_id, comic_book_id)
	     VALUES(:id, :comicId)`
	insertStruct(charactersDbEntries, charactersComicBooksInsertQuery)

	charactersSeriesInsertQuery := `INSERT INTO characters_series(character_id, series_id)
	     VALUES(:id, :seriesId)`
	insertStruct(charactersDbEntries, charactersSeriesInsertQuery)

}

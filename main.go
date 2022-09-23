package main

import (
	"bench/textsearch/database"
	"bench/textsearch/tables"
	"bench/textsearch/utils"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var port int = 10000

// User used for authentication
type User struct {
	Name     string `db:"name"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

// The request body for search requests
type SearchBody struct {
	SearchTable    string   `json:"searchTable"`
	ResultsColumns []string `json:"resultsColumns"`
	SearchTerm     string   `json:"searchTerm"`
	SearchColumn   string   `json:"searchColumn"`
}

// Helper function to encrypt passwords
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// Helper function to check if password is correct
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

// Register user in the database
func RegisterUser(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Db.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	// context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

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

func homePage(context *gin.Context) {
	fmt.Println("Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllSeries(context *gin.Context) {
	fmt.Println("Retrieving all Series")

	var series []tables.Series

	res := database.Db.Find(&series)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": series})
}

func returnComicBooks(context *gin.Context) {
	fmt.Println("Retrieving all Comic Books")

	var comicBooks []tables.ComicBook

	res := database.Db.Find(&comicBooks)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": comicBooks})
}

func returnCharacters(context *gin.Context) {
	fmt.Println("Retrieving all Characters")

	var characters_results []tables.Character

	res := database.Db.Find(&characters_results)

	if res.Error != nil {
		fmt.Printf("Error: %s", res.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": characters_results})
}

func returnSeries(context *gin.Context) {
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

func returnComicBook(context *gin.Context) {
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

func returnCharacter(context *gin.Context) {
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

func searchDatabase(context *gin.Context) {
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

// func getLargestComicBookId() int {
// 	currentId := 0

// 	for _, comicBook := range ComicBooks {

// 		if comicBook.Id > currentId {
// 			currentId = comicBook.Id
// 		}
// 	}

// 	return currentId

// }

// func createComicBook(context *gin.Context) {
// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var comicBook ComicBook

// 	currentId := getLargestComicBookId() + 1

// 	comicBook.Id = currentId

// 	json.Unmarshal(reqBody, &comicBook)

// 	fmt.Printf("Creating Comic Book with id: %d\n", currentId)

// 	ComicBooks = append(ComicBooks, comicBook)

// 	json.NewEncoder(context.Writer).Encode(comicBook)
// }

// func deleteComicBook(context *gin.Context) {
// 	vars := mux.Vars(r)
// 	key, _ := strconv.Atoi(vars["id"])

// 	fmt.Printf("Removing Comic Book with id: %d\n", key)

// 	for index, article := range ComicBooks {
// 		if article.Id == key {
// 			ComicBooks = append(ComicBooks[:index], ComicBooks[index+1:]...)
// 		}
// 	}
// }

// func updateComicBook(context *gin.Context) {
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

func initRouter() *gin.Engine {
	router := gin.Default()
	series := router.Group("/series")
	{
		series.GET("/", returnAllSeries)
		series.GET("/:id", returnSeries)
	}
	comicBooks := router.Group("/comic-books")
	{
		comicBooks.GET("/", returnComicBooks)
		comicBooks.GET("/:id", returnComicBook)
	}
	characters := router.Group("/characters")
	{
		characters.GET("/", returnCharacters)
		characters.GET("/:id", returnCharacter)
	}
	search := router.Group("/search")
	{
		search.POST("/", searchDatabase)
	}

	// api := router.Group("/api")
	// {
	// 	api.POST("/token", controllers.GenerateToken)
	// 	api.POST("/user/register", controllers.RegisterUser)
	// 	secured := api.Group("/secured").Use(middlewares.Auth())
	// 	{
	// 		secured.GET("/ping", controllers.Ping)
	// 	}
	// }
	return router
}

func main() {
	// get the executeDataMigration parameter from the cmd line
	executeDataMigrationPtr := flag.Bool("executeDataMigration", false, "bool value")
	executeBuildTables := flag.Bool("executeBuildTables", false, "bool value")
	flag.Parse()
	viper.SetConfigFile("env-vars.env")

	if *executeBuildTables {
		database.Connect()
		database.CreateTables()
	}

	if *executeDataMigrationPtr {
		database.Connect()
		utils.ExecuteMigration()
	}

	if !*executeDataMigrationPtr && !*executeBuildTables {
		database.Connect()
		router := initRouter()
		router.Run(fmt.Sprintf(":%d", port))
	}

}

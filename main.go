package main

import (
	"bench/textsearch/controllers"
	"bench/textsearch/database"
	"bench/textsearch/middlewares"
	"bench/textsearch/utils"
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var port int = 10000

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

// test call
func ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", ping)
		}
	}
	series := router.Group("/series").Use(middlewares.Auth())
	{
		series.GET("/", controllers.ReturnAllSeries)
		series.GET("/:id", controllers.ReturnSeries)
	}
	comicBooks := router.Group("/comic-books").Use(middlewares.Auth())
	{
		comicBooks.GET("/", controllers.ReturnComicBooks)
		comicBooks.GET("/:id", controllers.ReturnComicBook)
	}
	characters := router.Group("/characters").Use(middlewares.Auth())
	{
		characters.GET("/", controllers.ReturnCharacters)
		characters.GET("/:id", controllers.ReturnCharacter)
	}
	search := router.Group("/search").Use(middlewares.Auth())
	{
		search.POST("/", controllers.SearchDatabase)
	}
	return router
}

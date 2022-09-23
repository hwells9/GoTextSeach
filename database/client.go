package database

import (
	"bench/textsearch/authentication"
	"bench/textsearch/tables"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var dbError error

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

func Connect() {
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

	Db, dbError = gorm.Open(postgres.Open(postgresConnectionString), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func CreateTables() {
	err := Db.AutoMigrate(&authentication.User{})
	{
		if err != nil {
			log.Panicf("Error creating User table with error: %s", err)
		}
	}
	err1 := Db.AutoMigrate(&tables.Series{})
	{
		if err1 != nil {
			log.Panicf("Error creating Series table with error: %s", err1)
		}
	}
	err2 := Db.AutoMigrate(&tables.ComicBook{})
	{
		if err2 != nil {
			log.Panicf("Error creating ComicBook table with error: %s", err2)
		}
	}
	err3 := Db.AutoMigrate(&tables.Character{})
	{
		if err3 != nil {
			log.Panicf("Error creating Character table with error: %s", err3)
		}
	}
	err4 := Db.AutoMigrate(&tables.CharacterSeries{})
	{
		if err4 != nil {
			log.Panicf("Error creating CharacterSeries table with error: %s", err4)
		}
	}
	err5 := Db.AutoMigrate(&tables.CharacterComicBook{})
	{
		if err5 != nil {
			log.Panicf("Error creating CharacterComicBook table with error: %s", err5)
		}
	}
	log.Println("tables have been created!")
}

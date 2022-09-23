package tables

import (
	"gorm.io/gorm"
)

// Represents the Series models
type Series struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Represents Comic Book table
type ComicBook struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SeriesId    int    `json:"series_id"`
	Series      Series `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Represents Characters Table
type Character struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Represents Character Comic Book Table
type CharacterComicBook struct {
	gorm.Model
	CharacterId int       `json:"character_id"`
	ComicBookId int       `json:"comic_book_id"`
	Character   Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ComicBook   ComicBook `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Represents Character Series Table
type CharacterSeries struct {
	gorm.Model
	CharacterId int       `json:"character_id"`
	SeriesId    int       `json:"series_id"`
	Character   Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Series      Series    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

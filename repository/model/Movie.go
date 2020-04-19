package model

import "github.com/jinzhu/gorm"

// Movie is Gorm model of movie
type Movie struct {
	gorm.Model
	DirPath    string `json:"dirPath"`
	Title      string `json:"title"`
	IsPrepared bool   `json:"isPrepared"`
}

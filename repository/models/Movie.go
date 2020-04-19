package models

import "github.com/jinzhu/gorm"

// Movie is Gorm model of movie
type Movie struct {
	gorm.Model
	DirPath    string  `json:"dirPath"`
	DirName    string  `json:"dirName"`
	BaseName   string  `json:"baseName"`
	FileSize   float64 `json:"fileSize"`
	MimeType   string  `json:"mimeType"`
	IsPrepared bool    `json:"isPrepared"`
	IsGroupDir bool    `json:"isGroupDir"`
}

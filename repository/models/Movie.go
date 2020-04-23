package models

import "github.com/jinzhu/gorm"

// Movie is Gorm model of movie
type Movie struct {
	gorm.Model
	DirPath       string  `json:"dirPath"`
	DirName       string  `json:"dirName"`
	CleanDirName  string  `json:"cleanDirName"`
	BaseName      string  `json:"baseName"`
	CleanBaseName string  `json:"cleanBaseName"`
	FileSize      float64 `json:"fileSize"`
	MimeType      string  `json:"mimeType"`
	IsPrepared    bool    `json:"isPrepared"`
	IsInPrepare   bool    `json:"isInPrepare"`
	IsGroupDir    bool    `json:"isGroupDir"`
}

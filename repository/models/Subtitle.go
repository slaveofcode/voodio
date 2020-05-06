package models

import "github.com/jinzhu/gorm"

// Subtitle is Gorm model of subtitle
type Subtitle struct {
	gorm.Model
	DirPath       string `json:"dirPath"`
	DirName       string `json:"dirName"`
	CleanDirName  string `json:"cleanDirName"`
	BaseName      string `json:"baseName"`
	CleanBaseName string `json:"cleanBaseName"`
}

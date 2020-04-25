package config

import "github.com/jinzhu/gorm"

// ServerConfig will hold configuration for server handler to run
type ServerConfig struct {
	AppDir     string
	DB         *gorm.DB
	Port       int
	FFmpegBin  string
	TMDBApiKey string
}

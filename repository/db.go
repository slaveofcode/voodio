package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/voodio/repository/models"

	// dialect for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// OpenDB will create new database connection to Sqlite
func OpenDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Migrate will do migration of models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Movie{})
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/voodio/repository/models"
)

// MoviesPage will return function to handle Movie list
func MoviesPage(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var nonGroupMovies []models.Movie
		db.Where("is_group_dir = ?", false).Find(&nonGroupMovies)

		var groupMovies []models.Movie
		db.Where(&models.Movie{
			IsGroupDir: true,
		}).Group("dir_name").Find(&groupMovies)

		var allMovies []models.Movie

		allMovies = append(allMovies, nonGroupMovies...)
		allMovies = append(allMovies, groupMovies...)

		json, _ := json.Marshal(map[string]interface{}{
			"movies": allMovies,
			"count":  len(allMovies),
		})

		w.Write(json)

		return
	})
}

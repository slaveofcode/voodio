package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/voodio/repository/models"
)

// GroupMoviesPage will return function to handle Movie Grouplist
func GroupMoviesPage(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		movieID := strings.TrimSpace(r.URL.Query().Get("movieId"))

		var movie models.Movie
		if db.Where("id = ?", movieID).First(&movie).RecordNotFound() {
			json, _ := json.Marshal(map[string]interface{}{
				"found": false,
			})
			w.Write(json)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var groupMovies []models.Movie
		db.Where(&models.Movie{
			IsGroupDir: true,
			DirName:    movie.DirName,
			DirPath:    movie.DirPath,
		}).Find(&groupMovies)

		json, _ := json.Marshal(map[string]interface{}{
			"movies": groupMovies,
			"count":  len(groupMovies),
		})

		w.Write(json)

		return
	})
}

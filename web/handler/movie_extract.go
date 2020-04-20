package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/voodio/collections"
	"github.com/slaveofcode/voodio/repository/models"
	"github.com/slaveofcode/voodio/web/config"
)

// MovieExtractHLS will return function to handle extraction trigger of movie
func MovieExtractHLS(cfg *config.ServerConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get movie id
		movieId := strings.TrimSpace(r.URL.Query().Get("movieId"))

		if movieId == "" {
			w.Header().Set("content-type", "application/json")
			json, _ := json.Marshal(map[string]interface{}{
				"processed": false,
			})
			w.Write(json)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var movie models.Movie
		if cfg.DB.Where("id = ?", movieId).First(&movie).RecordNotFound() {
			w.Header().Set("content-type", "application/json")
			json, _ := json.Marshal(map[string]interface{}{
				"processed": false,
			})
			w.Write(json)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if movie.IsInPrepare || movie.IsPrepared {
			w.Header().Set("content-type", "application/json")
			json, _ := json.Marshal(map[string]interface{}{
				"processed": true,
			})
			w.Write(json)
			w.WriteHeader(http.StatusOK)
			return
		}

		// spawn worker to do extraction
		go func(db *gorm.DB, mov models.Movie) {
			db.Model(&mov).Update(&models.Movie{
				IsInPrepare: true,
			})

			// extract
			err := collections.DoExtraction(&mov, cfg.AppDir)

			if err != nil {
				logrus.Errorln("Something wrong when extracting HLS file", err)
				db.Model(&mov).Update(&models.Movie{
					IsInPrepare: false,
					IsPrepared:  false,
				})
				return
			}

			db.Model(&mov).Update(&models.Movie{
				IsInPrepare: false,
				IsPrepared:  true,
			})

			return
		}(cfg.DB, movie)

		w.Header().Set("content-type", "application/json")
		json, _ := json.Marshal(map[string]interface{}{
			"processed": true,
		})
		w.Write(json)
		w.WriteHeader(http.StatusOK)
		return

	})
}

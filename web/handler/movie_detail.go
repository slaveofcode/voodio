package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/slaveofcode/voodio/repository/models"
	"github.com/slaveofcode/voodio/web/config"
)

// MovieDetail will return function to handle movie detail request
func MovieDetail(cfg *config.ServerConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		movieID := strings.TrimSpace(r.URL.Query().Get("movieId"))

		var movie models.Movie
		if cfg.DB.Where("id = ?", movieID).First(&movie).RecordNotFound() {
			w.Header().Set("content-type", "application/json")
			json, _ := json.Marshal(map[string]interface{}{
				"found": false,
			})
			w.Write(json)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("content-type", "application/json")
		json, _ := json.Marshal(movie)
		w.Write(json)
		w.WriteHeader(http.StatusNotFound)
		return
	})
}

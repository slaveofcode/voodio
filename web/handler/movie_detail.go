package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/voodio/collections"
	"github.com/slaveofcode/voodio/repository/models"
	"github.com/slaveofcode/voodio/web/config"
)

type subinfo struct {
	SubName string `json:"subName"`
	SubFile string `json:"subFile"`
}

type movieResp struct {
	models.Movie
	Subtitles []subinfo
}

func getSubMovieInfos(db *gorm.DB, movie *models.Movie) []subinfo {
	var subs []models.Subtitle
	db.Where(&models.Subtitle{
		DirPath: movie.DirPath,
	}).Find(&subs)

	var subInfos []subinfo
	for _, s := range subs {
		subInfos = append(subInfos, subinfo{
			SubName: strings.Split(s.BaseName, ".")[0],
			SubFile: collections.GetVTTFileName(s.BaseName),
		})
	}

	return subInfos
}

// MovieDetail will return function to handle movie detail request
func MovieDetail(cfg *config.ServerConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		movieID := strings.TrimSpace(r.URL.Query().Get("movieId"))

		var movie models.Movie
		if cfg.DB.Where("id = ?", movieID).First(&movie).RecordNotFound() {

			json, _ := json.Marshal(map[string]interface{}{
				"found": false,
			})
			w.Write(json)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json, _ := json.Marshal(movieResp{
			Movie:     movie,
			Subtitles: getSubMovieInfos(cfg.DB, &movie),
		})
		w.Write(json)
		w.WriteHeader(http.StatusNotFound)
		return
	})
}

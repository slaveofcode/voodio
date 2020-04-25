package web

import (
	"net/http"
	"path/filepath"

	"github.com/rakyll/statik/fs"
	"github.com/slaveofcode/voodio/collections"

	// Import statik generated file
	_ "github.com/slaveofcode/voodio/statik"
	"github.com/slaveofcode/voodio/web/config"
	"github.com/slaveofcode/voodio/web/handler"
)

// NewRouter will return new router
func NewRouter(cfg *config.ServerConfig) *http.ServeMux {
	statikFS, err := fs.New()
	if err != nil {
		panic("Unable to spawn Web UI" + err.Error())
	}

	router := http.NewServeMux()
	router.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))

	router.Handle("/tmdb", CorsMiddleware(JSONMiddleware(handler.TMDBHandler(cfg))))

	router.Handle("/movies", CorsMiddleware(JSONMiddleware(handler.MoviesPage(cfg.DB))))
	router.Handle("/movies/prepare", CorsMiddleware(JSONMiddleware(handler.MovieExtractHLS(cfg))))
	router.Handle("/movies/detail", CorsMiddleware(JSONMiddleware(handler.MovieDetail(cfg))))

	staticPath := filepath.Join(cfg.AppDir, collections.ExtractionDirName)
	staticDirHandler := http.Dir(staticPath)
	router.Handle("/hls/", CorsMiddleware(http.StripPrefix("/hls/", http.FileServer(staticDirHandler))))

	return router
}

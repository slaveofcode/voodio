package web

import (
	"net/http"
	"path/filepath"

	"github.com/slaveofcode/voodio/collections"
	"github.com/slaveofcode/voodio/web/config"
	"github.com/slaveofcode/voodio/web/handler"
)

// NewRouter will return new router
func NewRouter(cfg *config.ServerConfig) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", CorsMiddleware(JSONMiddleware(handler.IndexPage())))
	router.Handle("/movies", CorsMiddleware(JSONMiddleware(handler.MoviesPage(cfg.DB))))
	router.Handle("/movies/prepare", CorsMiddleware(JSONMiddleware(handler.MovieExtractHLS(cfg))))
	router.Handle("/movies/detail", CorsMiddleware(JSONMiddleware(handler.MovieDetail(cfg))))

	staticPath := filepath.Join(cfg.AppDir, collections.ExtractionDirName)
	staticDirHandler := http.Dir(staticPath)
	router.Handle("/hls/", CorsMiddleware(http.StripPrefix("/hls/", http.FileServer(staticDirHandler))))

	return router
}

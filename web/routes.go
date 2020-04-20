package web

import (
	"net/http"

	"github.com/slaveofcode/pms/web/config"
	"github.com/slaveofcode/pms/web/handler"
)

// NewRouter will return new router
func NewRouter(cfg *config.ServerConfig) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", handler.IndexPage())
	router.Handle("/movies", handler.MoviesPage(cfg.DB))
	router.Handle("/movies/prepare", handler.MovieExtractHLS(cfg))
	router.Handle("/movies/detail", handler.MovieDetail(cfg))

	return router
}

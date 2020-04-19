package web

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/pms/web/handler"
)

// NewRouter will return new router
func NewRouter(db *gorm.DB) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", handler.IndexPage())
	router.Handle("/movies", handler.MoviesPage(db))

	return router
}

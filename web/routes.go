package web

import (
	"net/http"

	"github.com/slaveofcode/pms/web/handler"
)

// NewRouter will return new router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", handler.IndexPage())

	return router
}

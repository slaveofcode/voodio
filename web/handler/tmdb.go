package handler

import (
	"encoding/json"
	"net/http"

	"github.com/slaveofcode/voodio/web/config"
)

// TMDBHandler will return TMDB api handler
func TMDBHandler(cfg *config.ServerConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json, _ := json.Marshal(map[string]interface{}{
			"key": cfg.TMDBApiKey,
		})
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	})
}

package handler

import (
	"encoding/json"
	"net/http"
)

// IndexPage will return HandlerFunc for homepage
func IndexPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp, _ := json.Marshal(map[string]interface{}{
			"status": "OK",
		})
		w.Write(resp)
		return
	})
}

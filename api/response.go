package api

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func parseLimit(r *http.Request) int {
	n := 20
	if r != nil && r.URL.Query().Get("limit") == "100" {
		n = 100
	}
	return n
}

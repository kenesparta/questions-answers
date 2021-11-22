package infra

import "net/http"

func Headers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

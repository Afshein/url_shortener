package main

import (
	// "myproject/internal/storage"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		break
	case http.MethodPost:
		break
	default:
		w.Header().Set("Allow", "POST, GET")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

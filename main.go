package main

import (
	"crypto/sha256"
	"fmt"
	"myproject/storage"
)

// func main() {
// 	http.HandleFunc("/urlshort", func(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		break
// 	case http.MethodPost:
// 		break
// 	default:
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"status":"ok"}`))
// 	})
// }

func main() {
	hash := storage.Url_set("test")
	url := storage.Url_get(hash)
	fmt.Printf(url)
}

func shorten_url(url string) {
	s := "test"
	// hash the url
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	// hash the endpoint

	fmt.Printf("%x\n", bs)
}

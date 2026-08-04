package main

import (
	"log"
	"myproject/internal/storage"
	"net/http"
	"path"
)

func shortURLGet(w http.ResponseWriter, r *http.Request) {
	key := path.Base(r.URL.Path)
	longURL, err := storage.Url_get(key)

	log.Print("hit GET")

	if err != nil {
		http.Error(w, "Short URL does not exist", http.StatusBadRequest)
	}

	log.Printf("Redirecting to %s", longURL)
	http.Redirect(w, r, longURL, http.StatusFound)

}

func longURLPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to Parse Request", http.StatusBadRequest)
	}

	log.Print("hit POST")

	longURL := r.PostForm.Get("url")

	if longURL == "" {
		http.Error(w, "Must pass non empty string for url param", http.StatusBadRequest)
		return
	}

	shortUrl := storage.Url_set(longURL)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(
		`{
        "status": "ok",
        "shortUrl": "` + shortUrl + `"
    }`,
	))
}

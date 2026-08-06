package main

import (
	"log"
	"net/http"
	"path"
)

func (app *App) shortURLGet(w http.ResponseWriter, r *http.Request) {
	log.Print("hit GET")

	key := path.Base(r.URL.Path)
	longURL, err := app.urlMapping.GetLongURL(app.context, key)

	if err != nil {
		http.Error(w, "Short URL does not exist", http.StatusBadRequest)
		return
	}

	log.Printf("Redirecting to %s", longURL)
	http.Redirect(w, r, longURL, http.StatusFound)

}

func (app *App) longURLPost(w http.ResponseWriter, r *http.Request) {
	log.Print("hit POST")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to Parse Request", http.StatusBadRequest)
	}

	longURL := r.PostForm.Get("url")

	if longURL == "" {
		http.Error(w, "Must pass non empty string for url param", http.StatusBadRequest)
		return
	}

	shortUrl, err := app.urlMapping.CacheLongURL(app.context, longURL)

	if err != nil {
		log.Printf("Failed to cache long url %s", err)
		http.Error(w, "Failed to get shortURL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(
		`{
        "status": "ok",
        "shortUrl": "` + shortUrl + `"
    }`,
	))
}

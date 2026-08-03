package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/shorten", shortURLGet)
	router.HandlerFunc(http.MethodPost, "/", longURLPost)

	// Custom fallback
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		// Walk backwards until a registered route exists
		for len(path) > 1 {
			log.Print(path)
			path = path[:strings.LastIndex(path, "/")]
			if h, _, _ := router.Lookup(req.Method, path); h != nil {
				// Found a parent route → call it
				h(w, req, nil)
				return
			}
		}
		log.Print(req.URL.Path)
		http.NotFound(w, req)
	})

	return router
}

package main

import (
	"flag"
	"log"
	"myproject/internal/config"
	"net/http"
)

func main() {
	cfg := config.Load_config()
	flag.Parse()

	log.Printf("Starting server on %s", *cfg.PORT)
	err := http.ListenAndServe(
		*cfg.PORT,
		routes(),
	)
	log.Fatal(err)
}

package main

import (
	"context"
	"flag"
	"log"
	"myproject/internal/config"
	"myproject/internal/model"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	context    context.Context
	urlMapping model.URLMappingModel
}

func main() {
	cfg := config.Load_config()
	flag.Parse()

	app := App{
		context: context.Background(),
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	if err := rdb.Ping(app.context).Err(); err != nil {
		panic("Redis not reachable: " + err.Error())
	}

	log.Printf("Starting server on %s", *cfg.PORT)
	err := http.ListenAndServe(
		*cfg.PORT,
		app.routes(),
	)
	log.Fatal(err)
}

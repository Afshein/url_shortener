package main

import (
	"context"
	"flag"
	"log"
	"myproject/internal/config"
	"myproject/internal/model"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	context    context.Context
	urlMapping model.URLMappingRepo
}

func main() {
	cfg := config.Load_config()
	flag.Parse()

	context := context.Background()

	log.Printf("DATABASE_URL: %s", os.Getenv("DATABASE_URL"))
	pool, err := pgxpool.New(context, os.Getenv("DATABASE_URL"))

	if err != nil {
		panic("Postgres not reachable: " + err.Error())
	}

	rdb, err := model.OpenCacheDB(context)

	if err != nil {
		panic("Redis not reachable: " + err.Error())
	}

	app := App{
		context: context,
		urlMapping: model.URLMappingRepo{
			RDB: rdb,
			PDB: pool,
		},
	}

	log.Printf("Starting server on %s", *cfg.PORT)
	err = http.ListenAndServe(
		*cfg.PORT,
		app.routes(),
	)
	log.Fatal(err)
}

package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type URLMappingRepo struct {
	RDB *redis.Client
	PDB *pgxpool.Pool
}

func (m *URLMappingRepo) StoreLongURL(ctx context.Context, urlMap URLMapping) error {
	_, err := m.PDB.Exec(
		ctx,
		fmt.Sprintf(
			`INSERT INTO 
			 url_mappings(code, longURL)
		 	 VALUES ('%s', '%s')`,
			urlMap.ShortURL, urlMap.LongURL),
	)

	return err
}

func (m *URLMappingRepo) CacheLongURL(ctx context.Context, urlMap URLMapping) error {
	err := m.RDB.Set(ctx, urlMap.ShortURL, urlMap.LongURL, 0).Err()
	return err
}

func (m *URLMappingRepo) GetLongURL(ctx context.Context, code string) (string, error) {
	longURL, err := m.RDB.Get(ctx, code).Result()

	if err != nil {
		row := m.PDB.QueryRow(
			ctx,
			fmt.Sprintf(
				`SELECT longURL
				FROM url_mappings
				WHERE code = '%s'`,
				code),
		)

		err = row.Scan(&longURL)
	}

	return longURL, err
}

func OpenCacheDB(context context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	err := rdb.Ping(context).Err()

	return rdb, err
}

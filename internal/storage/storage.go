package storage

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password
	DB:       0,  // use default DB
	Protocol: 2,
})

func Url_set(url string) string {
	key := hash(url)
	rdb.Set(ctx, key, url, 0).Err()
	return key
}

func Url_get(hash string) (string, error) {
	url_val, err := rdb.Get(ctx, hash).Result()
	return url_val, err
}

func hash(val string) string {
	// hash the url
	h := sha256.New()
	h.Write([]byte(val))
	bs := h.Sum(nil)

	// Cast to uint64
	hash_int := binary.BigEndian.Uint64(bs)
	// Cast to string
	hash := strconv.FormatUint(hash_int, 10)

	log.Print(hash)

	return hash
}

package model

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type URLMapping struct {
	longURL  string
	shortURL string
}

type URLMappingModel struct {
	rdb *redis.Client
}

func (m *URLMappingModel) CacheLongURL(ctx context.Context, longURL string) (string, error) {
	key := hash(longURL)
	err := m.rdb.Set(ctx, key, longURL, 0).Err()
	return key, err
}

func (m *URLMappingModel) GetLongURL(ctx context.Context, hash string) (string, error) {
	longURL, err := m.rdb.Get(ctx, hash).Result()
	return longURL, err
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

	return hash
}

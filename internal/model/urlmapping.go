package model

import (
	"crypto/sha256"
	"encoding/binary"
	"strconv"
)

type URLMapping struct {
	LongURL  string
	ShortURL string
}

func NewURLMapping(longURL string) URLMapping {
	return URLMapping{
		LongURL:  longURL,
		ShortURL: hash(longURL),
	}
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

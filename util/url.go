package util

import (
	"math/rand"
	"time"
)

const (
	CHARS  = "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	LENGTH = 6
)

func GenShortUrl(longUrl string) {
}

func GeneralShortgUrl(longUrl string) string {
	var shortUrl string
	urlCharsLen := len(CHARS)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < LENGTH; i++ {
		randNumber := r.Intn(urlCharsLen)
		shortUrl += CHARS[randNumber : randNumber+1]
	}
	return shortUrl
}

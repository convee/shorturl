package cache

import (
	"fmt"

	credis "github.com/convee/goboot/db/redis"
)

const (
	SHORT_PRE = "shorturl_"
)

func SetShortUrlCache(shorturl string, longurl string) {
	credis.New("default")
	credis.Set(fmt.Sprintf("%s%s", SHORT_PRE, shorturl), longurl)
}

func GetLongurlByShorturl(shorturl string) string {
	credis.New("default")
	return credis.Get(fmt.Sprintf("%s%s", SHORT_PRE, shorturl))
}

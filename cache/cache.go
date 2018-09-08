package cache

import (
	"fmt"

	myredis "github.com/convee/goboot/db/redis"
)

const (
	SHORT_PRE = "shorturl_"
	LONG_PRE  = "longurl_"
)

func SetUrl(id string, longurl string) {
	myredis.New("default")
	myredis.Set(fmt.Sprintf("%s%s", SHORT_PRE, id), longurl)
}

func GetUrl(id string) (string, error) {
	myredis.New("default")
	str, err := myredis.Get(fmt.Sprintf("%s%s", SHORT_PRE, id))
	if err == nil {
		return str, nil
	}
	return "", err

}

func GetLongurl(token string) (string, error) {
	myredis.New("default")
	str, err := myredis.Get(fmt.Sprintf("%s%s", LONG_PRE, token))
	if err == nil {
		return str, nil
	}
	return "", err
}

func SetLongurl(token string, shorturl string) {
	myredis.New("default")
	myredis.Set(fmt.Sprintf("%s%s", LONG_PRE, token), shorturl)
}

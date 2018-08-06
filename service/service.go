package service

import (
	"github.com/convee/shorturl/cache"
	"github.com/convee/shorturl/mysql"
)

//添加一个短网址
func AddOneShortUrl(shorturl, longurl string) {

}

//查询一个短网址
func FindOneShortUrl(shorturl string) string {
	shorturlCache := cache.GetLongurlByShorturl(shorturl)
	if shorturlCache != "" {
		return shorturlCache
	}
	longurl := mysql.NewModel().GetAllShorturl(shorturl).LongUrl
	if longurl != "" {
		cache.SetShortUrlCache(shorturl, longurl)
		return longurl
	}
	return ""

}

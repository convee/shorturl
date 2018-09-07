package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/convee/goboot"
	"github.com/convee/shorturl/cache"
	"github.com/convee/shorturl/mysql"
	"github.com/convee/shorturl/util"
)

var (
	addr string
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1:8001", "http server")
	flag.Parse()
	goboot.Run("config.toml")
	cache.SetShortUrlCache("abc", "https://convee.cn")
	DecimalTo62()
	startHTTPServer(addr)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func startHTTPServer(addr string) {
	fmt.Println("http server starting on ", addr)
	http.HandleFunc("/", index)
	http.HandleFunc("/genShorturl", genShorturl)
	http.HandleFunc("/getLongurl", getLongurl)
	http.ListenAndServe(addr, nil)
}

//长网址转换成短网址
func genShorturl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	longurl := r.Form["longurl"][0]
	shorturl := util.GeneralShortgUrl(longurl)
	fmt.Println(shorturl)
}

//短网址获取长网址
func getLongurl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	shorturl := r.Form["shorturl"][0]
	var longurl string
	longurlCache, err := cache.GetLongurlByShorturl(shorturl)
	fmt.Println(longurlCache, err)
	if err == nil {
		longurl = longurlCache
	} else {
		longurlData, err := mysql.NewModel().GetLongurlByShorturl(shorturl)
		fmt.Println(longurlData, err)
		if err == nil {
			longurl = longurlData
			cache.SetShortUrlCache(shorturl, longurl)
		}
	}
	if longurl != "" {
		response := make(map[string]string)
		response["longurl"] = longurl
		response["shorturl"] = shorturl
		util.JsonReturn(w, util.Json{
			Error: 0,
			Msg:   "ok",
			Data:  response,
		})
	} else {
		util.JsonReturn(w, util.Json{
			Error: 1,
			Msg:   "no result",
		})
	}

}

func DecimalTo62() {
	shortUrl := util.DecimalToAny(10000000000, 62)
	fmt.Println(shortUrl)
}

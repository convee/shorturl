package main

import (
	"crypto/md5"
	"encoding/hex"
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
	startHTTPServer(addr)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func startHTTPServer(addr string) {
	fmt.Println("http server starting on ", addr)
	http.HandleFunc("/", index)
	http.HandleFunc("/genUrl", genUrl)
	http.HandleFunc("/getUrl", getUrl)
	http.ListenAndServe(addr, nil)
}

//长网址转换成短网址
func genUrl(w http.ResponseWriter, r *http.Request) {
	var shorturl string
	r.ParseForm()
	longurl := r.Form["longurl"][0]
	h := md5.New()
	h.Write([]byte("abcdefg!@"))
	token := hex.EncodeToString(h.Sum([]byte(longurl)))
	longCache, err := cache.GetLongurl(token)
	if err == nil {
		shorturl = longCache
	} else {
		id, err := mysql.NewModel().InsertShorturl("", longurl)
		fmt.Println(id)
		if err == nil {
			shorturl := util.DecimalToAny(int(id), 62)
			fmt.Println(shorturl)
			cache.SetLongurl(token, shorturl)
		}
	}
	if shorturl != "" {
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

//短网址获取长网址
func getUrl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	shorturl := r.Form["shorturl"][0]
	id := util.AnyToDecimal(shorturl, 64)
	var longurl string
	longurlCache, err := cache.GetUrl(string(id))
	if err == nil {
		longurl = longurlCache
	} else {
		longurlData, err := mysql.NewModel().GetLongurl(id)
		fmt.Println(longurlData, err)
		if err == nil {
			longurl = longurlData
			cache.SetUrl(string(id), longurl)
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

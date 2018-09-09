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
	url := r.Form["url"][0]
	h := md5.New()
	h.Write([]byte("abcdefg!@"))
	token := hex.EncodeToString(h.Sum([]byte(url)))
	longCache, err := cache.GetLongurl(token)
	if err == nil {
		shorturl = longCache
	} else {
		id, err := mysql.NewModel().AddUrl(url)
		fmt.Println(id)
		if err == nil {
			shorturl = util.DecimalToAny(int(id), 62)
			fmt.Println(shorturl)
			cache.SetLongurl(token, shorturl)
		}
	}
	fmt.Println(shorturl)
	if shorturl != "" {
		response := make(map[string]string)
		response["url"] = url
		response["short"] = shorturl
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
	short := r.Form["short"][0]
	id := util.AnyToDecimal(short, 62)
	var url string
	urlCache, err := cache.GetUrl(string(id))
	if err == nil {
		url = urlCache
	} else {
		urlData, err := mysql.NewModel().GetUrl(id)
		fmt.Println(urlData, err)
		if err == nil {
			url = urlData
			cache.SetUrl(string(id), url)
		}
	}
	if url != "" {
		response := make(map[string]string)
		response["url"] = url
		response["short"] = short
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

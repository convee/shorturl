package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"net/http"

	"github.com/convee/goboot"
	"github.com/convee/goboot/router"
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

func startHTTPServer(addr string) {
	fmt.Println("http server starting on ", addr)
	var r router.Router
	r.HandleFunc("/genUrl", genUrl).Get()
	r.HandleFunc("/getUrl", getUrl).Get()
	r.Handle("/jump", http.RedirectHandler("http://www.convee.cn", 302))
	r.HandleFunc("/(?P<short>.+)", redirectUrl).Get()
	http.ListenAndServe(addr, r)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("前置操作")
		next.ServeHTTP(w, r)
		fmt.Println("后置操作")
	})
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

func redirectUrl(w http.ResponseWriter, r *http.Request) {
	short := r.FormValue(":short")
	fmt.Println(short)
	id := util.AnyToDecimal(short, 62)
	var url string
	urlCache, err := cache.GetUrl(string(id))
	if err == nil {
		url = urlCache
	} else {
		urlData, err := mysql.NewModel().GetUrl(id)
		if err == nil {
			url = urlData
			cache.SetUrl(string(id), url)
		}
	}
	fmt.Println(url)
	if url != "" {
		http.Redirect(w, r, url, 302)
	} else {
		w.Write([]byte("404 page not found"))
	}
}

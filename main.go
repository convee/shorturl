package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/convee/goboot"
	"github.com/convee/shorturl/cache"
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
	fmt.Println("http server starting....")
	http.HandleFunc("/", index)
	http.HandleFunc("/gen", genShort)
	http.HandleFunc("/jump", jump)
	http.ListenAndServe(addr, nil)
}

//长网址转换成短网址
func genShort(w http.ResponseWriter, r *http.Request) {
	longUrl := "https://www.baidu.com"
	shortUrl := util.GeneralShortgUrl(longUrl)
	fmt.Println(shortUrl)
}

//短网址跳转
func jump(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello jump")
}

func DecimalTo62() {
	shortUrl := util.DecimalToAny(10000000000, 62)
	fmt.Println(shortUrl)
}

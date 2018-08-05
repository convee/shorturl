package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/convee/goboot"
	"github.com/convee/shorturl/app"
	"github.com/convee/shorturl/util"
)

var (
	addr = *flag.String("addr", "127.0.0.1:8001", "http server")
)

func main() {
	flag.Parse()
	goboot.Run("config.toml")
	DecimalTo62()
	startHTTPServer(addr)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func startHTTPServer(addr string) {
	fmt.Println("http server starting....")
	http.HandleFunc("/", index)
	http.ListenAndServe(addr, nil)
}
func GenShort() {
	longUrl := "https://www.baidu.com"
	shortUrl := app.GeneralShortgUrl(longUrl)
	fmt.Println(shortUrl)
}

func DecimalTo62() {
	shortUrl := util.DecimalToAny(10000000000, 62)
	fmt.Println(shortUrl)
}

package main

import (
	"github.com/convee/goboot"
	"github.com/convee/shorturl/app"
)

func main() {
	goboot.Run("config.toml")
	app.NewModel().GetAllShorturl()
	app.NewModel().GetShorturl()
}

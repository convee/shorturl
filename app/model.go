package app

import (
	"database/sql"
	"fmt"

	"github.com/convee/goboot/db/mysql"
)

type Model struct {
	db *sql.DB
}

type Shorturl struct {
	Id       int    `json:"id"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}

func NewModel() *Model {
	return &Model{db: mysql.New("test")}
}

func (m *Model) GetAllShorturl() {
	var shorturl Shorturl
	err := m.db.QueryRow("select id,long_url,short_url from short_url where id=?", 1).Scan(&shorturl.Id, &shorturl.LongUrl, &shorturl.ShortUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(shorturl)
}

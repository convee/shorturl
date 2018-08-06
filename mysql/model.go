package mysql

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

func (m *Model) GetAllShorturl(short string) Shorturl {
	var shorturl Shorturl
	err := m.db.QueryRow("select id,long_url,short_url from short_url where short_url=?", short).Scan(&shorturl.Id, &shorturl.LongUrl, &shorturl.ShortUrl)
	if err != nil {
		panic(err)
	}
	return shorturl
}

func (m *Model) InsertShorturl(shorturl string, longurl string) (id int64) {
	rs, err := m.db.Exec("insert into short_url (long_url, short_url) values (?, ?)", longurl, shorturl)
	if err != nil {
		panic(err)
	}
	id, err = rs.LastInsertId()
	if err != nil {
		panic(err)
	}
	return
}

func (m *Model) GetShorturl() {
	var shorturls []Shorturl
	rows, err := m.db.Query("select * from short_url")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var shorturl Shorturl
		rows.Scan(&shorturl.Id, &shorturl.ShortUrl, &shorturl.LongUrl)
		shorturls = append(shorturls, shorturl)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	fmt.Println(shorturls)
}

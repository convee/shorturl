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
	Shorturl string `json:"short_url"`
	Longurl  string `json:"long_url"`
}

func NewModel() *Model {
	return &Model{db: mysql.New("test")}
}

func (m *Model) GetLongurlByShorturl(short string) (string, error) {
	var longurl string
	err := m.db.QueryRow("SELECT `long_url` FROM short_url WHERE `short_url` = ? limit 1", short).Scan(&longurl)
	if err != nil {
		return "", err
	}
	return longurl, nil
}

func (m *Model) GetLongurl(id int) (string, error) {
	var longurl string
	err := m.db.QueryRow("SELECT `long_url` FROM short_url WHERE `id` = ? limit 1", id).Scan(&longurl)
	if err != nil {
		return "", err
	}
	return longurl, nil
}

func (m *Model) isExists(longurl string) (int, error) {
	var num int
	err := m.db.QueryRow("select count(id) as num from short_url where long_url = ?", longurl).Scan(&num)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (m *Model) InsertShorturl(shorturl string, longurl string) (int64, error) {
	rs, err := m.db.Exec("insert into short_url (long_url, short_url) values (?, ?)", longurl, shorturl)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *Model) GetAll() {
	var shorturls []Shorturl
	rows, err := m.db.Query("select * from short_url")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var shorturl Shorturl
		rows.Scan(&shorturl.Id, &shorturl.Shorturl, &shorturl.Longurl)
		shorturls = append(shorturls, shorturl)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	fmt.Println(shorturls)
}

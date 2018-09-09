package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/convee/goboot/db/mysql"
)

type Model struct {
	db *sql.DB
}

type Shorturl struct {
	Id         int64  `json:"id"`
	Url        string `json:"url"`
	CreateTime int64  `json:"create_time"`
}

func NewModel() *Model {
	return &Model{db: mysql.New("test")}
}

func (m *Model) GetUrl(id int) (string, error) {
	var url string
	err := m.db.QueryRow("SELECT `url` FROM shorturl WHERE `id` = ? limit 1", id).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (m *Model) AddUrl(url string) (int64, error) {
	now := time.Now().Unix()
	fmt.Println(now)
	rs, err := m.db.Exec("insert into shorturl (url,create_time) values (?,?)", url, now)
	fmt.Println(rs, err)

	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/magiconair/properties"
)

var db *sql.DB

func connect() {
	p := properties.MustLoadFile("./config.properties", properties.UTF8)
	user := p.GetString("user", "")
	password := p.GetString("password", "")
	url := fmt.Sprintf("%s:%s@/api_go_jwt_tasks?parseTime=true", user, password)

	conn, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	db = conn
}

func close() {
	db.Close()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	connect()
	result, err := db.Exec(query, args...)
	close()

	if err != nil {
		return nil, err
	}
	return result, nil
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	connect()
	rows, err := db.Query(query, args...)
	close()

	if err != nil {
		return nil, err
	}
	return rows, nil
}

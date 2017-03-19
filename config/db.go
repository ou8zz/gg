package config

import (
	"database/sql"
	"log"
	"runtime"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

type Settings struct {
	DbUrl   string
}

/**
 * 数据库链接
 */
var settings Settings = Settings{}
func Conn() *sql.DB {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("no caller info")
	}
	dbJson := filepath.Join(filepath.Dir(filename), "db.json")
	dbJsonByte, err := ioutil.ReadFile(dbJson)
	err = json.Unmarshal(dbJsonByte, &settings)
	db, _ := sql.Open("mysql", settings.DbUrl)
	err = db.Ping()
	if err != nil {
		fmt.Println("err : ", err)
		return nil
	}
	return db
}

/**
 * defer关闭数据库连接
 */
func CloseDb(conn sql.DB) {
	conn.Close()
	log.Println("close conn")
}

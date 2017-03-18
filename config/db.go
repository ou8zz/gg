package config

import (
	"database/sql"
	"log"
	)

/**
 * 数据库链接
 */
func Conn() *sql.DB {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/cailian?charset=utf8mb4")
	return db
}

/**
 * defer关闭数据库连接
 */
func DeferCloseDb(conn sql.DB) {
	conn.Close()
	log.Print("close conn")
}

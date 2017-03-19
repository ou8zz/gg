package model

import (
	"gg/config"
	"log"
	"database/sql"
)

// 实体Model
type User struct {
	Id int 		`json:"id"`
	Name  string 	`json:"name"`
	Email string 	`json:"email"`
	Tdate string 	`json:"tdate"`
}

// 查询sql, return userlist
func GetUser(u *User) ([]User, error) {
	// 动态参数
	param := ""
	var datas []interface{}
	if(u.Name != "") {
		param += " and name=?"
		datas = append(datas, u.Name)
	}
	if(u.Email != "") {
		param += " and email=?"
		datas = append(datas, u.Email)
	}
	sql := "select * from user where 1=1 " + param + " order by id"
	log.Println(sql)

	// SQL查询
	conn := config.Conn()
	defer config.CloseDb(*conn)
	query, err := conn.Query(sql, datas...)
	log.Println(err)

	// 迭代绑定json对象数组
	res := make([]User, 0)
	for query.Next() { //循环，让游标往下移动
		u := User{}
		err := query.Scan(&u.Id, &u.Name, &u.Email, &u.Tdate)
		log.Println(err)
		res = append(res, u)
	}
	return res, err
}

// insert数据
func SaveUser(u *User) (lid int64, err error) {
	// 事物控制
	conn := config.Conn()
	defer config.CloseDb(*conn)
	tx, err := conn.Begin()
	res, err := tx.Exec("insert into user (name, email, tdate) values (?, ?, ?)", "ole", u.Email, u.Tdate)
	res, err = tx.Exec("insert into user (name, email, tdate) values (?, ?, ?)", u.Name, u.Email, u.Tdate)
	defer Transtion(tx)
	if err != nil {
		panic(-1)
	}

	// 无事物
	sql := "insert into user (name, email, tdate) values (?, ?, ?)"
	log.Print(sql)
	res, err = conn.Exec(sql, "sherry", u.Email, "2017-03-18")

	// 获取最后一个insert的ID
	log.Print(res.LastInsertId())
	lid, _ = res.LastInsertId()
	return lid, err
}

func Transtion(tx *sql.Tx) {
	err := recover()
	if err != nil {
		tx.Rollback()
		log.Println("Exception model panic:", err)
	} else {
		tx.Commit()
	}
}
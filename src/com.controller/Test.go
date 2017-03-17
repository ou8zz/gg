package main

import (
	"net/http"
	"github.com/labstack/echo"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	e := echo.New()
	e.GET("/urlUser/:id", urlUser)
	e.POST("/getUser", getUser)
	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":85"))

}

/**
 * rest url test
 */
func urlUser(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	return c.String(http.StatusOK, id)
}

/**
 * 查询user列表
 */
func getUser(c echo.Context) error {
	// 绑定User表单对象
	u := &User{}
	c.Bind(u)

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
	log.Print(sql)

	// SQL查询
	conn := conn()
	defer deff(*conn)
	query, err := conn.Query(sql, datas...)
	checkErr(err)

	// 迭代绑定json对象数组
	res := make([]User, 0)
	for query.Next() { //循环，让游标往下移动
		u := User{}
		err := query.Scan(&u.Id, &u.Name, &u.Email, &u.Tdate)
		checkErr(err)
		res = append(res, u)
	}
	rd := ResponseData{0, "", res, 0, "", ""}
	return c.JSON(http.StatusCreated, rd)
}

// insert数据
func save(c echo.Context) error {
	// 在后面程序出现异常的时候就会捕获
	defer recoverException()

	// 绑定User表单对象
	u := &User{}
	if err := c.Bind(u); err != nil {
		fmt.Println(err, "save errrr")
		return err
	}

	// 事物控制
	conn := conn()
	defer deff(*conn)
	tx, _ := conn.Begin()
	res, err := tx.Exec("insert into user (name, email, tdate) values (?, ?, ?)", u.Name, u.Email, u.Tdate)
	res, err = tx.Exec("insert into user (name, email, tdate) values (?, ?, ?)", u.Name, u.Email, u.Tdate)

	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusOK, "faild")
	} else {
		tx.Commit()
	}

	// 无事物
	sql := "insert into user (name, email, tdate) values (?, ?, ?)"
	log.Print(sql)
	res, err = conn.Exec(sql, u.Name, u.Email, u.Tdate)

	// 获取最后一个insert的ID
	log.Print(res.LastInsertId())
	lid, _ := res.LastInsertId()

	rd := ResponseData{0, "新增成功", lid, 0, "", ""}
	return c.JSON(http.StatusCreated, rd)
}

/**
 * response返回json对象
 */
type ResponseData struct {
	ErrorCode int 		`json:"errorCode"`	// 必需 错误码。正常返回0 异常返回560 错误提示561对应errorInfo
	ErrorInfo string 	`json:"errorInfo"`	// 必需 错误信息。正常返回空”” 异常返回错误信息文本
	Data interface{} 	`json:"data"`		// 可选 返回数据内容。 如果有返回数据，可以是字符串或者数组JSON等等
	Total int 		`json:"total"`		// 可选 分页字段：总条数
	PageNum string 		`json:"pageNum"`	// 可选 分页字段：当前页数
	PageSize string 	`json:"pageSize"`	// 可选 分页字段：当前每页多少条数
}

/**
 * 数据库链接
 */
func conn() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cailian?charset=utf8mb4")
	checkErr(err)
	return db
}

// 在后面程序出现异常的时候就会捕获
func recoverException() {
	if r := recover(); r != nil {
		// 这里可以对异常进行一些处理和捕获
		log.Println("Exception:", r)
	}
}

// defer关闭数据库连接
func deff(conn sql.DB) {
	if conn != nil {
		conn.Close()
	}
	log.Print("close conn")
}

// 错误打印
func checkErr(errMasg error) {
	if errMasg != nil {
		//panic(errMasg)
		log.Print(errMasg)
	}
}

// 1. mian方法就是服务入库，对应的action应都统一管理
// 2. model实体和res实体返回直接建立在controller中使用
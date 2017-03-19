package controller

import (
	"net/http"
	"gg/utils"
	"github.com/labstack/echo"
	"fmt"
	"gg/service"
	"gg/model"
)

// 全局返回对象
var rd = utils.ResponseData{0, "", nil, 0, "", ""}

/**
 * rest url test
 */
func UrlUser(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	return c.String(http.StatusOK, id)
}

/**
 * 查询user列表
 */
func GetUser(c echo.Context) error {
	defer utils.RecoverException()

	// 绑定User表单对象
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		fmt.Print(err)
	}

	// 查询数据库
	rd.Data, rd.ErrorInfo = service.GetUser(u)
	return c.JSON(http.StatusCreated, rd)
}

/**
 * 保存数据
 */
func SaveUser(c echo.Context) error {
	// 在后面程序出现异常的时候就会捕获
	defer utils.RecoverException()

	// 绑定User表单对象
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		rd.Data = err
		rd.ErrorInfo = "提交参数错误"
		return c.JSON(http.StatusOK, rd)
	}

	// 数据保存
	rd.Data, rd.ErrorInfo = service.SaveUser(u)
	rd.Data = "新增成功"
	return c.JSON(http.StatusCreated, rd)
}

/**
 * 上传文件数据
 */
func UploadFile(c echo.Context) error {
	// 在后面程序出现异常的时候就会捕获
	defer utils.RecoverException()

	// 获取表单参数
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// 文件保存
	name, err = service.UploadFile(name, *avatar)
	if err != nil {
		rd.ErrorInfo = err
		return c.JSON(http.StatusCreated, rd)
	}

	rd.ErrorInfo = name
	rd.Data = "上传成功"
	return c.JSON(http.StatusCreated, rd)
}
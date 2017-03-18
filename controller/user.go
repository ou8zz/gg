package controller

import (
	"net/http"
	"gg/utils"
	"github.com/labstack/echo"
	"fmt"
	"gg/service"
	"gg/model"
	"errors"
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

	rd.Data, rd.ErrorInfo = service.GetUser(u)
	return c.JSON(http.StatusCreated, rd)
}

/**
 * 保存数据
 */
func SaveUser(c echo.Context) error {
	// 在后面程序出现异常的时候就会捕获
	//defer utils.RecoverException()
	defer func() (err error) {
		if r := recover(); r != nil {
			// 这里可以对异常进行一些处理和捕获
			fmt.Println("Exception:", r)

			//check exactly what the panic was and create error.
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknow panic")
			}
		}
		return err
	}()

	// 绑定User表单对象
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		rd.Data = err
		rd.ErrorInfo = "提交参数错误"
		return c.JSON(http.StatusOK, rd)
	}

	rd.Data, rd.ErrorInfo = service.SaveUser(u)
	rd.Data = "新增成功"
	return c.JSON(http.StatusCreated, rd)
}
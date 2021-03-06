package main

import (
	"github.com/labstack/echo"
	"gg/controller"
)

func main() {
	e := echo.New()
	e.Group("/admin/")
	e.GET("/urlUser/:id", controller.UrlUser)
	e.POST("/getUser", controller.GetUser)
	e.POST("/save", controller.SaveUser)
	e.POST("/upload", controller.UploadFile)

	e.Logger.Fatal(e.Start(":8080"))
}

package service

import (
	"gg/model"
	"os"
	"io"
	"mime/multipart"
)

// 查询userlist
func GetUser(u *model.User) ([]model.User, error) {
	res, err := model.GetUser(u)
	return res, err
}

// insert数据
func SaveUser(u *model.User) (int64, error) {
	// 事物控制
	lid, err := model.SaveUser(u)
	return lid, err
}

// 上传文件处理
func UploadFile(name string, avatar multipart.FileHeader) (nam string, err error) {
	src, err := avatar.Open()
	if err != nil {
		return name, err
	}
	defer src.Close()

	// 在指定目录创建新文件
	dst, err := os.Create("/Users/ole/dev/"+avatar.Filename)
	if err != nil {
		return name, err
	}
	defer dst.Close()

	// 复制内容到新文件中
	if _, err = io.Copy(dst, src); err != nil {
		return name, err
	}
	return name, err
}
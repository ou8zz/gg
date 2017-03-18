package service

import (
	"gg/model"
	"gg/utils"
)

// 查询userlist
func GetUser(u *model.User) ([]model.User, error) {
	res, err := model.GetUser(u)
	return res, err
}

// insert数据
func SaveUser(u *model.User) (int64, error) {
	defer utils.RecoverException()
	
	// 事物控制
	lid, err := model.SaveUser(u)
	panic(-3)
	return lid, err
}
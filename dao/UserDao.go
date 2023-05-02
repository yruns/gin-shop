package userDao

import (
	"gin-shop/model"
	"gin-shop/settings/db"
)

// VerifyUser 校验用户账号密码
func VerifyUser(username, password string) model.User {
	var user model.User
	db.Mysql.Table("user").Where("name = ? AND password = ?", username, password).Find(&user)
	return user
}

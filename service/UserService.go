package userService

import (
	"gin-shop/dao"
	"gin-shop/model"
	"gin-shop/request"
)

type UserService struct {
}

func Login(loginRequest request.LoginRequest) model.User {

	var user model.User
	// 校验账号密码
	user = userDao.VerifyUser(loginRequest.Name, loginRequest.Password)

	return user
}

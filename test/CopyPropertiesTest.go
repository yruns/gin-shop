package main

import (
	"fmt"
	"gin-shop/model"
	"gin-shop/request"
	"github.com/jinzhu/copier"
)

func main() {

	user := model.UserBasic{}

	loginRequest := request.LoginRequest{
		Name:     "1111",
		Password: "12222",
		Value:    "123",
	}

	//utils.CopyProperties(&loginRequest, &user)
	copier.Copy(&user, &loginRequest)

	fmt.Println(user)
}

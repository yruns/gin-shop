package controller

import (
	"gin-shop/middlewares"
	"gin-shop/request"
	"gin-shop/service"
	"gin-shop/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u *UserController) Register(router *gin.Engine) {
	router.POST("/login_pwd", login)
	router.GET("/captcha", captcha)
	//router.POST("/verifycha", verify)
	router.POST("/upload/avatar", uploadAvatar)
}

func uploadAvatar(c *gin.Context) {
	// 解析上传的参数：file，user_id
	userId, _ := c.Get("id")

	// 保存文件到本地

	// 保存文件路径到用户表的头像字段

}

// 生成验证码
func captcha(c *gin.Context) {
	utils.GenerateCaptcha(c)
}

// 登录
func login(c *gin.Context) {
	// 解析参数
	var loginRequest request.LoginRequest
	// 获取Body中的参数

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.Fail(c, "参数解析失败")
		return
	}

	// 判断验证码，暂不处理该逻辑

	// 判断用户名和密码
	user := userService.Login(loginRequest)

	// 返回结果
	if user.Id == 0 {
		utils.Fail(c, "用户名或密码错误")
	}

	// 生成jwt token
	if token, err := middlewares.GenerateToken(user.Id, user.Name); err != nil {
		utils.Fail(c, "token生成失败")
	} else {
		c.Header("Authorization", token)
		c.Header("Access-Control-Expose-Headers", "Authorization")
		utils.Ok(c, user)
	}
}

// 验证验证码
func verify(c *gin.Context) {
	var captcha utils.CaptchaResult
	if err := c.ShouldBindJSON(&captcha); err != nil {
		utils.Fail(c, err.Error())
		return
	}
	if utils.VerifyCaptcha(captcha.Id, captcha.VerifyValue) {
		utils.Ok(c, nil)
		return
	}
	utils.Fail(c, "验证码错误")
}

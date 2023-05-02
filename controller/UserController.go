package controller

import (
	"gin-shop/middlewares"
	"gin-shop/request"
	service "gin-shop/service"
	"gin-shop/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"strings"
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
	file, err := c.FormFile("avatar")
	if err != nil {
		logrus.Error(err.Error())
		utils.Fail(c, "图片上传失败，请重新上传")
		return
	}
	// 保存文件到本地
	uuid, _ := uuid.NewRandom()
	fileType := file.Filename[strings.LastIndex(file.Filename, "."):]
	fileName := uuid.String() + fileType
	localPath := "./upload/" + fileName

	err = c.SaveUploadedFile(file, localPath)
	if err != nil {
		logrus.Error(err.Error())
		utils.Fail(c, "图片保存失败，请重新上传")
		return
	}
	// 保存文件路径到用户表的头像字段
	minioPath, err, flag := service.SaveAvatar(userId.(int64), c, fileType, fileName, localPath)

	if err == nil && flag {
		utils.Ok(c, minioPath)
	} else {
		logrus.Error(err.Error())
		utils.Fail(c, "图片保存失败，请重新上传")
	}
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
	user := service.Login(loginRequest)

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

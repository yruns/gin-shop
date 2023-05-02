package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "请求成功",
		"data": data,
	})
}

//func Fail(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"code": 4002,
//		"msg":  "请求失败",
//		"data": nil,
//	})
//}

func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  msg,
		"data": nil,
	})
}

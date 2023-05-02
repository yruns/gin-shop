package middlewares

import (
	"gin-shop/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type Claims struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(viper.GetString("jwt.secret"))

func GenerateToken(Id int64, Name string) (string, error) {
	// 设置过期时间
	expireTime := time.Now().Add(3600 * time.Second)
	claims := Claims{
		Id:   Id,
		Name: Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 过滤登录和注册接口
		if strings.Contains(c.Request.URL.Path, "/login_pwd") || strings.Contains(c.Request.URL.Path, "/register") {
			c.Next()
			return
		}

		// 从请求头中获取Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.Fail(c, "Authorization header required")
			c.Abort()
			return
		}

		// 解析Token
		claims := Claims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		})

		if err != nil {
			utils.Fail(c, err.Error())
			c.Abort()
			return
		}

		// 验证Token是否有效
		if !token.Valid {
			utils.Fail(c, "Token已失效")
			c.Abort()
			return
		}

		// 验证是否过期
		if time.Unix(claims.ExpiresAt.Time.Unix(), 0).Sub(time.Now()) < 0 {
			utils.Fail(c, "Token已过期")
			c.Abort()
			return
		}

		// 将用户信息写入上下文
		c.Set("name", claims.Name)
		c.Set("id", claims.Id)

		c.Next()
	}
}

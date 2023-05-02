package userService

import (
	"fmt"
	"gin-shop/dao"
	"gin-shop/model"
	"gin-shop/request"
	"gin-shop/settings/db"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"time"
)

type UserService struct {
}

func Login(loginRequest request.LoginRequest) model.User {

	var user model.User
	// 校验账号密码
	user = userDao.VerifyUser(loginRequest.Name, loginRequest.Password)

	return user
}

func SaveAvatar(userId int64, c *gin.Context, fileType, fileName, localPath string) (string, error, bool) {

	// 写入到minio中
	contentType := "application/" + fileType

	// 获取当前年月日
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	// 生成文件名
	bucketName := viper.GetString("minio.mediaBucket")
	fmt.Println("bucketName:", bucketName)

	minioPath := fmt.Sprintf("%d/%d/%d/%s", year, month, day, fileName)
	fmt.Println("minioPath:", minioPath)

	// 使用FPutObject上传一个zip文件。
	_, err := db.Minio.FPutObject(c, bucketName, minioPath, localPath,
		minio.PutObjectOptions{ContentType: contentType})

	// 记录落库
	affected := db.Mysql.Table("user").
		Where("id = ?", userId).
		Update("avatar", minioPath).
		RowsAffected

	return minioPath, err, affected > 0
}

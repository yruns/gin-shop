package db

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var Minio *minio.Client

func InitMinio() (err error) {
	endpoint := viper.GetString("minio.endpoint")
	accessKeyID := viper.GetString("minio.accessKeyID")
	secretAccessKey := viper.GetString("minio.secretAccessKey")
	useSSL := false

	// Initialize minio client object.
	Minio, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	return err
}

package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func main() {
	ctx := context.Background()
	//endpoint := viper.GetString("minio.endpoint")
	//accessKeyID := viper.GetString("minio.accessKeyID")
	//secretAccessKey := viper.GetString("minio.secretAccessKey")
	endpoint := "127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
	log.Printf("初始化成功")

	buckets, err := minioClient.ListBuckets(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

	fmt.Println("============上传文件=============")
	objectName := "2023/05/02/1683025986pict.jpg"
	filePath := "H:\\Codefield\\Go\\gin-shop\\upload\\1683025986pict.jpg"
	contentType := "application/jpg"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(context.Background(), "mediafiles", objectName, filePath,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

}

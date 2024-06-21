package service

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var minioClient *s3.S3

func InitMinio() error {
	var err error
	minioClient, err = MinioConect()
	if err != nil {
		return err
	}
	return nil
}

func MinioConect() (*s3.S3, error) {
	// 環境変数を読み込む
	host := os.Getenv("MINIO_HOST")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	useSSL := false
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"), // 任意のリージョンを指定
		Endpoint:    aws.String(host),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		DisableSSL:  aws.Bool(useSSL),
		// S3ForcePathStyle: aws.Bool(true), // パススタイルのURLを使用するように設定
	})
	if err != nil {
		return nil, err
	}

	minioClient := s3.New(sess)
	return minioClient, nil
}

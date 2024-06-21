package util

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/service"
)

var bucketName = "mr-platform-content"

func InitContentFileBucket() error {
	err := service.InitMinio()
	if err != nil {
		fmt.Println("Failed to init minio", err)
		return err
	}
	buckets, err := GetContentFileBucket()
	if err != nil {
		fmt.Println("Failed to get buckets", err)
		return err
	}

	if !isExistBucket(buckets, bucketName) {
		fmt.Println("Bucket not found")
	}
	return nil
}

func GetContentUrl(minioUrl string) (string, error) {
	urlStr, err := service.GetObjectUrl(bucketName, minioUrl)
	if err != nil {
		fmt.Println("Failed to get object url", err)
		return "", err
	}
	return urlStr, nil
}

func GetContentFileBucket() (*s3.ListBucketsOutput, error) {
	buckets, err := service.GetBuckets()
	if err != nil {
		fmt.Println("Failed to get buckets", err)
		return nil, err
	}
	return buckets, nil
}

func CreateContentFileBucket(bucketName string) error {
	err := service.CreateBucket(bucketName)
	if err != nil {
		fmt.Println("Failed to create bucket", err)
		return err
	}
	return nil
}

func isExistBucket(buckets *s3.ListBucketsOutput, bucketName string) bool {
	for _, bucket := range buckets.Buckets {
		if *bucket.Name == bucketName {
			return true
		}
	}
	return false
}

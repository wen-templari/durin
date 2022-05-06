package util

import (
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadImg(source multipart.File) string {
	timeUnix := time.Now().UnixMilli() //单位秒
	target := strconv.FormatInt(timeUnix, 10)
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	accessKeyId := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")
	bucketName := "erebor"
	objectName := target
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Println("Error:", err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Println("Error:", err)
	}
	err = bucket.PutObject(objectName, source)
	if err != nil {
		log.Println("Error:", err)
	}
	fullTarget := "https://erebor.oss-cn-hangzhou.aliyuncs.com/" + target
	return fullTarget
}

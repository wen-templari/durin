package util

import (
	"log"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadImg(source multipart.File) string {
	timeUnix := time.Now().UnixMilli() //单位秒
	target := strconv.FormatInt(timeUnix, 10)
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := "***REMOVED***"
	accessKeySecret := "***REMOVED***"
	bucketName := "erebor"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := target
	// <yourLocalFileName>由本地文件路径加文件ƒ名包括后缀组成，例如/users/local/myfile.txt。
	// localFileName := source
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Println("Error:", err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Println("Error:", err)
	}

	// 上传文件。
	err = bucket.PutObject(objectName, source)
	if err != nil {
		log.Println("Error:", err)
	}
	fullTarget := "https://erebor.oss-cn-hangzhou.aliyuncs.com/" + target
	return fullTarget
}

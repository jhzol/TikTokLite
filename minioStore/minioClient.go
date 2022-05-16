package minioStore

import (
	"TikTokLite/log"
	"TikTokLite/util"
	"github.com/minio/minio-go/v6"
	"github.com/spf13/viper"
	"strings"
)

type Minio struct {
	MinioClient  *minio.Client
	endpoint     string
	port         string
	VideoBuckets string
	PicBuckets   string
}

func NewMinioClient() Minio {
	endpoint := viper.GetString("minio.host")
	port := viper.GetString("minio.port")
	endpoint = endpoint + ":" + port
	accessKeyID := viper.GetString("minio.accessKeyID")
	secretAccessKey := viper.GetString("minio.secretAccessKey")
	videoBucket := viper.GetString("minio.videobuckets")
	picBucket := viper.GetString("minio.picbuckets")
	useSSL := false

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Error(err)
	}
	//创建存储桶
	creatBucket(minioClient, videoBucket)
	creatBucket(minioClient, picBucket)
	if viper.GetString("minio.iswsl") == "true" {
		endpoint = viper.GetString("minio.winhost") + ":" + port
	}
	return Minio{minioClient, endpoint, port, videoBucket, picBucket}
}

func creatBucket(m *minio.Client, bucket string) {
	log.Debug("bucketname", bucket)
	found, err := m.BucketExists(bucket)
	if err != nil {
		log.Errorf("check %s bucketExists err:%s", bucket, err.Error())
	}
	if !found {
		m.MakeBucket(bucket, "us-east-1")
	}
}

func (m *Minio) UploadFile(bucket, file, userID string) (string, error) {
	var fileName strings.Builder
	var contentType, Suffix string
	if bucket == "video" {
		contentType = "video/mp4"
		Suffix = ".mp4"
	} else {
		contentType = "image/jpeg"
		Suffix = ".jpg"
	}
	fileName.WriteString(userID)
	fileName.WriteString("_")
	fileName.WriteString(util.GetCurrentTimeForString())
	fileName.WriteString(Suffix)
	n, err := m.MinioClient.FPutObject(bucket, fileName.String(), file, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Errorf("upload file error:%s", err.Error())
		return "", err
	}
	log.Info("upload file %dbyte success,fileName:%s", n, fileName)
	url := "http://" + m.endpoint + "/" + bucket + "/" + fileName.String()
	return url, nil
}

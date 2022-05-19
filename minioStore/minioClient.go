package minioStore

import (
	"TikTokLite/log"
	"TikTokLite/util"
	"github.com/minio/minio-go/v6"
	"github.com/spf13/viper"
	"strconv"
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
	// log.Debug("bucketname", bucket)
	found, err := m.BucketExists(bucket)
	if err != nil {
		log.Errorf("check %s bucketExists err:%s", bucket, err.Error())
	}
	if !found {
		m.MakeBucket(bucket, "us-east-1")
	}
	//设置桶策略
	policy := `{"Version": "2012-10-17",
				"Statement": 
					[{
						"Action":["s3:GetObject"],
						"Effect": "Allow",
						"Principal": {"AWS": ["*"]},
						"Resource": ["arn:aws:s3:::` + bucket + `/*"],
						"Sid": ""
					}]
				}`
	err = m.SetBucketPolicy(bucket, policy)
	if err != nil {
		log.Errorf("SetBucketPolicy %s  err:%s", bucket, err.Error())
	}
}

func (m *Minio) UploadFile(filetype, file, userID string) (string, error) {
	var fileName strings.Builder
	var contentType, Suffix, bucket string
	if filetype == "video" {
		contentType = "video/mp4"
		Suffix = ".mp4"
		bucket = m.VideoBuckets
	} else {
		contentType = "image/jpeg"
		Suffix = ".jpg"
		bucket = m.PicBuckets
	}
	fileName.WriteString(userID)
	fileName.WriteString("_")
	fileName.WriteString(strconv.FormatInt(util.GetCurrentTime(), 10))
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

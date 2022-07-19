package service

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/spf13/viper"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/qiniu/api.v7/v7/storage"
)

var (
	accessKey, secretKey, bucketName, domainName, upToken string
)

func initOSS() {
	accessKey = viper.GetString("oss.access_key")
	secretKey = viper.GetString("oss.secret_key")
	bucketName = viper.GetString("oss.bucket_name")
	domainName = viper.GetString("oss.domain_name")
}

func getToken() {
	var maxInt uint64 = 1 << 32
	initOSS()
	putPolicy := storage.PutPolicy{
		Scope:   bucketName,
		Expires: maxInt,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
}

func getObjectName(filename string, id uint32) (string, error) {
	i := strings.LastIndex(filename, ".")
	fileType := filename[i+1:]

	timeEpocNow := time.Now().Unix()
	objectName := strconv.FormatUint(uint64(id), 10) + "-" + strconv.FormatInt(timeEpocNow, 10) + "." + fileType
	return objectName, nil
}

func UploadFile(filename string, id uint32, r io.ReaderAt, dataLen int64) (string, error) {
	if upToken == "" {
		getToken()
	}

	objectName, err := getObjectName(filename, id)
	if err != nil {
		return "", err
	}

	// 下面是七牛云的oss所需信息，objectName对应key是文件上传路径
	cfg := storage.Config{Zone: &storage.ZoneHuanan, UseHTTPS: false, UseCdnDomains: true}
	formUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{Params: map[string]string{"x.name": "muxi-fresh-test"}}
	err = formUploader.Put(context.Background(), &ret, upToken, objectName, r, dataLen, &putExtra)
	if err != nil {
		return "", err
	}

	url := domainName + "/" + objectName
	return url, nil
}

func Download(url string) string {
	index := strings.LastIndex(url, "/")
	key := url[index+1:]
	domain := url[:index]
	fmt.Println(key, domain)
	mac := qbox.NewMac(accessKey, secretKey)
	deadline := time.Now().Add(time.Second * 3060).Unix() //  一小时有效期

	return storage.MakePrivateURL(mac, domain, key, deadline)
}

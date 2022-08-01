package service

import (
	"github.com/MuXiFresh-be/model/user"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
)

var (
	accessKey, secretKey, bucketName, domainName, upToken string
)

func GetProfile(id int) (*user.UserModel, error) {
	var user user.UserModel
	if err := user.GerInfo(id); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetToken() string {
	var maxInt uint64 = 1 << 32
	initOSS()
	putPolicy := storage.PutPolicy{
		Scope:   bucketName,
		Expires: maxInt,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
	return upToken
}

func initOSS() {
	accessKey = viper.GetString("oss.access_key")
	secretKey = viper.GetString("oss.secret_key")
	bucketName = viper.GetString("oss.bucket_name")
	domainName = viper.GetString("oss.domain_name")
}

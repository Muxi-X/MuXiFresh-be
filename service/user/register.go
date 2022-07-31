package service

import (
	MD5 "crypto/md5"
	"encoding/hex"
	USER "github.com/MuXiFresh-be/handler/user"
	User "github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
)

func Register(req USER.RegisterRequest) error {
	//前端自己验证两次密码是否一致
	//if req.Password != req.PasswordAgain {
	//	SendBadRequest(c, errno.ErrPasswordRepetition, nil, "please Re-enter the password", GetLine())
	//	return
	//}

	if err := User.IfExist(req.StudentId, req.Email, req.Name); err != nil {
		return errno.ServerErr(errno.ErrUserExisted, err.Error())
	}

	user := User.UserModel{
		Name:      req.Name,
		StudentId: req.StudentId,
		Email:     req.Email,
		Role:      1,
	}

	md5 := MD5.New()
	md5.Write([]byte(req.Password))
	user.HashPassword = hex.EncodeToString(md5.Sum(nil))
	if err := user.CreateUser(); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

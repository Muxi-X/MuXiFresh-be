package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/MuXiFresh-be/handler/user"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/pkg/errno"
)

func Register(req user.RegisterRequest) (err1, err2 error) {
	//前端自己验证两次密码是否一致
	//if req.Password != req.PasswordAgain {
	//	SendBadRequest(c, errno.ErrPasswordRepetition, nil, "please Re-enter the password", GetLine())
	//	return
	//}

	if err := model.IfExist(req.StudentId, req.Email, req.Name); err != nil {
		return errno.ErrUserExisted, err
	}

	var user model.UserModel
	user.Name = req.Name
	user.StudentId = req.StudentId
	user.Email = req.Email
	md5 := md5.New()
	md5.Write([]byte(req.Password))
	user.HashPassword = hex.EncodeToString(md5.Sum(nil))
	if err := model.DB.Self.Create(&user).Error; err != nil {
		return errno.ErrDatabase, err
	}
	return nil, nil
}

package service

import (
	"github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
)

func GetProfile(email string) (*user.UserModel, error) {
	user := &user.UserModel{}
	if err := user.GetInfo(email); err != nil {
		return nil, errno.ServerErr(errno.ErrGetUserInfo, err.Error())
	}
	return user, nil
}

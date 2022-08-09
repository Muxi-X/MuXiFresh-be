package service

import (
	"github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
)

func GetProfile(id int) (*user.UserModel, error) {
	var user user.UserModel
	if err := user.GerInfo(id); err != nil {
		return nil, errno.ServerErr(errno.ErrGetUserInfo, err.Error())
	}
	return &user, nil
}

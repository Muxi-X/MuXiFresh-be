package service

import (
	U "github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
)

// 更新用户信息

func UpdateInfo(email string, avatar string, name string) error {
	if err := U.UpdateInfo(email, avatar, name); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

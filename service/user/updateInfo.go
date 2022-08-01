package service

import (
	U "github.com/MuXiFresh-be/model/user"
)

// 更新用户信息
func UpdateInfo(email string, avatar string, name string) error {
	var user U.UserModel
	return user.UpdateInfo(email)
}

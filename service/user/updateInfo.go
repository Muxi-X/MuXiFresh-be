package service

import (
	U "github.com/MuXiFresh-be/model/user"
)

// 更新用户信息
func UpdateInfo(email string, avatar string, name string) error {
	return U.UpdateInfo(email, avatar, name)
}

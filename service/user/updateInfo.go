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

// 头像，姓名，学号
func UpdateInfor(email string, avatar string, name string, studentId string, college string, major string, grade string, gender string, contact_way string, contact_number string) error {
	if err := U.UpdateInfor(email, avatar, name, studentId, college, major, grade, gender, contact_way, contact_number); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
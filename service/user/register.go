package service

import (
	"github.com/MuXiFresh-be/handler/schedule"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
)

func Register(user user.UserModel) error {
	if err := model.DB.Self.Table("users").Create(&user).Error; err != nil {
		return errno.ErrDatabase
	}
	return nil
}

func Create(sche schedule.Schedules) error {
	if err := model.DB.Self.Table("schedules").Create(&sche).Error; err != nil {
		return errno.ErrDatabase
	}
	return nil
}

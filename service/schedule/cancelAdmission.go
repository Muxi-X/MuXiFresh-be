package service

import (
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/pkg/errno"
)

func CancelAdmit(name string) error {
	if err := model.DB.Self.Table("schedules").Where("name = ? ", name).Updates(map[string]interface{}{"admission_status": 0}).Error; err != nil {
		return errno.ErrDatabase
	}
	return nil
}

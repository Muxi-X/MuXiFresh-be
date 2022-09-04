package schedule

import (
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/schedule"
	"github.com/MuXiFresh-be/pkg/errno"
)

func CancelAdmit(email string) error {
	var num schedule.Admission_status
	if err := model.DB.Self.Table("schedules").Where("email = ? ", email).Select("admission_status").Find(&num).Error; err != nil {
		return errno.ErrDatabase
	}
	if err := model.DB.Self.Table("schedules").Where("email = ? ", email).Updates(map[string]interface{}{"admission_status": num.AdmissionStatus - 1}).Error; err != nil {
		return errno.ErrDatabase
	}
	return nil
}

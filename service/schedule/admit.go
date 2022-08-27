package schedule

import (
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/pkg/errno"
)

func Admit(name string, admit string) error {
	if err := model.DB.Self.Table("schedules").Where("name = ? ", name).Updates(map[string]interface{}{"admission_status": admit}).Error; err != nil {
		return errno.ErrDatabase
	}
	return nil
}

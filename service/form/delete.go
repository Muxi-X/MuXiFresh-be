package form

import (
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/model/schedule"
	"github.com/MuXiFresh-be/pkg/errno"
)

func DeleteForm(email string)  error {
	err :=form.Delete(email)
	if err != nil {
		return  errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

func DeleteSchedule(email string) error {
	err := schedule.Delete(email)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
package form

import (
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/model/schedule"
	"github.com/MuXiFresh-be/pkg/errno"
)

func EditForm(email string, group string, reason string,
	understand string, selfintroduction string, ifotherorganization string) error {
	if err := form.Edit(email,
		group, reason, understand,
		selfintroduction, ifotherorganization); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

func MoveGroup(email string,group string) error {
	if err := form.EditGroup(email,group); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

func MoveScheduleGroup(email string,group string) error {
	if err := schedule.EditGroup(email,group); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
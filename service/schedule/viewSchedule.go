package schedule

import (
	"fmt"

	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/schedule"
	"github.com/MuXiFresh-be/pkg/errno"
)

func ViewOwn(email string) (schedule.ScheduleModel, error) {
	var schedule schedule.ScheduleModel
	if err := model.DB.Self.Table("schedules").Where("email = ? ", email).Find(&schedule).Error; err != nil {
		fmt.Println(17, schedule)
		return schedule, errno.ErrDatabase
	}
	fmt.Println(18, schedule)
	return schedule, nil
}

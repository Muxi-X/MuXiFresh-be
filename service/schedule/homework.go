package schedule

import (
	"github.com/MuXiFresh-be/model/schedule"
)

func UpdateSchedule2(email string) error {
	if err := schedule.Update2(email); err != nil {
		return err
	}
	return nil
}

func UpdateSchedule3(email string) error {
	if err := schedule.Update3(email); err != nil {
		return err
	}
	return nil
}

func UpdateSchedule4(groupId uint) error {
	var group string
	switch groupId {
	case 1:
		group = "设计组"
	case 2:
		group = "产品组"
	case 3:
		group = "安卓组"
	case 4:
		group = "前端组"
	case 5:
		group = "后端组"

	}
	if err := schedule.Update4(group); err != nil {
		return err
	}
	return nil
}

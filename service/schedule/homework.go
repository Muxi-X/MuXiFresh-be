package schedule

import (
	"github.com/MuXiFresh-be/model/schedule"
)

func UpdateSchedule(email string) error {
	if err := schedule.Update2(email); err != nil {
		return err
	}
	return nil
}

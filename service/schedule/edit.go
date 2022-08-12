package schedule

import (
	"github.com/MuXiFresh-be/model/schedule"
)

func EditForm(email string, name string, major string, group string) error {
	return schedule.Edit(email, name, group, major)
}

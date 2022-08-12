package schedule

import (
	"github.com/MuXiFresh-be/model/schedule"
)

func CreateForm(email string, name string, studentId string, major string, group string) error {
	return schedule.Create(email, name, studentId, major, group)
}

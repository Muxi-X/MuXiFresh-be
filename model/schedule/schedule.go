package schedule

import (
	"github.com/MuXiFresh-be/model"
	"github.com/jinzhu/gorm"
)

type ScheduleModel struct {
	gorm.Model
	Email           string `gorm:"column:email;type:varchar(35);NOT NULL" json:"email"`
	Group           string `gorm:"column:group;type:varchar(20)" json:"group"`
	FormStatus      int    `gorm:"column:form_status;type:int(3);comment:'报名表状态 0-未提交 1-已提交';NOT NULL" json:"form_status"`
	WorkStatus      int    `gorm:"column:work_status;type:int(3);comment:'作业提交状态 0-未提交 1-已提交';NOT NULL" json:"work_status"`
	AdmissionStatus int    `gorm:"column:admission_status;type:int(3);comment:'录取状态 0-未录取 1-已录取';NOT NULL" json:"admission_status"`
}

func (s *ScheduleModel) TableName() string {
	return "schedules"
}

// Create ... create schedule
func (s *ScheduleModel) Create() error {
	return model.DB.Self.Create(s).Error
}

// Save ... save schedule.
func (s *ScheduleModel) Save() error {
	return model.DB.Self.Save(s).Error
}

func Create(email string, name string, studentId string, major string, group string) error {
	var schedule = ScheduleModel{

		Email: email,

		Group: group,
	}
	if err := schedule.Create(); err != nil {
		return err
	}
	return nil
}

func Edit(email string, name string, group string, major string) error {
	if err := model.DB.Self.Table("forms").Where("email=?", email).Updates(map[string]interface{}{"name": name, "major": major, "group": group}).Error; err != nil {
		return err
	}
	return nil
}

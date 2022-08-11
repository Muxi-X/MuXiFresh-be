package schedule

import (
	"github.com/MuXiFresh-be/model"
)

type ScheduleModel struct {
	Id              int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name            string `gorm:"column:name;type:varchar(20);NOT NULL" json:"name"`
	Email           string `gorm:"column:email;type:varchar(35);NOT NULL" json:"email"`
	StudentId       string `gorm:"column:student_id;type:char(10)" json:"student_id"`
	Collage         string `gorm:"column:collage;type:varchar(20)" json:"collage"`
	Major           string `gorm:"column:major;type:varchar(20)" json:"major"`
	Group           string `gorm:"column:group;type:varchar(20)" json:"group"`
	FormStatus      int    `gorm:"column:form_status;type:int(3);comment:报名表状态 0-未提交 1-已提交;NOT NULL" json:"form_status"`
	WorkStatus      int    `gorm:"column:work_status;type:int(3);comment:作业提交状态 0-未提交 1-已提交;NOT NULL" json:"work_status"`
	AdmissionStatus int    `gorm:"column:admission_status;type:int(3);comment:录取状态 0-未录取 1-已录取;NOT NULL" json:"admission_status"`
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

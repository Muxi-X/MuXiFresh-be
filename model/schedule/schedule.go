package schedule

import (
	"github.com/MuXiFresh-be/model"
)

type ScheduleModel struct {
	Email           string `gorm:"column:email;type:varchar(35);NOT NULL" json:"email"`
	Group           string `gorm:"column:group;type:varchar(20)" json:"group"`
	FormStatus      int    `gorm:"column:form_status;type:int(5);comment:'报名表状态 0-未提交 1-已报名 2-初试过 3-面试过';NOT NULL" json:"form_status"`
	WorkStatus      int    `gorm:"column:work_status;type:int(4);comment:'作业提交状态 0-未提交 1-已提交/未批阅 2-已批阅';NOT NULL" json:"work_status"`
	AdmissionStatus int    `gorm:"column:admission_status;type:int(3);comment:'录取状态 0-未录取 1-已录取';NOT NULL" json:"admission_status"`
	Name            string `gorm:"column:name;type:varchar(35);NOT NULL" json:"name"`
}

type Admission_status struct {
	AdmissionStatus int `gorm:"column:admission_status;type:int(3);comment:'录取状态 0-未录取 1-已录取';NOT NULL" json:"admission_status"`
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

func Edit(email string, name string, group string, major string) error {
	if err := model.DB.Self.Table("forms").Where("email=?", email).Updates(map[string]interface{}{"name": name, "major": major, "group": group}).Error; err != nil {
		return err
	}
	return nil
}

func Update(email string, group string) error {
	if err := model.DB.Self.Table("schedules").Where("email=?", email).Updates(map[string]interface{}{"group": group, "form_status": 1}).Error; err != nil {
		return err
	}
	return nil
}

func Update2(email string) error {
	if err := model.DB.Self.Table("schedules").Where("email=?", email).Updates(map[string]interface{}{"work_status": 1}).Error; err != nil {
		return err
	}
	return nil
}

func Update3(email string) error {
	if err := model.DB.Self.Table("schedules").Where("email=?", email).Updates(map[string]interface{}{"work_status": 0}).Error; err != nil {
		return err
	}
	return nil
}

func Update4(group string) error {
	if err := model.DB.Self.Table("schedules").Where("`group`=?", group).Updates(map[string]interface{}{"work_status": 0}).Error; err != nil {
		return err
	}
	return nil
}

//UPDATE `schedules` SET `work_status` = 1  WHERE (email='12345678')

func Delete(email string) error {
	var schedule = ScheduleModel{
		Email: email,
	}

	if err := model.DB.Self.Table("schedules").Where("email=?", email).Delete(&schedule).Error; err != nil {
		return err
	}
	return nil
}

func EditGroup(email string, group string) error {
	if err := model.DB.Self.Model(ScheduleModel{}).Where("email=?", email).Updates(ScheduleModel{
		Group: group,
	}).Error; err != nil {
		return err
	}
	return nil
}

package form

import (
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/jinzhu/gorm"
)

type FormModel struct {
	gorm.Model
	Email               string `json:"email" gorm:"column:email;not null"`
	Name                string `json:"name" gorm:"column:name;" `
	Avatar              string `json:"avatar" gorm:"column:avatar;"`
	StudentId           string `json:"student_id" gorm:"column:student_id;"`
	College             string `json:"college" gorm:"column:college;"`
	Major               string `json:"major" gorm:"column:major;"`
	Grade               string `json:"grade" gorm:"column:grade;"`
	Gender              string `json:"gender" gorm:"column:gender;"`
	ContactWay          string `json:"contact_way" gorm:"column:contact_way;"`
	ContactNumber       string `json:"contact_number" gorm:"column:contact_number;"`
	Group               string `json:"group" gorm:"column:group;"`
	Reason              string `json:"reason" gorm:"column:reason;"`
	Understand          string `json:"understand" gorm:"column:understand;"`
	SelfIntroduction    string `json:"self_introduction" gorm:"column:self_introduction;"`
	IfOtherOrganization string `json:"if_other_organization" gorm:"column:if_other_organization;"`
}

func (f *FormModel) TableName() string {
	return "forms"
}

// Create ... create form
//func (f *FormModel) Create() error {
//return model.DB.Self.Create(f).Error
//}

// Save ... save user
//func (f *FormModel) Save() error {
//return model.DB.Self.Save(f).Error
//}

func Create(email string, name string, avatar string, studentId string,
	college string, major string, grade string,
	gender string, contactway string, contactnumber string,
	group string, reason string, understand string,
	selfintroduction string, ifotherorganization string) (*FormModel, error) {
	var form = FormModel{
		Email:               email,
		Name:                name,
		Avatar:              avatar,
		StudentId:           studentId,
		College:             college,
		Major:               major,
		Grade:               grade,
		Gender:              gender,
		ContactWay:          contactway,
		ContactNumber:       contactnumber,
		Group:               group,
		Reason:              reason,
		Understand:          understand,
		SelfIntroduction:    selfintroduction,
		IfOtherOrganization: ifotherorganization,
	}
	tx := model.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&form).Error; err != nil {
		return nil, err
	}

	return &form, nil
}

func Edit(email string, name string, avatar string, studentid string,
	college string, major string, grade string,
	gender string, contactway string, contactnumber string,
	group string, reason string, understand string,
	selfintroduction string, ifotherorganization string) error {

	if err := model.DB.Self.Model(FormModel{}).Where("email=?", email).Updates(FormModel{
		Email:               email,
		Name:                name,
		Avatar:              avatar,
		StudentId:           studentid,
		College:             college,
		Major:               major,
		Grade:               grade,
		Gender:              gender,
		ContactWay:          contactway,
		ContactNumber:       contactnumber,
		Group:               group,
		Reason:              reason,
		Understand:          understand,
		SelfIntroduction:    selfintroduction,
		IfOtherOrganization: ifotherorganization,
	}).Error; err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

func Search(group string) ([]FormModel, error) {
	var form []FormModel
	if err := model.DB.Self.Table("forms").Where("`group`=?", group).Find(&form).Error; err != nil {
		return form, err
	}
	return form, nil
}

func View(email string) (*FormModel, error) {
	var form FormModel
	if err := model.DB.Self.Table("forms").Where("email=?", email).Find(&form).Error; err != nil {
		return nil, err
	}
	return &form, nil
}

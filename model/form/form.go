package form

import (
	"github.com/MuXiFresh-be/model"
)

type FormModel struct {
	Id                    uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Email                 string `json:"email" gorm:"column:email;default:null"`
	Name                  string `json:"name" gorm:"column:name;" binding:"required"`
	Avatar                string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	StudentId             string `json:"student_id" gorm:"column:student_id;"`
	College               string `json:"college" gorm:"column:college;"`
	Major                 string `json:"major" gorm:"column:major;"`
	Grade                 string `json:"grade" gorm:"column:grade;"`
	Gender                string `json:"gender" gorm:"column:gender;"`
	Contact_way           string `json:"contact_way" gorm:"column:contact_way;"`
	Contact_number        string `json:"contact_number" gorm:"column:contact_number;"`
	Group                 string `json:"group" gorm:"column:group;"`
	Reason                string `json:"reason" gorm:"column:reason;"`
	Understand            string `json:"understand" gorm:"column:understand;"`
	Self_introduction     string `json:"self_introduction" gorm:"column:self_introduction;"`
	If_other_organization string `json:"if_other_organization" gorm:"column:if_other_organization;"`
}

func (f *FormModel) TableName() string {
	return "forms"
}

// Create ... create form
func (f *FormModel) Create() error {
	return model.DB.Self.Create(f).Error
}

// Save ... save user
func (f *FormModel) Save() error {
	return model.DB.Self.Save(f).Error
}

// 
func ViewForm(email string) (FormModel, error) {
	var form FormModel
	if err := model.DB.Self.Table("forms").Where("email=?", email).Find(&form).Error; err != nil {
		return FormModel{}, err
	}
	return form, nil
}
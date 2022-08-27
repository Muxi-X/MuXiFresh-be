package form

import (
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/schedule"
	"github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/jinzhu/gorm"
)

type FormModel struct {
	gorm.Model
	Email string `json:"email" gorm:"column:email;not null;unique"`
	//Name                string `json:"name" gorm:"column:name;" `
	//Avatar              string `json:"avatar" gorm:"column:avatar;"`
	//StudentId           string `json:"student_id" gorm:"column:student_id;"`
	//College             string `json:"college" gorm:"column:college;"`
	//Major               string `json:"major" gorm:"column:major;"`
	//Grade               string `json:"grade" gorm:"column:grade;"`
	//Gender              string `json:"gender" gorm:"column:gender;"`
	//ContactWay          string `json:"contact_way" gorm:"column:contact_way;"`
	//ContactNumber       string `json:"contact_number" gorm:"column:contact_number;"`
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

func Create(email string,
	group string, reason string, understand string,
	selfintroduction string, ifotherorganization string) (*FormModel, error) {
	var form = FormModel{
		Email: email,
		//Name:                name,
		//Avatar:              avatar,
		//StudentId:           studentId,
		//College:             college,
		//Major:               major,
		//Grade:               grade,
		//Gender:              gender,
		//ContactWay:          contactway,
		//ContactNumber:       contactnumber,
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

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &form, nil
}

func Edit(email string,
	group string, reason string, understand string,
	selfintroduction string, ifotherorganization string) error {

	if err := model.DB.Self.Model(FormModel{}).Where("email=?", email).Updates(FormModel{
		Email: email,
		//Name:                name,
		//Avatar:              avatar,
		//StudentId:           studentid,
		//College:             college,
		//Major:               major,
		//Grade:               grade,
		//Gender:              gender,
		//ContactWay:          contactway,
		//ContactNumber:       contactnumber,
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

func GetUserInfoByGroup(status string) ([]user.UserModel, error) {

	var group string
	var formstatus int
	group = ""
	formstatus = 0 //formstatus默认为0-未提交，其他分别为1-已报名，2-初试过，3-面试过
	if status == "11" {
		group = "设计组"
		formstatus = 1
	}
	if status == "12" {
		group = "设计组"
		formstatus = 2
	}
	if status == "13" {
		group = "设计组"
		formstatus = 3
	}
	if status == "21" {
		group = "产品组"
		formstatus = 1
	}
	if status == "22" {
		group = "产品组"
		formstatus = 2
	}
	if status == "23" {
		group = "产品组"
		formstatus = 3
	}
	if status == "31" {
		group = "安卓组"
		formstatus = 1
	}
	if status == "32" {
		group = "安卓组"
		formstatus = 2
	}
	if status == "33" {
		group = "安卓组"
		formstatus = 3
	}
	if status == "41" {
		group = "前端组"
		formstatus = 1
	}
	if status == "42" {
		group = "前端组"
		formstatus = 2
	}
	if status == "43" {
		group = "前端组"
		formstatus = 3
	}
	if status == "51" {
		group = "后端组"
		formstatus = 1
	}
	if status == "52" {
		group = "后端组"
		formstatus = 2
	}
	if status == "53" {
		group = "后端组"
		formstatus = 3
	}

	//userinfo := make([]user.UserModel,100)
	//forminfo := make([]FormModel,100)
	/* var forminfo []FormModel
	if err := model.DB.Self.Table("forms").Where("`group`=?", group).Find(&forminfo).Error; err != nil {
		return []user.UserModel{}, err
	}

	len := len(forminfo) */

	/* userinfo := make([]user.UserModel,len)
	fmt.Print(len)
	for i := 0; i < len; i++ {
		if err := model.DB.Self.Table("users").Where("`email`=?", forminfo[i].Email).Find(&userinfo[i]).Error; err != nil {
			return userinfo, err
		}
	}
	return userinfo, nil */
	var scheduleinfo []schedule.ScheduleModel
	if formstatus == 1 {
		if err := model.DB.Self.Table("schedules").Where("`group`=? and admission_status= ? ", group, 0).Find(&scheduleinfo).Error; err != nil {
			return []user.UserModel{}, err
		}
	} else if formstatus != 1 {
		if err := model.DB.Self.Table("schedules").Where("`group`=? and admission_status= ?", group, formstatus-1).Find(&scheduleinfo).Error; err != nil {
			return []user.UserModel{}, err
		}
	}

	// if err := model.DB.Self.Table("schedules").Where("`group`=? and admission_status= ?", group, formstatus).Find(&scheduleinfo).Error; err != nil {
	// 	return []user.UserModel{}, err
	// }
	len := len(scheduleinfo)
	userinfo := make([]user.UserModel, len)

	for i := 0; i < len; i++ {
		if err := model.DB.Self.Table("users").Where("`email`=?", scheduleinfo[i].Email).Find(&userinfo[i]).Error; err != nil {
			return userinfo, err
		}
	}
	return userinfo, nil
}

func View(email string) (*FormModel, error) {
	var form FormModel
	if err := model.DB.Self.Table("forms").Where("email=?", email).Find(&form).Error; err != nil {
		return nil, err
	}
	return &form, nil
}

func Delete(email string) error {
	var form = FormModel{
		Email: email,
	}
	if err := model.DB.Self.Table("forms").Where("email=?", email).Delete(&form).Error; err != nil {
		return err
	}
	return nil
}

func EditGroup(email string, group string) error {
	if err := model.DB.Self.Model(FormModel{}).Where("email=?", email).Updates(FormModel{
		Group: group,
	}).Error; err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

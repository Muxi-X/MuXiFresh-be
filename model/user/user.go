package user

import (
	"errors"
	"fmt"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/file"
)

type UserModel struct {
	Id           uint32       `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string       `json:"name" gorm:"column:name;not null" binding:"required"`
	Email        string       `json:"email" gorm:"column:email;default:null"`
	Avatar       file.Picture `gorm:"foreignKey:Email;reference:Email"`
	Role         uint32       `json:"role" gorm:"column:role;" binding:"required"`
	Message      uint32       `json:"message" gorm:"column:message;" binding:"required"`
	HashPassword string       `json:"hash_password" gorm:"column:hash_password;" binding:"required"`
	StudentId    string       `json:"student_id" gorm:"column:student_id;"`
}

func (u *UserModel) TableName() string {
	return "users"
}

// Create ... create user
func (u *UserModel) Create() error {
	return model.DB.Self.Create(u).Error
}

// Save ... save user.
func (u *UserModel) Save() error {
	return model.DB.Self.Save(u).Error
}

// Get Information
func (user *UserModel) GerInfo(id int) error {
	if err := model.DB.Self.
		Where("id = ?", id).
		First(user).Error; err != nil {
		return err
	}
	var avatar file.Picture
	if err := model.DB.Self.
		Where("email  = ?", user.Email).
		First(&avatar).Error; err != nil {
		return err
	}
	user.Avatar = avatar
	return nil
}

func IfExist(id, email, name string) error {
	var user1 UserModel
	var user2 UserModel
	var user3 UserModel

	err1 := model.DB.Self.Debug().Where("student_id=?", id).First(&user1).Error
	err2 := model.DB.Self.Debug().Where("email=?", email).First(&user2).Error
	err3 := model.DB.Self.Debug().Where("name=?", name).First(&user3).Error

	s := []string{""}
	i := 0
	if err1 == nil {
		s = append(s, "*学号*")
		i++
	}

	if err2 == nil {
		s = append(s, "*邮箱*")
		i++
	}

	if err3 == nil {
		s = append(s, "*姓名*")
		i++

	}

	if i > 0 {
		s = append(s, "已被注册")
	}

	if i > 0 {
		return errors.New(fmt.Sprintf("%s", s))
	}

	return nil

}

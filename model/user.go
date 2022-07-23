package model

import (
	"errors"
	"fmt"
)

type UserModel struct {
	Id           uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string `json:"name" gorm:"column:name;" binding:"required"`
	Email        string `json:"email" gorm:"column:email;default:null"`
	Avatar       string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	Role         uint32 `json:"role" gorm:"column:role;" binding:"required"`
	Message      uint32 `json:"message" gorm:"column:message;" binding:"required"`
	HashPassword string `json:"hash_password" gorm:"column:hash_password;" binding:"required"`
	StudentId    string `json:"student_id" gorm:"column:student_id;"`
}

func (u *UserModel) TableName() string {
	return "user_models"
}

// Create ... create user
func (u *UserModel) Create() error {
	return DB.Self.Create(u).Error
}

// Save ... save user.
func (u *UserModel) Save() error {
	return DB.Self.Save(u).Error
}

func GetUserByStudentId(studentId string) (*UserModel, error) {
	return nil, nil
}

func IfExist(id, email, name string) error {
	var user1 UserModel
	var user2 UserModel
	var user3 UserModel

	err1 := DB.Self.Debug().Where("student_id=?", id).First(&user1).Error
	err2 := DB.Self.Debug().Where("email=?", email).First(&user2).Error
	err3 := DB.Self.Debug().Where("name=?", name).First(&user3).Error

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

package user

import (
	"errors"
	"fmt"
	"github.com/MuXiFresh-be/model"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model
	Name         string `json:"name" gorm:"column:name;not null" binding:"required"`
	Email        string `json:"email" gorm:"column:email;default:null;unique"`
	Avatar       string `json:"avatar"gorm:"column:avatar"`
	Role         uint32 `json:"role" gorm:"column:role;" binding:"required"`
	Message      uint32 `json:"message" gorm:"column:message;" binding:"required"`
	HashPassword string `json:"hash_password" gorm:"column:hash_password;" binding:"required"`
	StudentId    string `json:"student_id" gorm:"column:student_id;unique™"`
}

func (u *UserModel) TableName() string {
	return "users"
}

// Create ... create user
func (u *UserModel) CreateUser() error {
	tx := model.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Save ... save user.
func (u *UserModel) Save() error {
	return model.DB.Self.Save(u).Error
}

// Get Information
func (user *UserModel) GetInfo(email string) error {
	fmt.Printf("email:%s\n", email)
	if err := model.DB.Self.Model(UserModel{}).
		Where("email = ?", email).
		First(user).Error; err != nil {
		fmt.Println("error", err)
		return err
	}
	return nil
}

func IfExist(id, email, name string) error {
	var user1 UserModel
	var user2 UserModel
	var user3 UserModel

	err1 := model.DB.Self.Debug().Where("student_id=?", id).Find(&user1).Error
	err2 := model.DB.Self.Debug().Where("email=?", email).Find(&user2).Error
	err3 := model.DB.Self.Debug().Where("name=?", name).Find(&user3).Error

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

func UpdateInfo(email string, avatar string, name string) error {
	var user = UserModel{
		Email:  email,
		Avatar: avatar,
		Name:   name,
	}
	tx := model.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(user).Where("email = ?", user.Email).Update(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Authorize ...授权
func Authorize(email string, role int) error {
	Role := uint32(role)
	tx := model.DB.Self.Begin()
	var user UserModel
	model.DB.Self.Model(UserModel{}).Where("email = ?", email).Find(&user)
	if user.Role == 4 {
		return errors.New("can't change super admin")
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&UserModel{}).Where("email = ?", email).Update(UserModel{Role: Role}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetAdmin ...获取管理员
func GetAdimin(role int) ([]UserModel, error) {
	var user []UserModel
	err := model.DB.Self.Model(&UserModel{}).Where("role = ?", role).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

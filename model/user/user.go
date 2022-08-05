package user

import (
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/file"
)

type UserModel struct {
	ID           uint32       `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string       `json:"name" gorm:"column:name;not null" binding:"required"`
	Email        string       `json:"email" gorm:"column:email;default:null;unique"`
	Avatar       file.Picture `gorm:"foreignKey:Email;reference:Email"`
	Role         uint32       `json:"role" gorm:"column:role;" binding:"required"`
	Message      uint32       `json:"message" gorm:"column:message;" binding:"required"`
	HashPassword string       `json:"hash_password" gorm:"column:hash_password;" binding:"required"`
	StudentId    string       `json:"student_id" gorm:"column:student_id;unique™"`
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

	// var avatar file.Picture
	// avatar.Email = u.Email
	// avatar.Class = "Avatar"

	// if err := tx.Create(&avatar).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	return tx.Commit().Error
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

// func IfExist(id, email, name string) error {
// 	var user1 UserModel
// 	var user2 UserModel
// 	var user3 UserModel

// 	err1 := model.DB.Self.Debug().Where("student_id=?", id).First(&user1).Error
// 	err2 := model.DB.Self.Debug().Where("email=?", email).First(&user2).Error
// 	err3 := model.DB.Self.Debug().Where("name=?", name).First(&user3).Error

// 	s := []string{""}
// 	i := 0

// 	if err1 == nil {
// 		s = append(s, "*学号*")
// 		i++
// 	}

// 	if err2 == nil {
// 		s = append(s, "*邮箱*")
// 		i++
// 	}

// 	if err3 == nil {
// 		s = append(s, "*姓名*")
// 		i++

// 	}

// 	if i > 0 {
// 		s = append(s, "已被注册")
// 	}

// 	if i > 0 {
// 		return errors.New(fmt.Sprintf("%s", s))
// 	}

// 	return nil

// }

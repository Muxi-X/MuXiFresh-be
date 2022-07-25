package user

import "github.com/MuXiFresh-be/model"

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

func GetUserByStudentId(studentId string) (*UserModel, error) {
	return nil, nil
}

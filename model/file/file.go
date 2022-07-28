package file

import (
	"github.com/MuXiFresh-be/model"
	"github.com/jinzhu/gorm"
)

type HomeworkPublished struct {
	gorm.Model
	Publisher string `json:"publisher" gorm:"column:publisher"`
	Group     string `json:"group" gorm:"column:group"`
	Content   string `json:"content" gorm:"column:content"`
	Homeworks []Homework
}

type Homework struct {
	gorm.Model
	HomeworkID uint   `json:"homework_id" gorm:"column:homework_id"`
	Email      string `json:"email" gorm:"column:email"`
	URL        string `json:"url" gorm:"column:url"`
}

type Picture struct {
	gorm.Model
	Email string `json:"email" gorm:"column:email;unique"`
	Class string `json:"class" gorm:"column:class"`
	URL   string `json:"URL" gorm:"column:url"`
}

func (avatar *Picture) Update() error {
	return model.DB.Self.Where("email = ?", avatar.Email).Update(Picture{URL: avatar.URL}).Error
}

func (homework *Homework) Create() error {
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

	return tx.Commit().Error
}

package file

import (
	"github.com/jinzhu/gorm"
)

type HomeworkPublished struct {
	gorm.Model
	Publisher string `json:"publisher" gorm:"column:publisher"`
	GroupID   uint   `json:"group_id" gorm:"column:group_id"`
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

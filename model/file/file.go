package file

import (
	"github.com/jinzhu/gorm"
)

type HomeworkPublished struct {
	gorm.Model
	Publisher string `json:"publisher" gorm:"column:publisher"`
	GroupID   uint   `json:"group_id" gorm:"column:group_id"`
	Title     string `json:"title" gorm:"column:title"`
	Content   string `json:"content" gorm:"column:content"`
	FileUrl   string `json:"file_url"gorm:"column:url"`
	Homeworks []Homework
}

type Homework struct {
	gorm.Model
	HomeworkID uint   `json:"homework_id" gorm:"column:homework_id"`
	GroupID    uint   `json:"group_id"gorm:"column:group_id"`
	Title      string `json:"title" gorm:"column:title"`
	Content    string `json:"content" gorm:"column:content"`
	Email      string `json:"email" gorm:"column:email"`
	URL        string `json:"url" gorm:"column:url"`
	Status     int    `json:"status" gorm:"column:status"`
}

type Picture struct {
	gorm.Model
	Email string `json:"email" gorm:"column:email;unique"`
	Class string `json:"class" gorm:"column:class"`
	URL   string `json:"URL" gorm:"column:url"`
}

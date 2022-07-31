package comment

import (
	"errors"
	"github.com/MuXiFresh-be/model"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	HomeworkID uint   ` gorm:"column:homework_id"`
	Publisher  string ` gorm:"column:publisher"`
	Content    string `gorm:"size:255"`
}

func (comment *Comment) Create() error {
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
	if err := tx.Create(comment).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func DeleteComment(id string, email string) error {
	var comment Comment
	model.DB.Self.Where("id  = ?", id).Find(&comment)
	if comment.Publisher != email {
		return errors.New("permission denied")
	}
	tx := model.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where("id  = ?", id).
		Delete(&Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

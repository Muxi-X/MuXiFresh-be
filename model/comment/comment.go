package comment

import (
	"errors"
	"fmt"
	"github.com/MuXiFresh-be/model"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	HomeworkID uint   ` gorm:"column:homework_id"`
	Publisher  string ` gorm:"column:publisher"`
	Content    string `gorm:"size:255"`
}

func Create(email string, id uint, content string) error {
	var comment = Comment{
		HomeworkID: id,
		Publisher:  email,
		Content:    content,
	}
	fmt.Println("comment", comment)
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
	if err := tx.Create(&comment).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func DeleteComment(id string, email string) error {
	var comment Comment
	model.DB.Self.Model(Comment{}).Where("id  = ?", id).First(&comment)
	fmt.Println("publisher", comment.Publisher, "    email", email)
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

// 查询评论
func GetCommentList(id string, offset int, limit int) ([]Comment, int, error) {
	var item []Comment
	d := model.DB.Self.Table("comments").
		Where("homework_id = ?", id).
		Offset(offset).Limit(limit).
		Order("created_at desc").Find(&item)
	if err := d.Error; err != nil {
		return item, 0, err
	}
	var num int

	if err := d.Error; err != nil {
		return item, 0, err
	}
	for i, _ := range item {
		num = i + 1
	}
	return item, num, nil
}

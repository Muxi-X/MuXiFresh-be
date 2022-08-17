package comment

import (
	"errors"
	"fmt"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/file"
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/model/user"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	HomeworkID uint   ` gorm:"column:homework_id"`
	Publisher  string ` gorm:"column:publisher"`
	Content    string `gorm:"size:255"`
}
type List struct {
	Comment
	Name   string
	Avatar string
	Grade  string
	Group  string
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

	if err := tx.Model(file.Homework{}).Where("id = ?", id).Updates(file.Homework{Status: 1}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func DeleteComment(id string, email string) error {
	var comment Comment
	model.DB.Self.Model(Comment{}).Where("id  = ?", id).Find(&comment)
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
func GetCommentList(id string, offset int, limit int) ([]List, int, error) {
	var item []Comment
	d := model.DB.Self.Table("comments").
		Where("homework_id = ?", id).
		Offset(offset).Limit(limit).
		Order("created_at desc").Find(&item)
	if err := d.Error; err != nil {
		return nil, 0, err
	}
	num := 0
	if err := d.Error; err != nil {
		return nil, 0, err
	}
	list := make([]List, len(item))
	for i, m := range item {
		var u user.UserModel
		err2 := model.DB.Self.Model(user.UserModel{}).
			Where("email = ?", m.Publisher).
			Find(&u).Error
		if err2 != nil {
			return nil, 0, err2
		}
		var f form.FormModel
		err3 := model.DB.Self.Model(form.FormModel{}).
			Where("email = ?", m.Publisher).
			Find(&f).Error
		if err3 != nil {
			return nil, 0, err3
		}
		fmt.Println("u--------", u)
		list[i].Comment = m
		list[i].Name = u.Name
		list[i].Avatar = u.Avatar
		list[i].Grade = f.Grade
		list[i].Group = f.Group
		num = i + 1
	}
	return list, num, nil
}

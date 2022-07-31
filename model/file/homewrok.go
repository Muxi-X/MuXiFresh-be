package file

import "github.com/MuXiFresh-be/model"

// 提交作业
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
	if err := tx.Create(homework).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 发布作业
func (homework *HomeworkPublished) Publish() error {
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
	if err := tx.Create(homework).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

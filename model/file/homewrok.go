package file

import "github.com/MuXiFresh-be/model"

// Create ...提交作业
func Create(title string, content string, homeworkID uint, url string, email string) error {
	var homework = Homework{
		HomeworkID: homeworkID,
		Title:      title,
		Content:    content,
		Email:      email,
		URL:        url,
		Status:     0,
	}
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
	if err := tx.Create(&homework).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Publish 发布作业
func Publish(groupID uint, title string, content string, email string, url string) error {
	var homework = HomeworkPublished{
		GroupID:   groupID,
		Title:     title,
		Content:   content,
		Publisher: email,
		FileUrl:   url,
	}
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
	if err := tx.Create(&homework).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetHomework 查看不同组已发布的作业
func GetHomework(id int, offset int, limit int) ([]HomeworkPublished, int, error) {
	var item []HomeworkPublished
	d := model.DB.Self.Table("homework_publisheds").
		Where("group_id = ?", id).
		Offset(offset).Limit(limit).
		Order("created_at desc").Find(&item)
	if err := d.Error; err != nil {
		return item, 0, err
	}
	var num int = 0
	for i, _ := range item {
		num = i + 1
	}
	return item, num, nil
}

// ReviewHomework ...查阅作业
func ReviewHomework(id int) (*Homework, error) {
	var homework *Homework
	if err := model.DB.Self.Model(Homework{}).Where("id = ?", id).First(homework).Error; err != nil {
		return nil, err
	}
	return homework, nil
}

package file

import (
	"errors"
	"fmt"
	"github.com/MuXiFresh-be/model"
	form2 "github.com/MuXiFresh-be/model/form"
)

// Create ...提交作业
func Create(title string, content string, homeworkID uint, url string, email string) (*Homework, error) {
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
		return nil, err
	}

	var homeworkPublished HomeworkPublished
	if err := tx.Model(HomeworkPublished{}).Where("id = ?", homeworkID).
		Find(&homeworkPublished).Error; err != nil {
		return nil, err
	}
	homework.GroupID = homeworkPublished.GroupID

	if err := tx.Create(&homework).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &homework, nil
}

// Publish 发布作业
func Publish(groupID uint, title string, content string, email string, url string) (*HomeworkPublished, error) {
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
		return nil, err
	}
	if err := tx.Create(&homework).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &homework, nil
}

// GetHomeworkPublished 查看不同组已发布的作业
func GetHomeworkPublished(id int, offset int, limit int) ([]HomeworkPublished, int, error) {
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

// GetHomeworkHanded 查看上交的作业
func GetHomeworkHanded(groupId int, offset int, limit int) ([]Homework, int, error) {
	var handed []Homework
	d := model.DB.Self.Table("homeworks").
		Where("group_id = ?", groupId).
		Offset(offset).Limit(limit).
		Order("created_at desc").Find(&handed)
	if err := d.Error; err != nil {
		return handed, 0, err
	}
	var num int = 0
	for i, _ := range handed {
		num = i + 1
	}
	return handed, num, nil
}

// GetPublishedDetails 查看发布作业的细节
func GetPublishedDetails(id int) (*HomeworkPublished, error) {
	var homework HomeworkPublished
	if err := model.DB.Self.Model(HomeworkPublished{}).
		Where("id  = ?", id).Find(&homework).Error; err != nil {
		return nil, err
	}
	return &homework, nil
}

// ReviewHomework ...查阅作业
func ReviewHomework(id int) (*Homework, error) {
	var homework Homework
	if err := model.DB.Self.Model(Homework{}).Where("id = ?", id).Find(&homework).Error; err != nil {
		return nil, err
	}
	return &homework, nil
}

// UpdateUploaded ...修改上传的作业
func UpdateUploaded(id string, email string, title string, content string, fileUrl string) error {
	var hw Homework
	if err := model.DB.Self.Model(Homework{}).Where("id = ?", id).Find(&hw).Error; err != nil {
		return err
	}
	if hw.Email != email {
		return errors.New("permission denied")
	}
	if err := model.DB.Self.Model(Homework{}).Where("id = ?", id).Updates(Homework{
		Title:   title,
		Content: content,
		URL:     fileUrl,
	}).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePublished ...	修改发布的作业
func UpdatePublished(id, email, title, content, fileUrl string, groupID uint) error {
	var published HomeworkPublished
	if err := model.DB.Self.Model(HomeworkPublished{}).Where("id  = ?", id).Find(&published).Error; err != nil {
		return err
	}
	if published.Publisher != email {
		return errors.New("permission denied")
	}

	if err := model.DB.Self.Model(HomeworkPublished{}).Where("id = ?", id).Updates(HomeworkPublished{
		GroupID: groupID,
		Title:   title,
		Content: content,
		FileUrl: fileUrl,
	}).Error; err != nil {
		return err
	}
	return nil
}

// GetMine 查看某作业我的上传内容
func GetMine(email string, id string) ([]Homework, error) {
	var homework []Homework
	if err := model.DB.Self.Model(Homework{}).
		Where("email = ? AND homework_id = ?", email, id).
		Find(&homework).Error; err != nil {
		return nil, err
	}
	return homework, nil
}

// GetALl 查询我所有的作业
func GetAll(email string) ([]Homework, error) {
	var homework []Homework
	if err := model.DB.Self.Model(Homework{}).
		Where("email = ?", email).
		Find(&homework).Error; err != nil {
		return nil, err
	}
	return homework, nil
}

// GetAllPublished
func GetAllPublished(email string) ([]HomeworkPublished, error) {
	var form form2.FormModel
	if err := model.DB.Self.Model(form2.FormModel{}).
		Where("email = ?", email).
		Find(&form).Error; err != nil {
		return nil, err
	}
	var id uint
	switch form.Group {
	case "设计组":
		id = 1
	case "产品组":
		id = 2
	case "安卓组":
		id = 3
	case "前端组":
		id = 4
	case "后端组":
		id = 5

	}

	fmt.Println("id.....", id)
	var published []HomeworkPublished
	if err := model.DB.Self.Model(HomeworkPublished{}).
		Where("group_id = ?", id).
		Find(&published).Error; err != nil {
		return nil, err
	}
	return published, nil
}

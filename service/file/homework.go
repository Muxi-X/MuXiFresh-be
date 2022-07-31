package file

import (
	File "github.com/MuXiFresh-be/model/file"
)

// 提交作业
func HandInHomework(url string, email string, homeworkID uint) error {
	var homework File.Homework
	homework.URL = url
	homework.Email = email
	homework.HomeworkID = homeworkID

	return homework.Create()
}

// 发布作业
func PublishHomework(email string, groupID uint, content string) error {
	var homework = File.HomeworkPublished{
		Publisher: email,
		GroupID:   groupID,
		Content:   content,
	}
	return homework.Publish()

}

package file

import (
	Comment "github.com/MuXiFresh-be/model/comment"
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
func PublishHomework(email string, groupID uint, title string, content string) error {
	var homework = File.HomeworkPublished{
		Publisher: email,
		GroupID:   groupID,
		Content:   content,
	}
	return homework.Publish()

}

// 评论作业
func CommentHomework(email string, id uint, content string) error {
	var comment = Comment.Comment{
		Publisher:  email,
		HomeworkID: id,
		Content:    content,
	}
	return comment.Create()
}

// 删除评论
func Delete(id string, email string) error {
	return Comment.DeleteComment(id, email)
}

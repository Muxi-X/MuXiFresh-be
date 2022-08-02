package file

import (
	Comment "github.com/MuXiFresh-be/model/comment"
	File "github.com/MuXiFresh-be/model/file"
)

// 提交作业
func HandInHomework(title string, content string, homeworkID uint, url string, email string) error {
	var homework File.Homework
	homework.URL = url
	homework.Email = email
	homework.HomeworkID = homeworkID

	return homework.Create()
}

// 发布作业
func PublishHomework(email string, groupID uint, title string, content string, url string) error {
	var homework = File.HomeworkPublished{
		Publisher: email,
		Title:     title,
		GroupID:   groupID,
		Content:   content,
		FileUrl:   url,
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

// 获取评论
func GetComment(id string, offset int, limit int) ([]Comment.Comment, int, error) {
	var comments []Comment.Comment
	comments, num, err := Comment.GetCommentList(id, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return comments, num, err
}

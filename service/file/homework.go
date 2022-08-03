package file

import (
	Comment "github.com/MuXiFresh-be/model/comment"
	File "github.com/MuXiFresh-be/model/file"
)

// 提交作业
func HandInHomework(title string, content string, homeworkID uint, url string, email string) error {

	return File.Create(title, content, homeworkID, url, email)
}

// 发布作业
func PublishHomework(email string, ID uint, title string, content string, url string) error {

	return File.Publish(ID, title, content, email, url)

}

// 评论作业
func CommentHomework(email string, id uint, content string) error {
	return Comment.Create(email, id, content)
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

// 审阅作业
func Review(id string, offset int, limit int) ([]File.Homework, int, error) {
	return File.ReviewHomework(id, offset, limit)
}

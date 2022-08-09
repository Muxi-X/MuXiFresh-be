package file

import (
	Comment "github.com/MuXiFresh-be/model/comment"
	File "github.com/MuXiFresh-be/model/file"
	"github.com/MuXiFresh-be/pkg/errno"
)

// 提交作业
func HandInHomework(title string, content string, homeworkID uint, url string, email string) error {
	if err := File.Create(title, content, homeworkID, url, email); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// 发布作业
func PublishHomework(email string, ID uint, title string, content string, url string) error {
	if err := File.Publish(ID, title, content, email, url); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

// GetHomework ...获取不同组别的作业
func GetHomework(id int, offset int, limit int) ([]File.HomeworkPublished, int, error) {
	HW, num, err := File.GetHomework(id, offset, limit)
	if err != nil {
		return nil, 0, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return HW, num, nil
}

// 评论作业
func CommentHomework(email string, id uint, content string) error {
	if err := Comment.Create(email, id, content); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
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

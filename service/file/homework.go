package file

import (
	"fmt"
	Comment "github.com/MuXiFresh-be/model/comment"
	File "github.com/MuXiFresh-be/model/file"
	"github.com/MuXiFresh-be/pkg/errno"
)

// HandInHomework ...提交作业
func HandInHomework(title string, content string, homeworkID uint, url string, email string) (*File.Homework, error) {
	homework, err := File.Create(title, content, homeworkID, url, email)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return homework, nil
}

// PublishHomework ...发布作业
func PublishHomework(email string, ID uint, title string, content string, url string) (*File.HomeworkPublished, error) {
	homework, err := File.Publish(ID, title, content, email, url)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	return homework, nil
}

// GetHomeworkPublished ...获取不同组别的发布作业
func GetHomeworkPublished(id int, offset int, limit int) ([]File.HomeworkPublished, int, error) {
	HW, num, err := File.GetHomeworkPublished(id, offset, limit)
	if err != nil {
		return nil, 0, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return HW, num, nil
}

// GetHomeworkHanded ...获取不同组提交的作业
func GetHomeworkHanded(groupId int, offset int, limit int) ([]File.Homework, int, error) {
	handed, num, err := File.GetHomeworkHanded(groupId, offset, limit)
	if err != nil {
		return nil, 0, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return handed, num, nil
}

// CommentHomework ...评论作业
func CommentHomework(email string, id uint, content string) error {
	if err := Comment.Create(email, id, content); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// Delete ...删除评论
func Delete(id string, email string) error {
	if err := Comment.DeleteComment(id, email); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// GetComment ...获取评论
func GetComment(id string, offset int, limit int) ([]Comment.List, int, error) {
	comments, num, err := Comment.GetCommentList(id, offset, limit)
	if err != nil {
		return nil, 0, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return comments, num, nil
}

// Review ...查阅作业
func Review(id int) (*File.Homework, error) {
	homework, err := File.ReviewHomework(id)
	fmt.Println("------", homework, "-------")
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	return homework, nil
}

// GetHomeworkDetails ...获取作业的详细内容
func GetHomeworkDetails(id int, email string) (*File.HomeworkPublished, []File.Homework, error) {
	homeworkpublished, homework, err := File.GetPublishedDetails(id, email)
	if err != nil {
		return nil, nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return homeworkpublished, homework, nil
}

// ModifyHW ...修改作业
func ModifyHW(id string, email string, title string, content string, fileUrl string) error {
	if err := File.UpdateUploaded(id, email, title, content, fileUrl); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// ModifyPublished ...修改发布的作业
func ModifyPublished(id string, email string, groupID uint, title string, content string, fileUrl string) error {
	if err := File.UpdatePublished(id, email, title, content, fileUrl, groupID); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// GetMine ...获取我对应的提交作业
func GetMine(email string, id string) ([]File.Homework, error) {
	homework, err := File.GetMine(email, id)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return homework, nil
}

// GetAllMyHomework ...获取做过的全部作业
func GetAllMyHomework(email string) ([]File.Homework, error) {
	homework, err := File.GetAll(email)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return homework, nil
}

// GetAllPublished ...获取所有发布的作业
func GetAllPublished(email string) ([]File.HomeworkPublished, error) {
	published, err := File.GetAllPublished(email)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return published, nil
}

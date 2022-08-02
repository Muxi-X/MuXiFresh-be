package homework

type PublishHomeworkRequest struct {
	Title   string ` json:"title" bingding:"required"`
	Content string `json:"content" binding:"required"`
	GroupID uint   `json:"group_id" binding:"required"`
	FileUrl string `json:"file_url" `
}
type HomeworkRequest struct {
	Title      string ` json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	HomeworkID uint   `json:"homework_id" binding:"required"`
	FileUrl    string `json:"file_url" `
}

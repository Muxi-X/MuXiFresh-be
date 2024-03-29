package homework

type PublishHomeworkRequest struct {
	Title   string ` json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	GroupID uint   `json:"group_id" binding:"required"`
	FileUrl string `json:"file_url" `
}
type PublishedResponse struct {
	Title   string ` json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	GroupID uint   `json:"group_id" binding:"required"`
	FileUrl string `json:"file_url" `
	Status  int    `json:"status"`
}

type HomeworkRequest struct {
	Title      string ` json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	HomeworkID uint   `json:"homework_id" binding:"required"`
	FileUrl    string `json:"file_url" `
}

type ModifyPublishedHomeworkRequest struct {
	Title   string `json:"title" `
	Content string `json:"content" `
	GroupID uint   `json:"group_id"`
	FileUrl string `json:"file_url" `
}

type ModifyHomeworkRequest struct {
	Title   string ` json:"title" `
	Content string `json:"content"`
	FileUrl string `json:"file_url" `
}

type PerformanceResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	URL     string `json:"url"`
	Status  int    `json:"status"`
}

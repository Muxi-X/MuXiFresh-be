package homework

type HomeworkRequest struct {
	Title   string ` json:"title" bingding:"required"`
	Content string `json:"content" binding:"required"`
	GroupID uint   `json:"group_id" binding:"required"`
}

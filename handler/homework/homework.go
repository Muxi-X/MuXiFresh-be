package homework

type HomeworkRequest struct {
	Content string `json:"content" binding:"required"`
	GroupID uint   `json:"group_id" binding:"required"`
}

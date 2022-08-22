package form //用于存放和报名表相关的结构体

// createRequest create 请求
type createRequest struct {
	Name                string `json:"name" binding:"required"`
	Avatar              string `json:"avatar" binding:"required"`
	StudentId           string `json:"student_id" binding:"required"`
	College             string `json:"college" binding:"required"`
	Major               string `json:"major" binding:"required"`
	Grade               string `json:"grade" binding:"required"`
	Gender              string `json:"gender" binding:"required"`
	ContactWay          string `json:"contact_way" binding:"required"`
	ContactNumber       string `json:"contact_number" binding:"required"`
	Group               string `json:"group" binding:"required"`
	Reason              string `json:"reason" binding:"required"`
	Understand          string `json:"understand" binding:"required"`
	SelfIntroduction    string `json:"self_introduction" binding:"required"`
	IfOtherOrganization string `json:"if_other_organization" binding:"required"`
} // @name createRequest

// editRequest edit请求
type editRequest struct {
	Name                string `json:"name"`
	Avatar              string `json:"avatar"`
	StudentId           string `json:"student_id"`
	College             string `json:"college" `
	Major               string `json:"major"`
	Grade               string `json:"grade"`
	Gender              string `json:"gender"`
	ContactWay          string `json:"contact_way"`
	ContactNumber       string `json:"contact_number"`
	Group               string `json:"group" `
	Reason              string `json:"reason"`
	Understand          string `json:"understand"`
	SelfIntroduction    string `json:"self_introduction"`
	IfOtherOrganization string `json:"if_other_organization"`
} // @name editRequest

type searchRequest struct {
	Group string `json:"group"`
}

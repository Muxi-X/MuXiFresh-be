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

type viewothersRequest struct {
	Email string `json:"email"`
}

type searchResponse struct {
	Name          string `json:"name" gorm:"column:name;not null" binding:"required"`
	Email         string `json:"email" gorm:"column:email;default:null;unique"`
	Avatar        string `json:"avatar" gorm:"column:avatar"`
	//Role          uint32 `json:"role" gorm:"column:role;" binding:"required"`
	//Message       uint32 `json:"message" gorm:"column:message;" binding:"required"`
	//HashPassword  string `json:"hash_password" gorm:"column:hash_password;" binding:"required"`
	//StudentId     string `json:"student_id" gorm:"column:student_id;unique™"`
	College       string `json:"college" gorm:"column:college;"`
	//Major         string `json:"major" gorm:"column:major;"`
	Grade         string `json:"grade" gorm:"column:grade;"`
	//Gender        string `json:"gender" gorm:"column:gender;"`
	//ContactWay    string `json:"contact_way" gorm:"column:contact_way;"`
	//ContactNumber string `json:"contact_number" gorm:"column:contact_number;"`
}
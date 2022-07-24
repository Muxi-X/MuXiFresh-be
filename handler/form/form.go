package form //用于存放和报名表相关的结构体

// createRequest create 请求
type createRequest struct {
	Name                  string `json:"name"`
	Avatar                string `json:"avatar"`
	StudentId             string `json:"student_id"`
	College               string `json:"college" `
	Major                 string `json:"major"`
	Grade                 string `json:"grade"`
	Gender                string `json:"gender"`
	Contact_way           string `json:"contact_way"`
	Contact_number        string `json:"contact_number"`
	Group                 string `json:"group" `
	Reason                string `json:"reason"`
	Understand            string `json:"understand"`
	Self_introduction     string `json:"self_introduction"`
	If_other_organization string `json:"if_other_organization"`
} // @name createRequest

// editRequest edit请求
type editRequest struct {
	Name                  string `json:"name"`
	Avatar                string `json:"avatar"`
	StudentId             string `json:"student_id"`
	College               string `json:"college" `
	Major                 string `json:"major"`
	Grade                 string `json:"grade"`
	Gender                string `json:"gender"`
	Contact_way           string `json:"contact_way"`
	Contact_number        string `json:"contact_number"`
	Group                 string `json:"group" `
	Reason                string `json:"reason"`
	Understand            string `json:"understand"`
	Self_introduction     string `json:"self_introduction"`
	If_other_organization string `json:"if_other_organization"`
} // @name editRequest



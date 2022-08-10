package user

// RegisterRequest Register 请求
type RegisterRequest struct {
	StudentId string `json:"student_id" form:"student_id" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	//PasswordAgain string `json:"password_again" form:"password_again"`
	Email string `json:"email" form:"email" binding:"required"`
	Name  string `json:"name" form:"name" binding:"required"`
}

// loginRequest Login 请求
type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
} // @name loginRequest

// loginResponse Login 请求响应
type loginResponse struct {
	Token string `json:"token"`
} // @name loginResponse

// GetInfoRequest 获取 info 请求
type getInfoRequest struct {
	Ids []uint32 `json:"ids" binding:"required"`
} // @name GetInfoRequest

type userInfo struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

// GetInfoResponse 获取 info 响应
type getInfoResponse struct {
	List []userInfo `json:"list"`
} // @name getInfoResponse

// getProfileRequest 获取 profile 请求
type getProfileRequest struct {
	Id uint32 `json:"id"`
} // @name getProfileRequest

// userProfile 获取 profile 响应
type userProfile struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Role   uint32 `json:"role"`
} // @name userProfile

// listRequest 获取 userList 请求
type listRequest struct {
	Team  uint32 `json:"team"`
	Group uint32 `json:"group"`
} // @name listRequest

type user struct {
	Id     uint32 `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Role   uint32 `json:"role"`
} // @name user

// listResponse 获取 userList 响应
type listResponse struct {
	Count uint32 `json:"count"`
	List  []user `json:"list"`
} // @name listResponse

// updateInfoRequest 更新 userInfo 请求
type updateInfoRequest struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
} // @name updateInfoRequest

// updateTeamGroupRequest
type updateTeamGroupRequest struct {
	Ids   []uint32 `json:"ids"`
	Value uint32   `json:"value"`
	Kind  uint32   `json:"kind"`
} // @name updateTeamGroupRequest

// AdminResponse
type AdminResponse struct {
	ID        uint   `json:"id"`
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	Avatar    string `json:"avatar"`
}

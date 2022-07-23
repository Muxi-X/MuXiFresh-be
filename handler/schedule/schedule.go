package schedule

// GetInfoResponse 获取 info 响应
type GetInfoReponse struct {
	Name            string `json:"name"`
	Collage         string `json:"collage"`
	Major           string `json:"major"`
	Group           string `json:"group"`
	FormStatus      int    `json:"form_status"`
	WorkStatus      int    `json:"work_status"`
	AdmissionStatus int    `json:"admission_status"`
} // @name getInfoResponse

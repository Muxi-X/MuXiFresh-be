package form

import (
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/service/form"
	"github.com/gin-gonic/gin"
)

// @Summary "分组查看报名者信息"
// @Description "输入两位数字符如‘11’ -前一位表示（1-产品，2-安卓，3-设计，4-前端，5-后端）-后一位表示（1-已报名，2-初试过，3-面试过）"
// @Tags form
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body searchRequest true "按组查询报名表"
// @Success 200
// @Router /form/search [post]
func Search(c *gin.Context) {
	email := c.MustGet("email").(string)
	if email == "" {
		return
	}
	var group searchRequest
	if err := c.BindJSON(&group); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	/* Form, err := form.SearchForm(group.Group)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	} */
	//SendResponse(c, nil, Form)
	//fmt.Println("123", group.Group)
	Userinfo, err := form.SearchByGroupForUser(group.Group)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	len := len(Userinfo)
	responseinfo := make([]searchResponse, len)
	for i := 0; i < len; i++ {
		responseinfo[i].Name = Userinfo[i].Name
		responseinfo[i].Email = Userinfo[i].Email
		responseinfo[i].Avatar = Userinfo[i].Avatar
		responseinfo[i].College = Userinfo[i].College
		responseinfo[i].Grade = Userinfo[i].Grade
	}
	SendResponse(c, nil, responseinfo)
}

package user

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	. "github.com/MuXiFresh-be/handler"
	"github.com/MuXiFresh-be/model"
	userModel "github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req RegisterRequest
	var db *model.Database

	if err := c.ShouldBind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine()) //, errno.ErrBind)
		fmt.Println(req)
		return
	}

	if req.Password != req.PasswordAgain {
		SendBadRequest(c, errno.ErrPasswordRepetition, nil, "please Re-enter the password", GetLine()) //, errno.ErrPasswordRepetition)
		return
	}

	db.Init()
	defer db.Close()

	if err := userModel.IfExist(req.StudentId, req.Email, req.Name); err != nil {
		SendBadRequest(c, errno.ErrUserExisted, nil, err.Error(), GetLine())
		return
	}
	
	var user userModel.UserModel
	user.Name = req.Name
	user.StudentId = req.StudentId
	user.Email = req.Email
	md5 := md5.New()
	md5.Write([]byte(req.Password))
	user.HashPassword = hex.EncodeToString(md5.Sum(nil))
	if err := model.DB.Self.Create(&user).Error; err != nil {
		SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, "succeed in registration")

}

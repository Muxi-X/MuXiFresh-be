package service

import (
	Md5 "crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/MuXiFresh-be/model"
	"github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
	Token "github.com/MuXiFresh-be/pkg/token"
	"github.com/MuXiFresh-be/util"
)

//
// import (
// 	"github.com/MuXiFresh-be/model/user"
// 	"github.com/MuXiFresh-be/pkg/auth"
// 	"github.com/MuXiFresh-be/pkg/constvar"
// 	"github.com/MuXiFresh-be/pkg/token"
// 	// pb "github.com/MuXiFresh-be/proto"
// 	"github.com/MuXiFresh-be/util"
// )

// Login ... 登录
func Login(email string, pwd string) (string, error) {
	// 根据 studentId 在 DB 查询 user
	//user, err := userModel.GetUserByStudentId(studentId)

	var userInfo user.UserModel

	if err := model.DB.Self.Where("email = ?", email).First(&userInfo); err.Error != nil {
		fmt.Println(err, err.Error)
		return "", errno.ErrUserNotExisted
	}

	md5 := Md5.New()
	md5.Write([]byte(pwd))
	hashPwd := hex.EncodeToString(md5.Sum(nil))

	if userInfo.HashPassword != hashPwd {
		return "", errno.ErrPasswordIncorrect
	}

	//if err != nil {
	//	return "", errno.ServerErr(errno.ErrDatabase, err.Error())
	//}
	//if user == nil {
	// 注册
	//}
	// 	if err := auth.GetUserInfoFormOne(studentId, pwd); err != nil {
	// 		return "", err
	// 	}
	// 	info := &RegisterInfo{
	// 		StudentId: studentId,
	// 		Password:  pwd,
	// 		Role:      constvar.Normal,
	// 	}
	// 	// 用户未注册，自动注册
	// 	if err := user.RegisterUser(info); err != nil {
	// 		return "", err
	// 	}
	// 	// 注册后重新查询
	// 	user, err = user.GetUserByStudentId(studentId)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// }
	//
	// 生成 auth token
	var payload = Token.TokenPayload{
		Id:      userInfo.ID,
		Role:    userInfo.Role,
		Email:   userInfo.Email,
		Expired: util.GetExpiredTime(),
	}
	token, err := payload.GenerateToken()
	if err != nil {
		return "", err
	}

	return token, nil

}

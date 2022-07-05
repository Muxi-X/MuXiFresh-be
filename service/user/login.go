package service

import (
	"github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
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
func Login(studentId string, pwd string) (string, error) {
	// 根据 studentId 在 DB 查询 user
	user, err := user.GetUserByStudentId(studentId)

	if err != nil {
		return "", errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	if user == nil {
		// 注册
	}
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
	// // 生成 auth token
	// token, err := token.GenerateToken(&token.TokenPayload{
	// 	Id:      user.Id,
	// 	Expired: util.GetExpiredTime(),
	// })
	// if err != nil {
	// 	return "", err
	// }
	//
	// return token, nil

	return "", nil
}

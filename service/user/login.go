package service

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/MuXiFresh-be/model"
	userModel "github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
	"github.com/MuXiFresh-be/pkg/token"
	"github.com/MuXiFresh-be/util"
)

=======
>>>>>>> parent of 5f5d095 (feat: add init to log; refact: pkg/errno; feat: add demo to login)
=======
>>>>>>> parent of 5f5d095 (feat: add init to log; refact: pkg/errno; feat: add demo to login)
=======
>>>>>>> parent of 5f5d095 (feat: add init to log; refact: pkg/errno; feat: add demo to login)
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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func Login(email string, pwd string) (string, error) {
	// 根据 studentId 在 DB 查询 user
	//user, err := userModel.GetUserByStudentId(studentId)

	var userInfo userModel.UserModel
	var db *model.Database

	db.Init()

	if err := model.DB.Self.Where("email=?", email).First(&userInfo); err.Error != nil {
		fmt.Println(err, err.Error)
		return "", errno.ErrUserExisted
	}

	db.Close()

	md5 := md5.New()
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
=======
=======
>>>>>>> parent of 5f5d095 (feat: add init to log; refact: pkg/errno; feat: add demo to login)
=======
>>>>>>> parent of 5f5d095 (feat: add init to log; refact: pkg/errno; feat: add demo to login)
// 否则，用 code 获取 oauth 的 access token，并生成该应用的 auth token，返回给前端。
func Login(studentId string, pwd string) (string, error) {
	// 根据 studentId 在 DB 查询 user
	// user, err := user.GetUserByStudentId(studentId)
	//
	// if err != nil {
	// 	return "", err
	// }
	// if user == nil {
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> parent of 5f5d095 (feat: add init to log; refact: pkg/errno; feat: add demo to login)
=======
>>>>>>> parent of 5f5d095 (feat: add init to log; refact: pkg/errno; feat: add demo to login)
=======
>>>>>>> parent of 5f5d095 (feat: add init to log; refact: pkg/errno; feat: add demo to login)
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
	token, err := token.GenerateToken(email, util.GetExpiredTime())
	if err != nil {
		return token, err
	}

	return token, nil

}

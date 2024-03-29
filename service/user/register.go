package service

import (
	MD5 "crypto/md5"
	"encoding/hex"

	Schedule "github.com/MuXiFresh-be/model/schedule"
	User "github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
)

func Register(StudentId string, email string, name string, password string) error {
	//前端自己验证两次密码是否一致
	//if req.Password != req.PasswordAgain {
	//	SendBadRequest(c, errno.ErrPasswordRepetition, nil, "please Re-enter the password", GetLine())
	//	return
	//}

	if err := User.IfExist(StudentId, email, name); err != nil {
		return errno.ServerErr(errno.ErrUserExisted, err.Error())
	}

	user := User.UserModel{
		Name:      name,
		StudentId: StudentId,
		Email:     email,
		Role:      1,
	}

	md5 := MD5.New()
	md5.Write([]byte(password))
	user.HashPassword = hex.EncodeToString(md5.Sum(nil))
	if err := user.CreateUser(); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

func Create(email string, name string) error {
	var schedule = Schedule.ScheduleModel{
		AdmissionStatus: 0,
		Email:           email,
		FormStatus:      0,
		WorkStatus:      0,
		Group:           "",
		Name:            name,
	}

	if err := schedule.Create(); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

package form

import (
	"github.com/MuXiFresh-be/model/form"
	"github.com/MuXiFresh-be/model/user"
)

//根据组别返回该组别全部成员的form信息
func SearchByGroup(group string) ([]form.FormModel, error) {
	var forms []form.FormModel
	forms, err := form.Search(group)
	if err != nil {
		return forms, err
	}
	return forms, nil
}

//根据返回全部成员form信息中的email返回全部成员的User信息
func SearchByGroupForUser(group string) ([]user.UserModel, error) {
	var users []user.UserModel
	users, err := form.GetUserInfoByGroup(group)
	if err != nil {
		return users,err
	}
	return users, nil
}
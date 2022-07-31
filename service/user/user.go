package service

import (
	"github.com/MuXiFresh-be/model/user"
)

func GetProfile(id int) (*user.UserModel, error) {
	var user user.UserModel
	if err := user.GerInfo(id); err != nil {
		return nil, err
	}
	return &user, nil
}

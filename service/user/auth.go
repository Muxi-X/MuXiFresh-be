package service

import (
	"github.com/MuXiFresh-be/model/user"
	"github.com/MuXiFresh-be/pkg/errno"
)

// Authorize ... 授权
func Authorize(email string, role int) error {
	if err := user.Authorize(email, role); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

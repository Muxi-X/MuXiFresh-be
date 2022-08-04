package service

import "github.com/MuXiFresh-be/model/user"

// Authorize ... 授权
func Authorize(id int, role int) error {
	return user.Authorize(id, role)
}
